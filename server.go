package main

import (
	"github.com/rs/cors"
	"kw101/go-playground/graph"
	"kw101/go-playground/graph/generated"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

const defaultPort = "8080"

func setUpMiddleware(middlewareController *negroni.Negroni) {
	middlewareController.Use(negroni.NewRecovery())
	middlewareController.Use(negroni.NewLogger())

	// Handle Cors
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
	})
	middlewareController.Use(cors)
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
	setUpMiddleware(middlewareController)

	// Add Router
	middlewareController.UseHandler(createRouter())
	middlewareController.Run(":"+port)
}
