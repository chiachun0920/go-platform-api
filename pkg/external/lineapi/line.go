package lineapi

import (
	"github.com/chiachun0920/platform-api/pkg/dto/schema"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineAPI struct {
	client *linebot.Client
}

func NewLineAPI(secret, token string) *LineAPI {
	client, _ := linebot.New(secret, token)
	return &LineAPI{client: client}
}

func (c *LineAPI) Send(msg schema.MessagingRequest) error {
	var messages []linebot.SendingMessage
	messages = append(messages, c.createDefaultMsg(msg.Text))

	_, err := c.client.PushMessage(msg.To, messages...).Do()
	return err
}

func (c *LineAPI) createDefaultMsg(text string) *linebot.TextMessage {
	message := linebot.NewTextMessage(text)
	return message
}
