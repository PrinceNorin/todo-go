package main

import (
	_ "github.com/PrinceNorin/todo-go/docs"
	"github.com/PrinceNorin/todo-go/handler"
	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Todo Application
// @description This is a todo list management application
// @version 1.0
// @host localhost:8081
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler

	v1 := e.Group("/api/v1")
	{
		todo := handler.NewHandler()
		v1.GET("/todos", todo.FindTodos)
		v1.POST("/todos", todo.CreateTodo)
		v1.GET("/todos/:id", todo.GetTodo)
		v1.PUT("/todos/:id", todo.UpdateTodo)
		v1.DELETE("/todos/:id", todo.DeleteTodo)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8081"))
}
