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

type Todo struct {
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
func (todo *Todo) List(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.List request")

	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
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

func (todo *Todo) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Add request")

	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
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

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)
	//rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todo) Del(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Del request")

	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
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

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)
	//rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todo) Edit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Edit request")

	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
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

func (todo *Todo) Detail(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Todos.Detail request", "req", req, "post", req.Post, "body", req.Body, "body", req.GetBody(), "post", req.GetPost(), "method", req.Method, "header", req.Header, "url", req.Url, "path", req.Path)
	log.Log(json.Marshal(req))
	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", "example client not found")
	}

	// make request
	detailReq := &pb.DetailReq{
		Id: extractValue(req.Post["id"]),
	}
	log.Log("detailReq", detailReq)
	response, err := todosClient.Detail(ctx, detailReq)

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
