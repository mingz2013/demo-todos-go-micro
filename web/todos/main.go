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

	// register call handler
	service.HandleFunc("/todo/list", handler.TodosList)
	service.HandleFunc("/todo/edit", handler.TodosEdit)
	service.HandleFunc("/todo/detail", handler.TodosDetail)
	service.HandleFunc("/todo/add", handler.TodosAdd)
	service.HandleFunc("/todo/del", handler.TodosDel)

	service.Handle("/static", http.FileServer(http.Dir("html/static")))

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	service.HandleFunc("*", http.NotFound)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
