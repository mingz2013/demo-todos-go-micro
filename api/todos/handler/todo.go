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
	listReq := &pb.ListReq{}

	err := json.Unmarshal([]byte(req.Body), listReq)
	if err != nil {
		log.Log(err)
		return errors.InternalServerError("go.micro.api.todos.todos.list", "body json error")
	}

	log.Log("listReq", listReq)
	response, err := todosClient.List(ctx, listReq)

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
	log.Log("Received Todos.Add request", req)

	// extract the client from the context
	todosClient, ok := client.TodoFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.todos.todos.add", "example client not found")
	}

	// make request
	addReq := &pb.AddReq{
		//Name: extractValue(req.Post["name"]),
		//Title:   extractValue(req.Post["title"]),
		//Content: extractValue(req.Post["content"]),
	}

	err := json.Unmarshal([]byte(req.Body), addReq)
	if err != nil {
		log.Log(err)
		return errors.InternalServerError("go.micro.api.todos.todos.add", "body json error")
	}

	log.Log("addReq", addReq)
	response, err := todosClient.Add(ctx, addReq)
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
	delReq := &pb.DelReq{}

	err := json.Unmarshal([]byte(req.Body), delReq)
	if err != nil {
		log.Log(err)
		return errors.InternalServerError("go.micro.api.todos.todos.del", "body json error")
	}

	log.Log("delReq", delReq)
	response, err := todosClient.Del(ctx, delReq)

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
	reqSrv := &pb.EditReq{}

	err := json.Unmarshal([]byte(req.Body), reqSrv)
	if err != nil {
		log.Log(err)
		return errors.InternalServerError("go.micro.api.todos.todos.Edit", "body json error")
	}

	log.Log("reqSrv", reqSrv)
	response, err := todosClient.Edit(ctx, reqSrv)
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
	reqSrv := &pb.DetailReq{}

	err := json.Unmarshal([]byte(req.Body), reqSrv)
	if err != nil {
		log.Log(err)
		return errors.InternalServerError("go.micro.api.todos.todos.Detail", "body json error")
	}

	log.Log("reqSrv: ", reqSrv)
	response, err := todosClient.Detail(ctx, reqSrv)

	if err != nil {
		return errors.InternalServerError("go.micro.api.todos.todos.detail", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)
	//rsp = response
	log.Log("rsp: ", rsp)
	return nil
}
