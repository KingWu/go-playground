package main

import (
	"kw101/go-playground/app/server"
)

func main() {
	server.Run()
}

// import (
// 	"kw101/go-playground/api/graphql"
// 	"log"
// 	"net/http"

// 	graphqlGo "github.com/graph-gophers/graphql-go"
// 	"github.com/graph-gophers/graphql-go/relay"
// )

// func main() {
// 	schema := graphqlGo.MustParseSchema(graphql.Schema, &graphql.Resolver{})
// 	http.Handle("/api/v1", &relay.Handler{Schema: schema})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }