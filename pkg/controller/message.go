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
	messaging   service.Messaging
}

func NewMessageController(repo service.MessageRepository, messaging service.Messaging) *MessageController {
	return &MessageController{
		messageRepo: repo,
		messaging:   messaging,
	}
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

func (controller *MessageController) SendMessage(c *gin.Context) {
	var req schema.MessagingRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	if err := usecase.SendMessage(controller.messaging, req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
