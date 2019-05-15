package services

import (
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/dao"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todos"
)

type TodoService struct {
}

func newTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Add(req *pb.AddReq) (err error) {
	todo := &pb.Todo{
		Id:      "1",
		Title:   req.Title,
		Content: req.Content,
	}

	err = dao.GetTodoDao().Add(todo)

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

	return
}

func (s *TodoService) Get(req *pb.DetailReq) (todo *pb.Todo, err error) {
	todo, err = dao.GetTodoDao().Get(req.Id)
	return
}

func (s *TodoService) GetAll() (todos []*pb.Todo, err error) {
	todos, err = dao.GetTodoDao().GetAll()
	return
}
