package postgres

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/odpf/compass/core/asset"
)

type LineageRepository struct {
	client *Client
}

// NewLineageRepository initializes lineage repository
func NewLineageRepository(client *Client) (*LineageRepository, error) {
	if client == nil {
		return nil, errors.New("postgres client is nil")
	}
	return &LineageRepository{
		client: client,
	}, nil
}

// GetGraph returns a graph that contains list of relations of a given node
func (repo *LineageRepository) GetGraph(ctx context.Context, node asset.LineageNode) (asset.LineageGraph, error) {
	var graph asset.LineageGraph

	upstreams, err := repo.getUpstreamsGraph(ctx, node)
	if err != nil {
		return graph, fmt.Errorf("error fetching upstreams graph: %w", err)
	}
	downstreams, err := repo.getDownstreamsGraph(ctx, node)
	if err != nil {
		return graph, fmt.Errorf("error fetching upstreams graph: %w", err)
	}

	graph = append(upstreams, downstreams...)

	return graph, nil
}

// Upsert insert or delete connections of a given node by comparing them with current state
func (repo *LineageRepository) Upsert(ctx context.Context, node asset.LineageNode, upstreams, downstreams []asset.LineageNode) error {
	currentGraph, err := repo.getDirectLineage(ctx, node)
	if err != nil {
		return fmt.Errorf("error getting node's direct lineage: %w", err)
	}
	newGraph := repo.buildGraph(node, upstreams, downstreams)
	toInserts, toRemoves := repo.compareGraph(currentGraph, newGraph)
	toRemoves = repo.filterSelfDeleteOnly(node, toRemoves)

	return repo.client.RunWithinTx(ctx, func(tx *sqlx.Tx) error {
		err := repo.insertGraph(ctx, tx, toInserts)
		if err != nil {
			return fmt.Errorf("error inserting graph: %w", err)
		}

		err = repo.removeGraph(ctx, tx, toRemoves)
		if err != nil {
			return fmt.Errorf("error removing graph: %w", err)
		}

		return nil
	})
}

func (repo *LineageRepository) buildGraph(node asset.LineageNode, upstreams, downstreams []asset.LineageNode) (graph asset.LineageGraph) {
	for _, us := range upstreams {
		graph = append(graph, asset.LineageEdge{
			Source: us.URN,
			Target: node.URN,
			Prop: map[string]interface{}{
				"root": node.URN, // this is to note which node is updating the relation
			},
		})
	}
	for _, ds := range downstreams {
		graph = append(graph, asset.LineageEdge{
			Source: node.URN,
			Target: ds.URN,
			Prop: map[string]interface{}{
				"root": node.URN, // this is to note which node is updating the relation
			},
		})
	}

	return
}

// filterSelfDeleteOnly filters edges that are not created by the given node
// it uses prop["root"] field to figure out which node (source or target) is latest updater of the edge,
// and only allow that node to delete the relation
func (repo *LineageRepository) filterSelfDeleteOnly(node asset.LineageNode, toRemoves asset.LineageGraph) (res asset.LineageGraph) {
	for _, edge := range toRemoves {
		rootURN, ok := edge.Prop["root"]
		if ok && rootURN != node.URN {
			continue
		}
		res = append(res, edge)
	}

	return
}

func (repo *LineageRepository) insertGraph(ctx context.Context, execer sqlx.ExecerContext, graph asset.LineageGraph) error {
	if len(graph) == 0 {
		return nil
	}

	builder := sq.Insert("lineage_graph").Columns("source", "target", "prop").PlaceholderFormat(sq.Dollar)
	for _, edge := range graph {
		builder = builder.Values(edge.Source, edge.Target, edge.Prop)
	}
	builder = builder.Suffix("ON CONFLICT DO NOTHING")

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("error building query: %w", err)
	}

	_, err = execer.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("error executing insert query: %w", err)
	}

	return nil
}

func (repo *LineageRepository) removeGraph(ctx context.Context, execer sqlx.ExecerContext, graph asset.LineageGraph) error {
	if len(graph) == 0 {
		return nil
	}

	builder := sq.Delete("lineage_graph").PlaceholderFormat(sq.Dollar)
	conditions := sq.Or{}
	for _, edge := range graph {
		conditions = append(conditions,
			sq.Eq{"source": edge.Source, "target": edge.Target},
		)
	}
	builder = builder.Where(conditions)

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("error building query: %w", err)
	}
	_, err = execer.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (repo *LineageRepository) compareGraph(current, new asset.LineageGraph) (toInserts, toRemoves asset.LineageGraph) {
	if len(current) == 0 && len(new) == 0 {
		return
	}

	currMap := map[string]asset.LineageEdge{}
	for _, c := range current {
		key := c.Source + c.Target
		currMap[key] = c
	}

	for _, n := range new {
		key := n.Source + n.Target
		_, exists := currMap[key]
		if exists {
			// if exists, it means that both new and current have it.
			// we remove it from the map,
			// so that what's left in the map is the that only exists in current
			// and have to be removed
			delete(currMap, key)
		} else {
			toInserts = append(toInserts, n)
		}
	}

	for _, edge := range currMap {
		toRemoves = append(toRemoves, edge)
	}

	return
}

func (repo *LineageRepository) getUpstreamsGraph(ctx context.Context, node asset.LineageNode) (asset.LineageGraph, error) {
	var graph asset.LineageGraph

	query := repo.getUpstreamQuery()

	var gm LineageGraphModel
	err := repo.client.db.SelectContext(ctx, &gm, query, node.URN)
	if err != nil {
		return graph, err
	}

	graph = gm.toGraph()

	return graph, nil
}

func (repo *LineageRepository) getDownstreamsGraph(ctx context.Context, node asset.LineageNode) (asset.LineageGraph, error) {
	var graph asset.LineageGraph

	query := repo.getDownstreamQuery()

	var gm LineageGraphModel
	err := repo.client.db.SelectContext(ctx, &gm, query, node.URN)
	if err != nil {
		return graph, err
	}

	graph = gm.toGraph()

	return graph, nil
}

func (repo *LineageRepository) getUpstreamQuery() string {
	return `
		WITH RECURSIVE search_graph (
			source, target, prop, depth, path
		) AS (
				SELECT
					source,
					target,
					prop,
					1 as depth,
					ARRAY[target] as path
				FROM
					lineage_graph
				WHERE
					target = $1
			UNION ALL
				SELECT
					lg.source,
					lg.target,
					lg.prop,
					sg.depth + 1,
					sg.path || lg.target
				FROM
					lineage_graph lg,
					search_graph sg
				WHERE
					lg.target = sg.source
					AND lg.target <> ALL(sg.path)
		)
		
		SELECT source, target, prop FROM search_graph;  
	`
}

func (repo *LineageRepository) getDownstreamQuery() string {
	return `
		WITH RECURSIVE search_graph (
			source, target, prop, depth, path
		) AS (
				SELECT
					source,
					target,
					prop,
					1 as depth,
					ARRAY[source] as path
				FROM
					lineage_graph
				WHERE
					source = $1
			UNION ALL
				SELECT
					lg.source,
					lg.target,
					lg.prop,
					sg.depth + 1,
					sg.path || lg.source
				FROM
					lineage_graph lg,
					search_graph sg
				WHERE
					lg.source = sg.target
					AND lg.source <> ALL(sg.path)
		)
		
		SELECT source, target, prop FROM search_graph;
	`
}

func (repo *LineageRepository) getDirectLineage(ctx context.Context, node asset.LineageNode) (graph asset.LineageGraph, err error) {
	query := `SELECT * FROM lineage_graph WHERE (source = $1 OR target = $1)`
	var gm LineageGraphModel
	if err = repo.client.db.SelectContext(ctx, &gm, query, node.URN); err != nil {
		err = fmt.Errorf("error running fetch direct nodes query: %w", err)
		return
	}

	graph = gm.toGraph()

	return
}
