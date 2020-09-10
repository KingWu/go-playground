package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	// "github.com/vektah/gqlparser/v2/gqlerror"
	// "github.com/99designs/gqlgen/graphql"
	"context"
	"fmt"
	dbSql "kw101/go-playground/app/database/sql"
	"kw101/go-playground/graph/generated"
	"kw101/go-playground/graph/model"
	"log"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	conn := r.DB

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
		Text:   input.Text,
		UserID: fmt.Sprintf("%d", userID),
		// User: &model.User{ID: fmt.Sprintf("%d", userID), Name: input.Name},
	}
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, limit *int) ([]*model.Todo, error) {
	var qlimit uint = 10
	if limit != nil  && *limit > 0 {
		qlimit = uint(*limit)
	}

	conn := r.DB
	sql, args, _ := dbSql.ListToDo(qlimit)
	rows, _ := conn.Query(context.Background(), sql, args...)

	var todos []*model.Todo
	for rows.Next() {
		var id int
		var text string
		var userID int
		err := rows.Scan(&id, &text, &userID)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &model.Todo{
			ID:     strconv.Itoa(id),
			Text:   text,
			UserID: fmt.Sprintf("%d", userID),
		})
	}

	// Try error msg
	// graphql.AddError(ctx, &gqlerror.Error{
	// 	Message: "Error message",
	// 	Extensions: map[string]interface{}{
	// 		"code": "100001",
	// 	},
	// })

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	log.Printf("resolver User: todo id [%s]", obj.UserID)
	return r.UserLoader.Load(obj.UserID)
	// conn := r.DB
	// sql, args, _ := dbSql.GetUser(obj.UserID)

	// var id int
	// var name string
	// conn.QueryRow(context.Background(), sql, args...).Scan(&id, &name)
	// return &model.User{
	// 	ID:   fmt.Sprintf("%d", id),
	// 	Name: name,
	// }, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
