package controller

import (
	"net/http"
	"time"

	"github.com/chiachun0920/platform-api/pkg/dto"
	"github.com/chiachun0920/platform-api/pkg/dto/schema"
	"github.com/chiachun0920/platform-api/pkg/service"
	"github.com/chiachun0920/platform-api/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageRepo  service.MessageRepository
	customerRepo service.CustomerRepository
	messaging    service.Messaging
}

func NewMessageController(
	msgRepo service.MessageRepository,
	customerRepo service.CustomerRepository,
	messaging service.Messaging,
) *MessageController {
	return &MessageController{
		messageRepo:  msgRepo,
		customerRepo: customerRepo,
		messaging:    messaging,
	}
}

func convertTime(milseconds int64) time.Time {
	return time.Unix(0, milseconds*int64(time.Millisecond))
}

func (controller *MessageController) WebhookLine(c *gin.Context) {
	var req schema.LineWebhookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	event := req.Events[0]

	if err := usecase.SaveMessage(controller.messageRepo, &dto.Message{
		MessageType: event.MessageData.Type,
		Text:        event.MessageData.Text,
		Sender:      event.SourceData.UserID,
		CreatedAt:   convertTime(event.Timestamp),
		UpdatedAt:   convertTime(event.Timestamp),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	customerId := req.Events[0].SourceData.UserID
	if err := usecase.UpdateCustomerProfile(
		controller.customerRepo,
		controller.messaging,
		customerId,
	); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
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
