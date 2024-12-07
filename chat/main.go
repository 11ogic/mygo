package main

import (
	"github.com/gin-gonic/gin"
	"mygo/chat/middleware"
	"mygo/chat/router"
)

func main() {
	r := gin.Default()

	r.Use(middleware.Cors())

	router.RegisterRouter(r)

	err := r.Run()

	if err != nil {
		return
	}
}
