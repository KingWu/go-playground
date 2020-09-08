package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"fmt"
	dbSql "kw101/go-playground/app/database/sql"
	"context"
	"kw101/go-playground/graph/generated"
	"kw101/go-playground/graph/model"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	conn := ctx.Value("db").(*pgxpool.Pool)

	// Create User
	sql, args, _ := dbSql.CreateUser(input.Name)
	log.Print(sql)
	log.Print(args)
	
	userID := 0
	err := conn.QueryRow(context.Background(), sql, args...).Scan(&userID)
	log.Printf("user id: %d", userID)

	// Create ToDo
	sql, args, _ = dbSql.CreateToDo(userID, input.Text)

	log.Print(sql)
	log.Print(args)

	_, err = conn.Exec(context.Background(), sql, args...)
	log.Print(err)

	todo := &model.Todo{
		Text: input.Text,
		User: &model.User{ID: fmt.Sprintf("%d", userID), Name: input.Name},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
