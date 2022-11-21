package service

import "github.com/chiachun0920/platform-api/pkg/dto"

type MessageRepository interface {
	ListMessages() ([]*dto.Message, error)
	SaveMessage(m *dto.Message) error
}
