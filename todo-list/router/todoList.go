package router

import (
	"github.com/gin-gonic/gin"
	"mygo/todo-list/controller"
)

func RegisterRouter(r *gin.Engine) {
	todoList := controller.TodoList{}

	r.Use(todoList.Cors)
	r.GET("/list", todoList.List)
}
