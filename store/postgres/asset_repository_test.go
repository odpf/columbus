package postgres_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
	"github.com/odpf/columbus/asset"
	"github.com/odpf/columbus/store/postgres"
	"github.com/odpf/columbus/user"
	"github.com/odpf/salt/log"
	"github.com/ory/dockertest/v3"
	"github.com/r3labs/diff/v2"
	"github.com/stretchr/testify/suite"
)

var defaultAssetUpdaterUserID = uuid.NewString()

type AssetRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *postgres.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.AssetRepository
	userRepo   *postgres.UserRepository
	users      []user.User
	builder    sq.SelectBuilder
}

func (r *AssetRepositoryTestSuite) SetupSuite() {
	var err error

	logger := log.NewLogrus()
	r.client, r.pool, r.resource, err = newTestClient(logger)
	if err != nil {
		r.T().Fatal(err)
	}

	r.ctx = context.TODO()
	r.userRepo, err = postgres.NewUserRepository(r.client)
	if err != nil {
		r.T().Fatal(err)
	}

	r.repository, err = postgres.NewAssetRepository(r.client, r.userRepo, defaultGetMaxSize, defaultProviderName)
	if err != nil {
		r.T().Fatal(err)
	}

	r.users = r.createUsers(r.userRepo)
}

func (r *AssetRepositoryTestSuite) createUsers(userRepo user.Repository) []user.User {
	var err error
	users := []user.User{}

	user1 := user.User{UUID: uuid.NewString(), Email: "user-test-1@odpf.io", Provider: defaultProviderName}
	user1.ID, err = userRepo.Create(r.ctx, &user1)
	r.Require().NoError(err)
	users = append(users, user1)

	user2 := user.User{UUID: uuid.NewString(), Email: "user-test-2@odpf.io", Provider: defaultProviderName}
	user2.ID, err = userRepo.Create(r.ctx, &user2)
	r.Require().NoError(err)
	users = append(users, user2)

	user3 := user.User{UUID: uuid.NewString(), Email: "user-test-3@odpf.io", Provider: defaultProviderName}
	user3.ID, err = userRepo.Create(r.ctx, &user3)
	r.Require().NoError(err)
	users = append(users, user3)

	user4 := user.User{UUID: uuid.NewString(), Email: "user-test-4@odpf.io", Provider: defaultProviderName}
	user4.ID, err = userRepo.Create(r.ctx, &user4)
	r.Require().NoError(err)
	users = append(users, user4)

	return users
}

func (r *AssetRepositoryTestSuite) TearDownSuite() {
	// Clean tests
	err := r.client.Close()
	if err != nil {
		r.T().Fatal(err)
	}
	err = purgeDocker(r.pool, r.resource)
	if err != nil {
		r.T().Fatal(err)
	}
}

func (r *AssetRepositoryTestSuite) insertRecord() (assets []asset.Asset) {
	filePath := "./testdata/mock-asset-data.json"
	testFixtureJSON, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []asset.Asset{}
	}

	var data []asset.Asset
	if err = json.Unmarshal(testFixtureJSON, &data); err != nil {
		return []asset.Asset{}
	}

	for _, d := range data {
		ast := asset.Asset{
			URN:         d.URN,
			Name:        d.Name,
			Type:        d.Type,
			Service:     d.Service,
			Description: d.Description,
			Data:        d.Data,
			Version:     asset.BaseVersion,
			UpdatedBy:   r.users[0],
		}

		id, err := r.repository.Upsert(r.ctx, &ast)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		ast.ID = id
		assets = append(assets, ast)
	}

	return assets
}

func (r *AssetRepositoryTestSuite) TestBuildFilterQuery() {
	r.builder = sq.Select(`a.test as test`)

	testCases := []struct {
		description   string
		config        asset.Filter
		expectedQuery string
	}{
		{
			description: "should return sql query with types filter",
			config: asset.Filter{
				Types: []asset.Type{asset.TypeTable},
			},
			expectedQuery: `type IN ($1)`,
		},
		{
			description: "should return sql query with services filter",
			config: asset.Filter{
				Services: []string{"mysql", "kafka"},
			},
			expectedQuery: `service IN ($1,$2)`,
		},
		{
			description: "should return sql query with query fields filter",
			config: asset.Filter{
				QueryFields: []string{"name", "description"},
				Query:       "demo",
			},
			expectedQuery: `(name ILIKE $1 OR description ILIKE $2)`,
		},
		{
			description: "should return sql query with nested data query filter",
			config: asset.Filter{
				QueryFields: []string{"data.landscape.properties.project-id", "description"},
				Query:       "columbus_002",
			},
			expectedQuery: `(data->'landscape'->'properties'->>'project-id' ILIKE $1 OR description ILIKE $2)`,
		},
		{
			description: "should return sql query with asset's data fields filter",
			config: asset.Filter{
				Data: map[string]string{
					"entity":  "odpf",
					"country": "th",
				},
			},
			expectedQuery: `data->>'entity' = 'odpf' AND data->>'country' = 'th'`,
		},
		{
			description: "should return sql query with asset's nested data fields filter",
			config: asset.Filter{
				Data: map[string]string{
					"landscape.properties.project-id": "columbus_001",
					"country":                         "vn",
				},
			},
			expectedQuery: `data->'landscape'->'properties'->>'project-id' = 'columbus_001' AND data->>'country' = 'vn'`,
		},
	}

	for _, testCase := range testCases {
		r.Run(testCase.description, func() {
			result := r.repository.BuildFilterQuery(r.builder, testCase.config)
			query, _, err := result.ToSql()
			r.Require().NoError(err)
			query, err = sq.Dollar.ReplacePlaceholders(query)
			r.Require().NoError(err)

			actualQuery := strings.Split(query, "WHERE ")
			r.Equal(testCase.expectedQuery, actualQuery[1])
		})
	}
}

func (r *AssetRepositoryTestSuite) TestGetAll() {
	assets := r.insertRecord()

	r.Run("should return all assets limited by default size", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{})
		r.Require().NoError(err)
		r.Require().Len(results, defaultGetMaxSize)
		for i := 0; i < defaultGetMaxSize; i++ {
			r.assertAsset(&assets[i], &results[i])
		}
	})

	r.Run("should override default size using GetConfig.Size", func() {
		size := 6
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Size: size,
		})
		r.Require().NoError(err)
		r.Require().Len(results, size)
		for i := 0; i < size; i++ {
			r.assertAsset(&assets[i], &results[i])
		}
	})

	r.Run("should fetch assets by offset defined in GetConfig.Offset", func() {
		offset := 2
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Offset: offset,
		})
		r.Require().NoError(err)
		for i := offset; i > defaultGetMaxSize+offset; i++ {
			r.assertAsset(&assets[i], &results[i-offset])
		}
	})

	r.Run("should filter using type", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Types:         []asset.Type{asset.TypeTable},
			SortBy:        "urn",
			SortDirection: "desc",
		})
		r.Require().NoError(err)

		expectedURNs := []string{"i-undefined-dfgdgd-avi", "e-test-grant2"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter using service", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Services: []string{"mysql", "kafka"},
			SortBy:   "urn",
		})
		r.Require().NoError(err)

		expectedURNs := []string{"c-demo-kafka", "f-john-test-001", "i-test-grant", "i-undefined-dfgdgd-avi"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter using query fields", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			QueryFields: []string{"name", "description"},
			Query:       "demo",
			SortBy:      "urn",
		})
		r.Require().NoError(err)

		expectedURNs := []string{"c-demo-kafka", "e-test-grant2"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter only using nested query data fields", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			QueryFields: []string{"data.landscape.properties.project-id", "data.title"},
			Query:       "columbus_001",
			SortBy:      "urn",
		})
		r.Require().NoError(err)

		expectedURNs := []string{"i-test-grant", "j-xcvcx"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter using query field with nested query data fields", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			QueryFields: []string{"data.landscape.properties.project-id", "description"},
			Query:       "columbus_002",
			SortBy:      "urn",
		})
		r.Require().NoError(err)

		expectedURNs := []string{"g-jane-kafka-1a", "h-test-new-kafka"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter using asset's data fields", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Data: map[string]string{
				"entity":  "odpf",
				"country": "th",
			},
		})
		r.Require().NoError(err)

		expectedURNs := []string{"e-test-grant2", "h-test-new-kafka", "i-test-grant"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})

	r.Run("should filter using asset's nested data fields", func() {
		results, err := r.repository.GetAll(r.ctx, asset.Filter{
			Data: map[string]string{
				"landscape.properties.project-id": "columbus_001",
				"country":                         "vn",
			},
		})
		r.Require().NoError(err)

		expectedURNs := []string{"j-xcvcx"}
		r.Equal(len(expectedURNs), len(results))
		for i := range results {
			r.Equal(expectedURNs[i], results[i].URN)
		}
	})
}

func (r *AssetRepositoryTestSuite) TestGetCount() {
	// populate assets
	total := 12
	typ := asset.TypeJob
	service := []string{"service-getcount"}
	for i := 0; i < total; i++ {
		ast := asset.Asset{
			URN:       fmt.Sprintf("urn-getcount-%d", i),
			Type:      typ,
			Service:   service[0],
			UpdatedBy: r.users[0],
		}
		id, err := r.repository.Upsert(r.ctx, &ast)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		ast.ID = id
	}

	r.Run("should return total assets with filter", func() {
		actual, err := r.repository.GetCount(r.ctx, asset.Filter{
			Types:    []asset.Type{typ},
			Services: service,
		})
		r.Require().NoError(err)
		r.Equal(total, actual)
	})
}

func (r *AssetRepositoryTestSuite) TestGetByID() {
	r.Run("return error from client if asset not an uuid", func() {
		_, err := r.repository.GetByID(r.ctx, "invalid-uuid")
		r.Error(err)
		r.Contains(err.Error(), "invalid asset id: \"invalid-uuid\"")
	})

	r.Run("return NotFoundError if asset does not exist", func() {
		uuid := "2aabb450-f986-44e2-a6db-7996861d5004"
		_, err := r.repository.GetByID(r.ctx, uuid)
		r.ErrorAs(err, &asset.NotFoundError{AssetID: uuid})
	})

	r.Run("return correct asset from db", func() {
		asset1 := asset.Asset{
			URN:       "urn-gbi-1",
			Type:      "table",
			Service:   "bigquery",
			Version:   asset.BaseVersion,
			UpdatedBy: r.users[1],
		}
		asset2 := asset.Asset{
			URN:       "urn-gbi-2",
			Type:      "topic",
			Service:   "kafka",
			Version:   asset.BaseVersion,
			UpdatedBy: r.users[1],
		}

		var err error
		id, err := r.repository.Upsert(r.ctx, &asset1)
		r.Require().NoError(err)
		r.NotEmpty(id)
		asset1.ID = id

		id, err = r.repository.Upsert(r.ctx, &asset2)
		r.Require().NoError(err)
		r.NotEmpty(id)
		asset2.ID = id

		result, err := r.repository.GetByID(r.ctx, asset2.ID)
		r.NoError(err)
		asset2.UpdatedBy = r.users[1]
		r.assertAsset(&asset2, &result)
	})

	r.Run("return owners if any", func() {

		ast := asset.Asset{
			URN:     "urn-gbi-3",
			Type:    "table",
			Service: "bigquery",
			Owners: []user.User{
				r.users[1],
				r.users[2],
			},
			UpdatedBy: r.users[1],
		}

		id, err := r.repository.Upsert(r.ctx, &ast)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		ast.ID = id

		result, err := r.repository.GetByID(r.ctx, ast.ID)
		r.NoError(err)
		r.Len(result.Owners, len(ast.Owners))
		for i, owner := range result.Owners {
			r.Equal(ast.Owners[i].ID, owner.ID)
		}
	})
}

func (r *AssetRepositoryTestSuite) TestFind() {
	r.Run("return NotFoundError if asset does not exist", func() {
		urn := "some-urn"
		typ := asset.TypeDashboard
		service := "bigquery"
		_, err := r.repository.Find(r.ctx, urn, typ, service)
		r.ErrorAs(err, &asset.NotFoundError{URN: urn, Type: typ, Service: service})
	})

	r.Run("return correct asset from db", func() {
		asset1 := asset.Asset{
			URN:       "urn-find-1",
			Type:      "table",
			Service:   "bigquery",
			Version:   asset.BaseVersion,
			UpdatedBy: r.users[1],
		}
		asset2 := asset.Asset{
			URN:       "urn-find-2",
			Type:      "topic",
			Service:   "kafka",
			Version:   asset.BaseVersion,
			UpdatedBy: r.users[1],
		}

		var err error
		id, err := r.repository.Upsert(r.ctx, &asset1)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		asset1.ID = id

		id, err = r.repository.Upsert(r.ctx, &asset2)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		asset2.ID = id

		result, err := r.repository.Find(r.ctx, asset2.URN, asset2.Type, asset2.Service)
		r.NoError(err)
		asset2.UpdatedBy = r.users[1]
		r.assertAsset(&asset2, &result)

		// clean up
		err = r.repository.Delete(r.ctx, asset1.ID)
		r.Require().NoError(err)
		err = r.repository.Delete(r.ctx, asset2.ID)
		r.Require().NoError(err)
	})

	r.Run("return owners if any", func() {
		ast := asset.Asset{
			URN:     "urn-find-3",
			Type:    "table",
			Service: "bigquery",
			Owners: []user.User{
				r.users[1],
				r.users[2],
			},
			UpdatedBy: r.users[1],
		}

		id, err := r.repository.Upsert(r.ctx, &ast)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		ast.ID = id

		result, err := r.repository.Find(r.ctx, ast.URN, ast.Type, ast.Service)
		r.NoError(err)
		r.Len(result.Owners, len(ast.Owners))
		for i, owner := range result.Owners {
			r.Equal(ast.Owners[i].ID, owner.ID)
		}

		// clean up
		err = r.repository.Delete(r.ctx, ast.ID)
		r.Require().NoError(err)
	})
}

func (r *AssetRepositoryTestSuite) TestVersions() {
	assetURN := uuid.NewString() + "urn-u-2-version"
	// v0.1
	astVersioning := asset.Asset{
		URN:       assetURN,
		Type:      "table",
		Service:   "bigquery",
		UpdatedBy: r.users[1],
	}

	id, err := r.repository.Upsert(r.ctx, &astVersioning)
	r.Require().NoError(err)
	r.Require().NotEmpty(id)
	astVersioning.ID = id

	// v0.2
	astVersioning.Description = "new description in v0.2"
	id, err = r.repository.Upsert(r.ctx, &astVersioning)
	r.Require().NoError(err)
	r.Require().Equal(id, astVersioning.ID)

	// v0.3
	astVersioning.Owners = []user.User{
		{
			Email: "user@odpf.io",
		},
		{
			Email:    "meteor@odpf.io",
			Provider: "meteor",
		},
	}
	id, err = r.repository.Upsert(r.ctx, &astVersioning)
	r.Require().NoError(err)
	r.Require().Equal(id, astVersioning.ID)

	// v0.4
	astVersioning.Data = map[string]interface{}{
		"data1": float64(12345),
	}
	id, err = r.repository.Upsert(r.ctx, &astVersioning)
	r.Require().NoError(err)
	r.Require().Equal(id, astVersioning.ID)

	// v0.5
	astVersioning.Labels = map[string]string{
		"key1": "value1",
	}

	id, err = r.repository.Upsert(r.ctx, &astVersioning)
	r.Require().NoError(err)
	r.Require().Equal(id, astVersioning.ID)

	r.Run("should return 3 last versions of an assets if there are exist", func() {

		expectedAssetVersions := []asset.Asset{
			{
				ID:      astVersioning.ID,
				URN:     assetURN,
				Type:    "table",
				Service: "bigquery",
				Version: "0.5",
				Changelog: diff.Changelog{
					diff.Change{Type: "create", Path: []string{"labels", "key1"}, From: interface{}(nil), To: "value1"},
				},
				UpdatedBy: r.users[1],
			},
			{
				ID:      astVersioning.ID,
				URN:     assetURN,
				Type:    "table",
				Service: "bigquery",
				Version: "0.4",
				Changelog: diff.Changelog{
					diff.Change{Type: "create", Path: []string{"data", "data1"}, From: interface{}(nil), To: float64(12345)},
				},
				UpdatedBy: r.users[1],
			},
			{
				ID:      astVersioning.ID,
				URN:     assetURN,
				Type:    "table",
				Service: "bigquery",
				Version: "0.3",
				Changelog: diff.Changelog{
					diff.Change{Type: "create", Path: []string{"owners", "0", "email"}, From: interface{}(nil), To: "user@odpf.io"},
					diff.Change{Type: "create", Path: []string{"owners", "1", "email"}, From: interface{}(nil), To: "meteor@odpf.io"},
				},
				UpdatedBy: r.users[1],
			},
		}

		assetVersions, err := r.repository.GetVersionHistory(r.ctx, asset.Filter{Size: 3}, astVersioning.ID)
		r.NoError(err)
		// making updatedby user time empty to make ast comparable
		for i := 0; i < len(assetVersions); i++ {
			assetVersions[i].UpdatedBy.CreatedAt = time.Time{}
			assetVersions[i].UpdatedBy.UpdatedAt = time.Time{}
			assetVersions[i].CreatedAt = time.Time{}
			assetVersions[i].UpdatedAt = time.Time{}
		}
		r.Equal(expectedAssetVersions, assetVersions)
	})

	r.Run("should return current version of an assets", func() {
		expectedLatestVersion := asset.Asset{
			ID:          astVersioning.ID,
			URN:         assetURN,
			Type:        "table",
			Service:     "bigquery",
			Description: "new description in v0.2",
			Data:        map[string]interface{}{"data1": float64(12345)},
			Labels:      map[string]string{"key1": "value1"},
			Version:     "0.5",
			UpdatedBy:   r.users[1],
		}

		ast, err := r.repository.GetByID(r.ctx, astVersioning.ID)
		// hard to get the internally generated user id, we exclude the owners from the assertion
		astOwners := ast.Owners
		ast.Owners = nil
		r.NoError(err)
		// making updatedby user time empty to make ast comparable
		ast.UpdatedBy.CreatedAt = time.Time{}
		ast.UpdatedBy.UpdatedAt = time.Time{}
		ast.CreatedAt = time.Time{}
		ast.UpdatedAt = time.Time{}
		r.Equal(expectedLatestVersion, ast)

		r.Len(astOwners, 2)
	})

	r.Run("should return current version of an assets with by version", func() {
		expectedLatestVersion := asset.Asset{
			ID:          astVersioning.ID,
			URN:         assetURN,
			Type:        "table",
			Service:     "bigquery",
			Description: "new description in v0.2",
			Data:        map[string]interface{}{"data1": float64(12345)},
			Labels:      map[string]string{"key1": "value1"},
			Version:     "0.5",
			UpdatedBy:   r.users[1],
		}

		ast, err := r.repository.GetByVersion(r.ctx, astVersioning.ID, "0.5")
		// hard to get the internally generated user id, we exclude the owners from the assertion
		astOwners := ast.Owners
		ast.Owners = nil
		r.NoError(err)
		// making updatedby user time empty to make ast comparable
		ast.UpdatedBy.CreatedAt = time.Time{}
		ast.UpdatedBy.UpdatedAt = time.Time{}
		ast.CreatedAt = time.Time{}
		ast.UpdatedAt = time.Time{}
		r.Equal(expectedLatestVersion, ast)

		r.Len(astOwners, 2)
	})

	r.Run("should return a specific version of an asset", func() {
		selectedVersion := "0.3"
		expectedAsset := asset.Asset{
			ID:          astVersioning.ID,
			URN:         assetURN,
			Type:        "table",
			Service:     "bigquery",
			Description: "new description in v0.2",
			Version:     "0.3",
			Changelog: diff.Changelog{
				diff.Change{Type: "create", Path: []string{"owners", "0", "email"}, From: interface{}(nil), To: "user@odpf.io"},
				diff.Change{Type: "create", Path: []string{"owners", "1", "email"}, From: interface{}(nil), To: "meteor@odpf.io"},
			},
			UpdatedBy: r.users[1],
		}
		expectedOwners := []user.User{
			{
				Email: "user@odpf.io",
			},
			{
				Email:    "meteor@odpf.io",
				Provider: "meteor",
			},
		}
		astVer, err := r.repository.GetByVersion(r.ctx, astVersioning.ID, selectedVersion)
		// hard to get the internally generated user id, we exclude the owners from the assertion
		astOwners := astVer.Owners
		astVer.Owners = nil
		r.Assert().NoError(err)
		// making updatedby user time empty to make ast comparable
		astVer.UpdatedBy.CreatedAt = time.Time{}
		astVer.UpdatedBy.UpdatedAt = time.Time{}
		astVer.CreatedAt = time.Time{}
		astVer.UpdatedAt = time.Time{}
		r.Assert().Equal(expectedAsset, astVer)

		for i := 0; i < len(astOwners); i++ {
			astOwners[i].ID = ""
		}
		r.Assert().Equal(expectedOwners, astOwners)
	})
}

func (r *AssetRepositoryTestSuite) TestUpsert() {
	r.Run("on insert", func() {
		r.Run("set ID to asset and version to base version", func() {
			ast := asset.Asset{
				URN:       "urn-u-1",
				Type:      "table",
				Service:   "bigquery",
				Version:   "0.1",
				UpdatedBy: r.users[0],
			}
			id, err := r.repository.Upsert(r.ctx, &ast)
			r.NoError(err)
			r.NotEmpty(id)
			ast.ID = id

			assetInDB, err := r.repository.GetByID(r.ctx, ast.ID)
			r.Require().NoError(err)
			r.NotEqual(time.Time{}, assetInDB.CreatedAt)
			r.NotEqual(time.Time{}, assetInDB.UpdatedAt)
			r.assertAsset(&ast, &assetInDB)
		})

		r.Run("should store owners if any", func() {
			ast := asset.Asset{
				URN:     "urn-u-3",
				Type:    "table",
				Service: "bigquery",
				Owners: []user.User{
					r.users[1],
					r.users[2],
				},
				UpdatedBy: r.users[0],
			}

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.Require().NotEmpty(id)
			ast.ID = id

			actual, err := r.repository.GetByID(r.ctx, ast.ID)
			r.NoError(err)

			r.Len(actual.Owners, len(ast.Owners))
			for i, owner := range actual.Owners {
				r.Equal(ast.Owners[i].ID, owner.ID)
			}
		})

		r.Run("should create owners as users if they do not exist yet", func() {
			ast := asset.Asset{
				URN:     "urn-u-3a",
				Type:    "table",
				Service: "bigquery",
				Owners: []user.User{
					{Email: "newuser@example.com", Provider: defaultProviderName},
				},
				UpdatedBy: r.users[0],
			}

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.NotEmpty(id)

			actual, err := r.repository.GetByID(r.ctx, id)
			r.NoError(err)

			r.Len(actual.Owners, len(ast.Owners))
			for i, owner := range actual.Owners {
				r.Equal(ast.Owners[i].Email, owner.Email)
				r.NotEmpty(id)
			}
		})
	})

	r.Run("on update", func() {
		r.Run("should not create nor updating the asset if asset is identical", func() {
			ast := asset.Asset{
				URN:       "urn-u-2",
				Type:      "table",
				Service:   "bigquery",
				UpdatedBy: r.users[0],
			}
			identicalAsset := ast

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.NotEmpty(id)
			ast.ID = id

			id, err = r.repository.Upsert(r.ctx, &identicalAsset)
			r.Require().NoError(err)
			r.NotEmpty(id)
			identicalAsset.ID = id

			r.Equal(ast.ID, identicalAsset.ID)
		})

		r.Run("should delete old owners if it does not exist on new asset", func() {
			ast := asset.Asset{
				URN:     "urn-u-4",
				Type:    "table",
				Service: "bigquery",
				Owners: []user.User{
					r.users[1],
					r.users[2],
				},
				UpdatedBy: r.users[0],
			}
			newAsset := ast
			newAsset.Owners = []user.User{
				r.users[2],
			}

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.NotEmpty(id)
			ast.ID = id

			id, err = r.repository.Upsert(r.ctx, &newAsset)
			r.Require().NoError(err)
			r.NotEmpty(id)
			newAsset.ID = id

			actual, err := r.repository.GetByID(r.ctx, ast.ID)
			r.NoError(err)
			r.Len(actual.Owners, len(newAsset.Owners))
			for i, owner := range actual.Owners {
				r.Equal(newAsset.Owners[i].ID, owner.ID)
			}
		})

		r.Run("should create new owners if it does not exist on old asset", func() {
			ast := asset.Asset{
				URN:     "urn-u-4",
				Type:    "table",
				Service: "bigquery",
				Owners: []user.User{
					r.users[1],
				},
				UpdatedBy: r.users[0],
			}
			newAsset := ast
			newAsset.Owners = []user.User{
				r.users[1],
				r.users[2],
			}

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.NotEmpty(id)
			ast.ID = id

			id, err = r.repository.Upsert(r.ctx, &newAsset)
			r.Require().NoError(err)
			r.NotEmpty(id)
			newAsset.ID = id

			actual, err := r.repository.GetByID(r.ctx, ast.ID)
			r.NoError(err)
			r.Len(actual.Owners, len(newAsset.Owners))
			for i, owner := range actual.Owners {
				r.Equal(newAsset.Owners[i].ID, owner.ID)
			}
		})

		r.Run("should create users from owners if owner emails do not exist yet", func() {
			ast := asset.Asset{
				URN:     "urn-u-4a",
				Type:    "table",
				Service: "bigquery",
				Owners: []user.User{
					r.users[1],
				},
				UpdatedBy: r.users[0],
			}
			newAsset := ast
			newAsset.Owners = []user.User{
				r.users[1],
				{Email: "newuser@example.com", Provider: defaultProviderName},
			}

			id, err := r.repository.Upsert(r.ctx, &ast)
			r.Require().NoError(err)
			r.NotEmpty(id)
			ast.ID = id

			id, err = r.repository.Upsert(r.ctx, &newAsset)
			r.Require().NoError(err)
			r.NotEmpty(id)
			newAsset.ID = id

			actual, err := r.repository.GetByID(r.ctx, ast.ID)
			r.NoError(err)
			r.Len(actual.Owners, len(newAsset.Owners))
			for i, owner := range actual.Owners {
				r.Equal(newAsset.Owners[i].Email, owner.Email)
				r.NotEmpty(id)
			}
		})
	})
}

func (r *AssetRepositoryTestSuite) TestDelete() {
	r.Run("return error from client if any", func() {
		err := r.repository.Delete(r.ctx, "invalid-uuid")
		r.Error(err)
		r.Contains(err.Error(), "invalid asset id: \"invalid-uuid\"")
	})

	r.Run("return NotFoundError if asset does not exist", func() {
		uuid := "2aabb450-f986-44e2-a6db-7996861d5004"
		err := r.repository.Delete(r.ctx, uuid)
		r.ErrorAs(err, &asset.NotFoundError{AssetID: uuid})
	})

	r.Run("should delete correct asset", func() {
		asset1 := asset.Asset{
			URN:       "urn-del-1",
			Type:      "table",
			Service:   "bigquery",
			UpdatedBy: user.User{ID: defaultAssetUpdaterUserID},
		}
		asset2 := asset.Asset{
			URN:       "urn-del-2",
			Type:      "topic",
			Service:   "kafka",
			Version:   asset.BaseVersion,
			UpdatedBy: user.User{ID: defaultAssetUpdaterUserID},
		}

		var err error
		id, err := r.repository.Upsert(r.ctx, &asset1)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		asset1.ID = id

		id, err = r.repository.Upsert(r.ctx, &asset2)
		r.Require().NoError(err)
		r.Require().NotEmpty(id)
		asset2.ID = id

		err = r.repository.Delete(r.ctx, asset1.ID)
		r.NoError(err)

		_, err = r.repository.GetByID(r.ctx, asset1.ID)
		r.ErrorAs(err, &asset.NotFoundError{AssetID: asset1.ID})

		asset2FromDB, err := r.repository.GetByID(r.ctx, asset2.ID)
		r.NoError(err)
		r.Equal(asset2.ID, asset2FromDB.ID)

		// cleanup
		err = r.repository.Delete(r.ctx, asset2.ID)
		r.NoError(err)
	})
}

func (r *AssetRepositoryTestSuite) assertAsset(expectedAsset *asset.Asset, actualAsset *asset.Asset) bool {
	// sanitize time to make the assets comparable
	expectedAsset.CreatedAt = time.Time{}
	expectedAsset.UpdatedAt = time.Time{}
	expectedAsset.UpdatedBy.CreatedAt = time.Time{}
	expectedAsset.UpdatedBy.UpdatedAt = time.Time{}

	actualAsset.CreatedAt = time.Time{}
	actualAsset.UpdatedAt = time.Time{}
	actualAsset.UpdatedBy.CreatedAt = time.Time{}
	actualAsset.UpdatedBy.UpdatedAt = time.Time{}

	return r.Equal(expectedAsset, actualAsset)
}

func TestAssetRepository(t *testing.T) {
	suite.Run(t, &AssetRepositoryTestSuite{})
}
