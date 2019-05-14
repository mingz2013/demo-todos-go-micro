package handler

import (
	"context"
	"github.com/micro/go-log"

	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
)

type Todos struct{}

func (todo *Todos) Add(ctx context.Context, req *pb.AddReq, rsp *pb.AddResp) error {
	log.Log("Todos.Add")
	rsp.Error = nil
	rsp.Success = true
	return nil
}

func (todo *Todos) Del(ctx context.Context, req *pb.DelReq, rsp *pb.DelResp) error {
	log.Log("Todos.Del")
	rsp.Error = nil
	rsp.Success = true
	return nil
}

func (todo *Todos) Edit(ctx context.Context, req *pb.EditReq, rsp *pb.EditResp) error {
	log.Log("Todos.Edit")
	rsp.Error = nil
	rsp.Success = true

	return nil
}

func (todo *Todos) List(ctx context.Context, req *pb.ListReq, rsp *pb.ListResp) error {
	log.Log("Todos.List")
	rsp.Todos = []*pb.Todo{}
	rsp.Success = true
	rsp.Error = nil
	log.Log(rsp)
	return nil
}

func (todo *Todos) Detail(ctx context.Context, req *pb.DetailReq, rsp *pb.DetailResp) error {
	log.Log("Todos.Detail")
	rsp.Error = nil
	rsp.Success = true
	rsp.Todo = &pb.Todo{}
	return nil
}
