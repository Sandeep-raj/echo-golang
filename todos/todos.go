package todos

import "github.com/labstack/echo/v4"

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TodosGroupHandler(todo *echo.Group) {
	todo.GET("/", GetTodos)
	todo.GET("/userid/:userid", GetTodosByUserId)
	todo.GET("/id/:id", GetTodosById)
}
