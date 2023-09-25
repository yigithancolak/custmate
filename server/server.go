package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/yigithancolak/custmate/graph"
	"github.com/yigithancolak/custmate/postgresdb"
	"github.com/yigithancolak/custmate/token"
	"github.com/yigithancolak/custmate/util"
)

type Server struct {
	Config   *util.Config
	GraphQL  *handler.Server
	Resolver *graph.Resolver
}

func NewServer(config *util.Config) (*Server, error) {
	db, err := postgresdb.NewDB(config)
	if err != nil {
		return nil, err
	}

	jwtMaker, err := token.NewJWTMaker(config)
	if err != nil {
		return nil, err
	}
	store := postgresdb.NewStore(db, jwtMaker)
	resolver := graph.NewResolver(store)
	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return &Server{
		Config:   config,
		GraphQL:  gqlServer,
		Resolver: resolver,
	}, nil
}
