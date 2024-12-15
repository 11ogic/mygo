package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/todo-list/service"
	"net/http"
)

type TodoList struct {
	service.TodoList
}

func (l *TodoList) Cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func (l *TodoList) List(ctx *gin.Context) {
	result := l.GetList()

	fmt.Println(result)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": result,
	})
}
