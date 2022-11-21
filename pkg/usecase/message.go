package usecase

import (
	"github.com/chiachun0920/platform-api/pkg/dto"
	"github.com/chiachun0920/platform-api/pkg/dto/schema"
	"github.com/chiachun0920/platform-api/pkg/service"
)

func SaveMessage(repo service.MessageRepository, msg *dto.Message) error {
	err := repo.SaveMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func ListMessages(repo service.MessageRepository, customerId string) ([]*dto.Message, error) {
	messages, err := repo.ListMessages(customerId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func SendMessage(messaging service.Messaging, msg schema.MessagingRequest) error {
	return messaging.Send(msg)
}
