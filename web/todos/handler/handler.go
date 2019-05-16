package handler

import (
	"encoding/json"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
	"net/http"
)

func TodosList(w http.ResponseWriter, r *http.Request) {
	log.Log("TodosList.....")
	defer log.Log("todolist....", w)
	// decode the incoming request as json
	//var request map[string]interface{}
	var request pb.ListReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Log(err)
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodoInterfaceService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.List(r.Context(), &request)
	if err != nil {
		log.Log(err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Log("TodosList...", "rsp", rsp)

	// we want to augment the response
	//response := map[string]interface{}{
	//	"todos": rsp.Todos,
	//	"ref":   time.Now().UnixNano(),
	//}
	//log.Log(rsp)

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		log.Log(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosEdit(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//var request map[string]interface{}
	var request pb.EditReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodoInterfaceService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Edit(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	//response := map[string]interface{}{
	//	//"todos": rsp.Todos,
	//	"ref": time.Now().UnixNano(),
	//}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosAdd(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//var request map[string]interface{}
	var request pb.AddReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodoInterfaceService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Add(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	//response := map[string]interface{}{
	//	//"todos": rsp.Todos,
	//	"ref": time.Now().UnixNano(),
	//}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosDetail(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//var request map[string]interface{}
	var request pb.DetailReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodoInterfaceService("go.micro.srv.todos", client.DefaultClient)

	rsp, err := todosClient.Detail(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	//response := map[string]interface{}{
	//	"todos": rsp.Todo,
	//	"ref":   time.Now().UnixNano(),
	//}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodosDel(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//var request map[string]interface{}
	var request pb.DelReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	todosClient := pb.NewTodoInterfaceService("go.micro.srv.todos", client.DefaultClient)
	rsp, err := todosClient.Del(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Log(rsp)

	// we want to augment the response
	//response := map[string]interface{}{
	//	//"todos": rsp.Todos,
	//	"ref": time.Now().UnixNano(),
	//}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
