package main

import (
	"github.com/micro/go-log"
	"time"

	"github.com/micro/go-micro"
	"github.com/mingz2013/demo-todos-go-micro/api/todos/client"
	"github.com/mingz2013/demo-todos-go-micro/api/todos/handler"

	pb "github.com/mingz2013/demo-todos-go-micro/api/todos/proto/todo"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.todos"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.TodoWrapper(service)),
	)

	// Register Handler
	pb.RegisterTodoHandler(service.Server(), new(handler.Todo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
