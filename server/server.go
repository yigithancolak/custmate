package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph"
	"github.com/yigithancolak/custmate/postgresdb"
	"github.com/yigithancolak/custmate/util"
)

type Server struct {
	Config   *util.Config
	DB       *sqlx.DB
	GraphQL  *handler.Server
	Resolver *graph.Resolver
}

func NewServer(config *util.Config) (*Server, error) {
	db, err := postgresdb.NewDB(config)
	if err != nil {
		return nil, err
	}

	resolver := graph.NewResolver(db)
	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return &Server{
		Config:   config,
		DB:       db,
		GraphQL:  gqlServer,
		Resolver: resolver,
	}, nil
}
