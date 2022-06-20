package main

import (
	"ks-aggregator-rates/graph"
	"ks-aggregator-rates/graph/generated"
	"ks-aggregator-rates/internal/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.SetupDb()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: db,
				},
			}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
