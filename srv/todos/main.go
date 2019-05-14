package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/handler"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.todos"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterTodosHandler(service.Server(), new(handler.Todos))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.todos", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.todos", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
