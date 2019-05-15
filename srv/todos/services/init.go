package services

var todoService *TodoService

func GetTodosService() *TodoService {
	if todoService == nil {
		todoService = newTodoService()
	}

	return todoService
}
