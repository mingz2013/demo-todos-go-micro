package services

import (
	"github.com/micro/go-log"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/dao"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/datastore"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
	"strconv"
)

type TodoService struct {
}

func newTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Add(req *pb.AddReq) (err error) {
	todo := &pb.Todo{
		Id:      strconv.Itoa(datastore.GetKeyId()),
		Title:   req.Title,
		Content: req.Content,
	}

	err = dao.GetTodoDao().Add(todo)
	log.Log(todo, err)
	return
}

func (s *TodoService) Del(req *pb.DelReq) (err error) {

	err = dao.GetTodoDao().Del(req.Id)

	return
}

func (s *TodoService) Edit(req *pb.EditReq) (err error) {
	todo := &pb.Todo{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}
	err = dao.GetTodoDao().Update(todo)
	log.Log(todo, err)
	return
}

func (s *TodoService) Get(req *pb.DetailReq) (todo *pb.Todo, err error) {
	todo, err = dao.GetTodoDao().Get(req.Id)
	log.Log(todo, err)
	return
}

func (s *TodoService) GetAll() (todos []*pb.Todo, err error) {
	todos, err = dao.GetTodoDao().GetAll()
	log.Log(todos, err)
	return
}
