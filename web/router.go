package web

import (
	"github.com/gorilla/mux"
	"github.com/odpf/columbus/models"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger                  logrus.FieldLogger
	TypeRepository          models.TypeRepository
	RecordRepositoryFactory models.RecordRepositoryFactory
	RecordSearcher          models.RecordSearcher
	LineageProvider         LineageProvider
	Middlewares             []mux.MiddlewareFunc
}

func NewRouter(config Config) *mux.Router {
	router := mux.NewRouter()

	for _, middleware := range config.Middlewares {
		router.Use(middleware)
	}

	setupRoutes(router, config)

	return router
}

func setupRoutes(router *mux.Router, config Config) {
	typeHandler := NewTypeHandler(
		config.Logger.WithField("reporter", "type-handler"),
		config.TypeRepository,
		config.RecordRepositoryFactory,
	)
	searchHandler := NewSearchHandler(
		config.Logger.WithField("reporter", "search-handler"),
		config.RecordSearcher,
		config.TypeRepository,
	)
	lineageHandler := NewLineageHandler(
		config.Logger.WithField("reporter", "lineage-handler"),
		config.LineageProvider,
	)

	router.PathPrefix("/ping").Handler(NewHeartbeatHandler())
	router.PathPrefix("/v1/types").Handler(typeHandler)
	router.PathPrefix("/v1/entities").Handler(typeHandler) // For backward compatibility
	router.PathPrefix("/v1/search").Handler(searchHandler)
	router.PathPrefix("/v1/lineage").Handler(lineageHandler)
}
