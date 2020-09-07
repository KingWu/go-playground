package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	// "os"
	"log"
	// "math/rand"
	"context"
	// "fmt"
	"kw101/go-playground/graph/generated"
	"kw101/go-playground/graph/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)



func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	conn := ctx.Value("db").(*pgxpool.Pool) 

	// sql, _, _ := sq.Insert("app.user").
	// 	Columns("name").
	// 	Values("user " + input.UserID).
	// 	ToSql()

	// // Create user
	// conn.Exec(context.Background(), sql)

	sql, args, _ := sq.Insert("app.todo").
		Columns("text").
    Values(input.Text).
		ToSql()

	log.Print(sql)
	log.Print(args)
	
	_, err := conn.Exec(context.Background(), sql, args)
	log.Print(err)

	todo := &model.Todo{
		Text:   input.Text,
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
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
