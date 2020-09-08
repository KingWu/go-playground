package graph

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"kw101/go-playground/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	DB *pgxpool.Pool
}
