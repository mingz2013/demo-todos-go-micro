package handler

import (
	"context"
	"encoding/json"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/errors"
	"github.com/mingz2013/demo-todos-go-micro/api/todos/client"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
)

type Todos struct {
}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Call is called by the API as /todos/example/call with post body {"name": "foo"}
func (todo *Todos) List(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.List request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.list", "example client not found")
	}

	// make request
	response, err := todosClient.List(ctx, &pb.ListReq{})
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.list", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	log.Log(rsp)
	return nil
}

func (todo *Todos) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Add request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.add", "example client not found")
	}

	// make request
	response, err := todosClient.Add(ctx, &pb.AddReq{
		//Name: extractValue(req.Post["name"]),
		Title:   extractValue(req.Post["title"]),
		Content: extractValue(req.Post["content"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.add", err.Error())
	}

	//b, _ := json.Marshal(response)

	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todos) Del(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Del request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.Del", "example client not found")
	}

	// make request
	response, err := todosClient.Del(ctx, &pb.DelReq{
		Id: extractValue(req.Post["id"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.Del", err.Error())
	}

	//b, _ := json.Marshal(response)

	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todos) Edit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Edit request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.Edit", "example client not found")
	}

	// make request
	response, err := todosClient.Edit(ctx, &pb.EditReq{
		Id:      extractValue(req.Post["id"]),
		Title:   extractValue(req.Post["title"]),
		Content: extractValue(req.Post["content"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.Edit", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)
	//rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todos) Detail(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Detail request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", "example client not found")
	}

	// make request
	response, err := todosClient.Detail(ctx, &pb.DetailReq{
		Id: extractValue(req.Post["id"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)
	//rsp = response
	log.Log(rsp)
	return nil
}
