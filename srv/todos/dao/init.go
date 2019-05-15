package dao

import "github.com/mingz2013/demo-todos-go-micro/srv/todos/datastore"

var todoDao *TodoDao

func GetTodoDao() *TodoDao {
	if todoDao == nil {
		session := datastore.GetSession()
		todoDao = newTodoDao(session)
	}

	return todoDao

}
