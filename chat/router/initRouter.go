package router

import (
	"github.com/gin-gonic/gin"
	"mygo/chat/controller"
)

func RegisterRouter(r *gin.Engine) {
	chat := controller.Chat{}
	r.POST("/chat", chat.DoChat)
}
