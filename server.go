package main

import (
	"log"
	"time"
	"kw101/go-playground/server"
	"os"
	"gopkg.in/tylerb/graceful.v1"
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
	log.Printf("Start server listening on :%s", port)
	graceful.Run(":" + port, 10*time.Second, middleware)
}
