package handler

import (
	"encoding/json"
	"github.com/micro/go-log"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
)

func TodosList(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodosService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.List(r.Context(), &pb.ListReq{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"todos": rsp.Todos,
		"ref":   time.Now().UnixNano(),
	}
	log.Log(rsp)

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosEdit(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodosService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Edit(r.Context(), &pb.EditReq{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	response := map[string]interface{}{
		//"todos": rsp.Todos,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosAdd(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodosService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Add(r.Context(), &pb.AddReq{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	response := map[string]interface{}{
		//"todos": rsp.Todos,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosDetail(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodosService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Detail(r.Context(), &pb.DetailReq{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	response := map[string]interface{}{
		"todos": rsp.Todo,
		"ref":   time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosDel(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodosService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Del(r.Context(), &pb.DelReq{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	response := map[string]interface{}{
		//"todos": rsp.Todos,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
