package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/yigithancolak/custmate/directives"
	"github.com/yigithancolak/custmate/graph"
	"github.com/yigithancolak/custmate/middleware"
	"github.com/yigithancolak/custmate/postgresdb"
	"github.com/yigithancolak/custmate/token"
	"github.com/yigithancolak/custmate/util"
)

type Server struct {
	Config   *util.Config
	GraphQL  *handler.Server
	Resolver *graph.Resolver
	Router   *chi.Mux
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

	router := chi.NewRouter()
	router.Use(middleware.Middleware(store.Organizations))

	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	return &Server{
		Config:   config,
		GraphQL:  gqlServer,
		Resolver: resolver,
		Router:   router,
	}, nil
}
