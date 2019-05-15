package handler

import (
	"context"
	"github.com/micro/go-log"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/services"

	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
)

type Todos struct {
}

func (todo *Todos) Add(ctx context.Context, req *pb.AddReq, rsp *pb.AddResp) error {
	log.Log("Todos.Add")

	err := services.GetTodosService().Add(req)

	if err != nil {
		rsp.Success = false
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()
		return err
	}

	rsp.Success = true
	rsp.Error = nil
	return nil
}

func (todo *Todos) Del(ctx context.Context, req *pb.DelReq, rsp *pb.DelResp) error {
	log.Log("Todos.Del")

	err := services.GetTodosService().Del(req)

	if err != nil {
		rsp.Success = false
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()
		return err
	}

	rsp.Success = true
	rsp.Error = nil
	return nil
}

func (todo *Todos) Edit(ctx context.Context, req *pb.EditReq, rsp *pb.EditResp) error {
	log.Log("Todos.Edit")

	err := services.GetTodosService().Edit(req)

	if err != nil {
		rsp.Success = false
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()
		return err
	}

	rsp.Success = true
	rsp.Error = nil
	return nil
}

func (todo *Todos) List(ctx context.Context, req *pb.ListReq, rsp *pb.ListResp) error {
	log.Log("Todos.List")

	todos, err := services.GetTodosService().GetAll()

	if err != nil {
		rsp.Success = false
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()
		return err
	}

	rsp.Success = true
	rsp.Error = nil
	rsp.Todos = todos

	return nil

	log.Log(rsp)
	return nil
}

func (todo *Todos) Detail(ctx context.Context, req *pb.DetailReq, rsp *pb.DetailResp) error {
	log.Log("Todos.Detail")
	detail, err := services.GetTodosService().Get(req)

	if err != nil {
		rsp.Success = false
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()
		return err
	}

	rsp.Success = true
	rsp.Error = nil
	rsp.Todo = detail

	return nil

	log.Log(rsp)
	return nil
}
