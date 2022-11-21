package lineapi

import (
	"github.com/chiachun0920/platform-api/pkg/dto"
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

func (c *LineAPI) GetProfile(customerId string) (*dto.CustomerProfile, error) {
	profile, err := c.client.GetProfile(customerId).Do()
	if err != nil {
		return nil, err
	}

	customerProfile := &dto.CustomerProfile{
		UserID:      profile.UserID,
		DisplayName: profile.DisplayName,
		PictureURL:  profile.PictureURL,
	}
	return customerProfile, nil
}

func (c *LineAPI) createDefaultMsg(text string) *linebot.TextMessage {
	message := linebot.NewTextMessage(text)
	return message
}
