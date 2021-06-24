package main

import (
	"order/graph"
	"order/graph/generated"
	"order/database"
	"log"
	"net/http"
	"os"
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4002"

// HeaderMiddleware - add header as context
func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()
		profileID := r.Header.Get("Context")
		ctx := context.WithValue(c, graph.Key{}, profileID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
        ORDER: db.Order,
				APRA: db.Apra,
    }}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
