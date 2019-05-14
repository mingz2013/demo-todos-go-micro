package subscriber

import "context"
import pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
import "github.com/micro/go-log"

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *pb.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *pb.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
