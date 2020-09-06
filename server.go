package main

import (
	"kw101/go-playground/config"
	"net/http"
	"log"
	"time"
	"kw101/go-playground/server"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	middleware := server.BuildMiddleware()
	middleware.UseHandler(server.CreateRouter())
	// Run server
	log.Printf("Environment: %s", config.Env())

	port := config.Port()
	if config.Env() == config.Development {
		// Http
		log.Printf("Start http server listening on :%s", port)
		graceful.Run(":" + port, 10*time.Second, middleware)
	} else {
		// Https 
		log.Printf("Start https server listening on :%s", port)
		server := &http.Server{
			Addr:    ":" + port,
			Handler: middleware,
		}
		graceful.ListenAndServeTLS(server, "cert/localhost.crt", "cert/localhost.key", 10*time.Second)
	}
}
