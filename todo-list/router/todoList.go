package router

import (
	"github.com/gin-gonic/gin"
	"mygo/todo-list/controller"
	"mygo/todo-list/service"
)

func RegisterRouter(r *gin.Engine, list *service.TodoList) {
	todoList := controller.TodoList{
		TodoList: list,
	}

	r.Use(todoList.Cors)
	r.GET("/list", todoList.List)
}
