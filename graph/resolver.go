package graph

import (
	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *pgxpool.Pool
	UserLoader *dataloader.Loader
}
