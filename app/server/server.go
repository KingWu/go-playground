package server

import (
	"github.com/graph-gophers/dataloader"
	"fmt"
	"kw101/go-playground/app/database/sql"
	"kw101/go-playground/graph"
	"kw101/go-playground/config"
	"context"
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"log"
	"time"
	"gopkg.in/tylerb/graceful.v1"
	"kw101/go-playground/graph/model"
)

func connectDatabasePool(databaseUrl string) *pgxpool.Pool {
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
	pool := connectDatabasePool(config.DatabaseUrl())

	middleware := BuildMiddleware(
		[]string{"http://foo.com", "http://foo.com:8080"},
		config.Env() != config.Production,
	)
	
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		log.Print("Fetching User2: todo id")
		log.Print(keys.Keys())
		sql, args, _ := sql.GetUsers(keys.Keys())
		log.Print(sql)
		log.Print(args)
		rows, err := pool.Query(context.Background(), sql, args...)
		log.Print(err)
		var users []*dataloader.Result
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			users = append(users, &dataloader.Result{
				Data: &model.User{
					ID: fmt.Sprintf("%d", id),
					Name: name,
				},
			})
		}
		log.Print(users)
		return users
	}
	
	// Inject var into context
	// cache := &dataloader.NoCache{}
	resolver := &graph.Resolver{ DB: pool,
		UserLoader: dataloader.NewBatchedLoader(batchFn, dataloader.WithClearCacheOnBatch()),
	}
	middleware.UseHandler(CreateRouter(resolver))

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
