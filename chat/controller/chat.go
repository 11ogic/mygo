package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"mygo/chat/model"
	"mygo/chat/util"
	"net/http"
)

type Chat struct {
}

func (c *Chat) DoChat(ctx *gin.Context) {
	body := &model.Chat{}

	if err := ctx.ShouldBindJSON(body); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	prompt := util.CreatePrompt()

	var data = map[string]any{
		"text": body.Text,
	}

	msg, _ := prompt.FormatMessages(data)

	content := []llms.MessageContent{
		llms.TextParts(msg[0].GetType(), msg[0].GetContent()),
		llms.TextParts(msg[1].GetType(), msg[1].GetContent()),
	}

	llm := util.CreateOllama(ctx, "qwen")

	res, err := llm.GenerateContent(context.Background(), content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
