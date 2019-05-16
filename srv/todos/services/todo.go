package services

import (
	"github.com/micro/go-log"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/dao"
	"github.com/mingz2013/demo-todos-go-micro/srv/todos/datastore"
	pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"
	"gopkg.in/mgo.v2"
	"strconv"
)

type TodoService struct {
}

func newTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Add(req *pb.AddReq) (pbErr *pb.Error) {
	todo := &pb.Todo{
		Id:      strconv.Itoa(datastore.GetKeyId()),
		Title:   req.Title,
		Content: req.Content,
	}

	err := dao.GetTodoDao().Add(todo)

	log.Log(todo, err)

	if err != nil {
		pbErr = &pb.Error{}
		pbErr.Detail = err.Error()
		pbErr.Code = 500
		return
	}

	return
}

func (s *TodoService) Del(req *pb.DelReq) (pbErr *pb.Error) {

	err := dao.GetTodoDao().Del(req.Id)
	log.Log(err)
	if err != nil {
		pbErr = &pb.Error{}

		if err == mgo.ErrNotFound {
			pbErr.Detail = err.Error()
			pbErr.Code = 404
			return
		}

		pbErr.Detail = err.Error()
		pbErr.Code = 500
		return
	}

	return
}

func (s *TodoService) Edit(req *pb.EditReq) (pbErr *pb.Error) {
	todo := &pb.Todo{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}
	err := dao.GetTodoDao().Update(todo)
	log.Log(todo, err)

	if err != nil {
		pbErr = &pb.Error{}

		if err == mgo.ErrNotFound {
			pbErr.Detail = err.Error()
			pbErr.Code = 404
			return
		}

		pbErr.Detail = err.Error()
		pbErr.Code = 500
		return
	}

	return
}

func (s *TodoService) Get(req *pb.DetailReq) (todo pb.Todo, pbErr *pb.Error) {
	todo, err := dao.GetTodoDao().Get(req.Id)
	log.Log(todo, err)

	if err != nil {
		pbErr = &pb.Error{}

		if err == mgo.ErrNotFound {
			pbErr.Detail = err.Error()
			pbErr.Code = 404
			return
		}

		pbErr.Detail = err.Error()
		pbErr.Code = 500
		return
	}

	return
}

func (s *TodoService) GetAll() (todos []*pb.Todo, pbErr *pb.Error) {
	todos, err := dao.GetTodoDao().GetAll()
	log.Log(todos, err)

	if err != nil {
		pbErr = &pb.Error{}

		if err == mgo.ErrNotFound {
			pbErr.Detail = err.Error()
			pbErr.Code = 404
			return
		}

		pbErr.Detail = err.Error()
		pbErr.Code = 500
		return
	}
	return
}
