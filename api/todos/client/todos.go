package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	pb "github.com/mingz2013/demo-todos-go-micro/api/todos/proto/todos"
)

type exampleKey struct{}

// FromContext retrieves the client from the Context
func TodosFromContext(ctx context.Context) (pb.TodosService, bool) {
	c, ok := ctx.Value(exampleKey{}).(pb.TodosService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func TodosWrapper(service micro.Service) server.HandlerWrapper {
	client := pb.NewTodosService("go.micro.srv.todos", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
