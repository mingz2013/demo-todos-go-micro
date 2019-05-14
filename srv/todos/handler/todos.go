package handler

import (
	"context"
	"github.com/micro/go-log"

	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
)

type Todos struct{}

//// Call is a single request handler called via client.Call or the generated client code
//func (e *Todos) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
//	log.Log("Received Example.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}
//
//// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *Todos) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
//	log.Logf("Received Example.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Logf("Responding: %d", i)
//		if err := stream.Send(&pb.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *Todos) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Logf("Got ping %v", req.Stroke)
//		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}

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
