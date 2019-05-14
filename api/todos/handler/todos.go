package handler

import (
	"context"
	"github.com/micro/go-log"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/mingz2013/demo-todos-go-micro/api/todos/client"
	pb "github.com/mingz2013/demo-todos-go-micro/api/todos/proto/todos"
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
func (todo *Todos) List(ctx context.Context, req *pb.ListReq, rsp *pb.ListResp) error {
	log.Log("Received Todos.List request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.list", "example client not found")
	}

	// make request
	response, err := todosClient.List(ctx, req)
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.list", err.Error())
	}

	//b, _ := json.Marshal(response)

	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todos) Add(ctx context.Context, req *pb.AddReq, rsp *pb.AddResp) error {
	log.Log("Received Todos.Add request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.add", "example client not found")
	}

	// make request
	response, err := todosClient.Add(ctx, req)
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

func (todo *Todos) Del(ctx context.Context, req *pb.DelReq, rsp *pb.DelResp) error {
	log.Log("Received Todos.Del request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.Del", "example client not found")
	}

	// make request
	response, err := todosClient.Del(ctx, req)
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

func (todo *Todos) Edit(ctx context.Context, req *pb.EditReq, rsp *pb.EditResp) error {
	log.Log("Received Todos.Edit request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.Edit", "example client not found")
	}

	// make request
	response, err := todosClient.Edit(ctx, req)
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.Edit", err.Error())
	}

	//b, _ := json.Marshal(response)

	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	rsp = response
	log.Log(rsp)
	return nil
}

func (todo *Todos) Detail(ctx context.Context, req *pb.DetailReq, rsp *pb.DetailResp) error {
	log.Log("Received Todos.Detail request")

	// extract the client from the context
	todosClient, ok := client.TodosFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", "example client not found")
	}

	// make request
	response, err := todosClient.Detail(ctx, req)
	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", err.Error())
	}

	//b, _ := json.Marshal(response)

	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	rsp = response
	log.Log(rsp)
	return nil
}
