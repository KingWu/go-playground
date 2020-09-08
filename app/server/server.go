package server

import (
	"kw101/go-playground/config"
	"context"
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
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

func createDatabase(databaseUrl string) *pgxpool.Pool {
	log.Printf("**** Start Connect DB")
	dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Printf("**** Connected DB")
	return dbpool
}

func Run() {
	pool := createDatabase(config.DatabaseUrl())

	middleware := BuildMiddleware(
		[]string{"http://foo.com", "http://foo.com:8080"},
		config.Env() != config.Production,
	)

	// Inject var into context
	ctx := context.WithValue(context.Background(), "db", pool)
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
