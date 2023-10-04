package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/yigithancolak/custmate/server"
	"github.com/yigithancolak/custmate/util"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	config, err := util.LoadConfig(".", appEnv, "env")
	if err != nil {
		log.Fatalf("error while loading config: %v", err)
	}

	srv, err := server.NewServer(config)
	if err != nil {
		log.Fatalf("error while creating server: %v", err)
	}

	srv.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	srv.Router.Handle("/query", srv.GraphQL)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, srv.Router))
}
