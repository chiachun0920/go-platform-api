package controller

import (
	"net/http"

	"github.com/chiachun0920/platform-api/pkg/dto"
	"github.com/chiachun0920/platform-api/pkg/dto/schema"
	"github.com/chiachun0920/platform-api/pkg/service"
	"github.com/chiachun0920/platform-api/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageRepo service.MessageRepository
}

func NewMessageController(repo service.MessageRepository) *MessageController {
	return &MessageController{messageRepo: repo}
}

func (controller *MessageController) WebhookLine(c *gin.Context) {
	var req schema.LineWebhookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	text := req.Events[0].MessageData.Text
	if err := usecase.SaveMessage(controller.messageRepo, &dto.Message{
		Text: text,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
