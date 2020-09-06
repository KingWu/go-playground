package main

import (
	"kw101/go-playground/graph"
	"kw101/go-playground/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

const defaultPort = "8080"

func buildMiddleware(middlewareController *negroni.Negroni) {
	middlewareController.Use(negroni.NewRecovery())
	middlewareController.Use(negroni.NewLogger())
}

func createRouter() *httprouter.Router{
	router := httprouter.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.HandlerFunc("GET", "/", playground.Handler("GraphQL playground", "/query"))
	router.Handler("POST", "/query", srv)
	return router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	middlewareController := negroni.New()
	buildMiddleware(middlewareController)

	middlewareController.UseHandler(createRouter())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, middlewareController))
}
