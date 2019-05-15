package main

import (
	"github.com/micro/go-log"
	"net/http"

	"github.com/micro/go-web"
	"github.com/mingz2013/demo-todos-go-micro/web/todos/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.todos"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	service.Handle("/static", http.FileServer(http.Dir("html/static")))

	// register call handler
	service.HandleFunc("/todos/list", handler.TodosList)
	service.HandleFunc("/todos/edit", handler.TodosEdit)
	service.HandleFunc("/todos/detail", handler.TodosDetail)
	service.HandleFunc("/todos/add", handler.TodosAdd)
	service.HandleFunc("/todos/del", handler.TodosDel)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
