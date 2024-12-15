package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/todo-list/db"
	"mygo/todo-list/model"
	"mygo/todo-list/router"
)

func main() {
	r := gin.Default()

	database := db.Connect("practice_todo_list")

	err := database.AutoMigrate(&model.TodoList{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("auto migrate database success...")

	router.RegisterRouter(r)

	err = r.Run()

	if err != nil {
		return
	}
}
