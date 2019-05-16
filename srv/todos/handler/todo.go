package handler

import (
	"context"
	"github.com/micro/go-log"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/services"

	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
)

type Todo struct {
}

func (todo *Todo) Add(ctx context.Context, req *pb.AddReq, rsp *pb.AddResp) error {
	log.Log("Todos.Add")

	detail, err := services.GetTodosService().Add(req)

	if err != nil {
		rsp.Success = false
		rsp.Error = err
		return nil
	}

	rsp.Success = true
	rsp.Error = nil
	rsp.Todo = detail
	return nil
}

func (todo *Todo) Del(ctx context.Context, req *pb.DelReq, rsp *pb.DelResp) error {
	log.Log("Todos.Del:-> ", req)

	err := services.GetTodosService().Del(req)

	if err != nil {
		rsp.Success = false
		rsp.Error = err
		return nil
	}

	rsp.Success = true
	rsp.Error = nil
	return nil
}

func (todo *Todo) Edit(ctx context.Context, req *pb.EditReq, rsp *pb.EditResp) error {
	log.Log("Todos.Edit:->", req)

	err := services.GetTodosService().Edit(req)
	log.Log("Todos.Edit:-> err: ", err)
	if err != nil {
		rsp.Success = false
		rsp.Error = err
		return nil
	}

	rsp.Success = true
	rsp.Error = nil
	return nil
}

func (todo *Todo) List(ctx context.Context, req *pb.ListReq, rsp *pb.ListResp) error {
	log.Log("Todos.List")

	todos, err := services.GetTodosService().GetAll()

	if err != nil {
		rsp.Success = false
		rsp.Error = err
		return nil
	}

	rsp.Success = true
	rsp.Error = nil
	rsp.Todos = todos

	log.Log(rsp)

	return nil
}

func (todo *Todo) Detail(ctx context.Context, req *pb.DetailReq, rsp *pb.DetailResp) error {
	log.Log("Todos.Detail--> ", " req:", req, " id: ", req.Id, "rsp: ", rsp)
	detail, err := services.GetTodosService().Get(req)

	if err != nil {
		rsp.Success = false
		rsp.Error = err
		return nil
	}

	rsp.Success = true
	rsp.Error = nil
	rsp.Todo = detail

	log.Log("rsp: ", rsp)
	return nil
}
