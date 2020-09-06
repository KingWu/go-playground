package server

import (
	"golang.org/x/net/context"
	"kw101/go-playground/config"
	"net/http"
	"log"
	"time"
	"gopkg.in/tylerb/graceful.v1"
)

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func Run() {
	middleware := BuildMiddleware(
		[]string{"http://foo.com", "http://foo.com:8080"},
		config.Env() != config.Production,
	)

	// Inject var into context
	ctx := context.WithValue(context.Background(), "db", "test")
	middleware.UseHandler(&ContextInjector{ctx, CreateRouter()})

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
