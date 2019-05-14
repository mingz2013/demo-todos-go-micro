// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/todos/todos.proto

package go_micro_api_todos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Todos service

type TodosService interface {
	Add(ctx context.Context, in *AddReq, opts ...client.CallOption) (*AddResp, error)
	Del(ctx context.Context, in *DelReq, opts ...client.CallOption) (*DelResp, error)
	Edit(ctx context.Context, in *EditReq, opts ...client.CallOption) (*EditResp, error)
	List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListResp, error)
	Detail(ctx context.Context, in *DetailReq, opts ...client.CallOption) (*DetailResp, error)
}

type todosService struct {
	c    client.Client
	name string
}

func NewTodosService(name string, c client.Client) TodosService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.todos"
	}
	return &todosService{
		c:    c,
		name: name,
	}
}

func (c *todosService) Add(ctx context.Context, in *AddReq, opts ...client.CallOption) (*AddResp, error) {
	req := c.c.NewRequest(c.name, "Todos.Add", in)
	out := new(AddResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosService) Del(ctx context.Context, in *DelReq, opts ...client.CallOption) (*DelResp, error) {
	req := c.c.NewRequest(c.name, "Todos.Del", in)
	out := new(DelResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosService) Edit(ctx context.Context, in *EditReq, opts ...client.CallOption) (*EditResp, error) {
	req := c.c.NewRequest(c.name, "Todos.Edit", in)
	out := new(EditResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosService) List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListResp, error) {
	req := c.c.NewRequest(c.name, "Todos.List", in)
	out := new(ListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosService) Detail(ctx context.Context, in *DetailReq, opts ...client.CallOption) (*DetailResp, error) {
	req := c.c.NewRequest(c.name, "Todos.Detail", in)
	out := new(DetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Todos service

type TodosHandler interface {
	Add(context.Context, *AddReq, *AddResp) error
	Del(context.Context, *DelReq, *DelResp) error
	Edit(context.Context, *EditReq, *EditResp) error
	List(context.Context, *ListReq, *ListResp) error
	Detail(context.Context, *DetailReq, *DetailResp) error
}

func RegisterTodosHandler(s server.Server, hdlr TodosHandler, opts ...server.HandlerOption) error {
	type todos interface {
		Add(ctx context.Context, in *AddReq, out *AddResp) error
		Del(ctx context.Context, in *DelReq, out *DelResp) error
		Edit(ctx context.Context, in *EditReq, out *EditResp) error
		List(ctx context.Context, in *ListReq, out *ListResp) error
		Detail(ctx context.Context, in *DetailReq, out *DetailResp) error
	}
	type Todos struct {
		todos
	}
	h := &todosHandler{hdlr}
	return s.Handle(s.NewHandler(&Todos{h}, opts...))
}

type todosHandler struct {
	TodosHandler
}

func (h *todosHandler) Add(ctx context.Context, in *AddReq, out *AddResp) error {
	return h.TodosHandler.Add(ctx, in, out)
}

func (h *todosHandler) Del(ctx context.Context, in *DelReq, out *DelResp) error {
	return h.TodosHandler.Del(ctx, in, out)
}

func (h *todosHandler) Edit(ctx context.Context, in *EditReq, out *EditResp) error {
	return h.TodosHandler.Edit(ctx, in, out)
}

func (h *todosHandler) List(ctx context.Context, in *ListReq, out *ListResp) error {
	return h.TodosHandler.List(ctx, in, out)
}

func (h *todosHandler) Detail(ctx context.Context, in *DetailReq, out *DetailResp) error {
	return h.TodosHandler.Detail(ctx, in, out)
}
