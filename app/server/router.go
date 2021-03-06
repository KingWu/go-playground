package server

import (
	"github.com/99designs/gqlgen/graphql"
	"context"
	"net/http"
	"kw101/go-playground/graph"
	"kw101/go-playground/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/julienschmidt/httprouter"
)

func CreateRouter(resolver *graph.Resolver) *httprouter.Router{
	router := httprouter.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		// TODO: Unknown error and report bugs to trackers
		return graphql.DefaultRecover(ctx, err)
	})

	router.HandlerFunc("GET", "/", playground.Handler("GraphQL playground", "/query"))
	router.HandlerFunc("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"hello\": \"world\"}"))
	})
	router.Handler("POST", "/query", srv)
	return router
}
