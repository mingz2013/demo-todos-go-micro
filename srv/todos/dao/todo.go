package dao

import (
	"github.com/micro/go-log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

import pb "github.com/mingz2013/demo-todos-go-micro/srv/todos/proto/todo"

const dbName = "todos"
const todoCollection = "todo"

type TodoDao struct {
	session *mgo.Session
}

func newTodoDao(session *mgo.Session) *TodoDao {
	return &TodoDao{session: session}
}

func (dao *TodoDao) Add(todo *pb.Todo) (err error) {
	err = dao.collection().Insert(todo)
	return
}

func (dao *TodoDao) Del(id string) (err error) {
	log.Log("Del: ", "id: ", id)
	err = dao.collection().Remove(bson.M{"id": id})
	log.Log("Del: err: ", err)
	return
}

func (dao *TodoDao) Update(todo *pb.Todo) (err error) {
	log.Log("Update ", "todo: ", todo)
	err = dao.collection().Update(bson.M{"id": todo.Id}, bson.M{"$set": todo})
	log.Log("Update: err: ", err)
	return
}

func (dao *TodoDao) Get(id string) (todo pb.Todo, err error) {
	log.Log("Get ", "id: ", id)
	err = dao.collection().Find(bson.M{"id": id}).One(&todo)
	log.Log("Get: err: ", err)
	return
}

func (dao *TodoDao) GetAll() (todos []*pb.Todo, err error) {
	err = dao.collection().Find(nil).All(&todos)
	return
}

func (dao *TodoDao) Close() {
	dao.session.Close()
}

func (dao *TodoDao) collection() *mgo.Collection {
	return dao.session.DB(dbName).C(todoCollection)
}
