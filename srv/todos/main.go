package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/datastore"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/handler"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/subscriber"
	"os"
	"time"
)

const defaultHost = "localhost:27017"

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := datastore.CreateSession(host)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	datastore.CreateRedisClient()
	defer datastore.GetRedisClient().Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.todos"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterTodoInterfaceHandler(service.Server(), new(handler.Todo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.todos", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.todos", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
