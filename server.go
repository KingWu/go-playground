package main

import (
	"kw101/go-playground/server"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	middleware := server.BuildMiddleware()
	middleware.UseHandler(server.CreateRouter())
	// Run server
	middleware.Run(":"+port)
}
