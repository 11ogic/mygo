package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mygo/todo-list/db"
	"mygo/todo-list/model"
	"mygo/todo-list/router"
	"mygo/todo-list/service"
)

func main() {
	r := gin.Default()

	database := db.Connect("practice_todo_list")

	err := database.AutoMigrate(&model.TodoList{})

	s := &service.TodoList{
		DB: database,
	}

	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	fmt.Println("auto migrate database success...")
	router.RegisterRouter(r, s)

	err = r.Run()

	if err != nil {
		return
	}
}
