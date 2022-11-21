package service

import "github.com/chiachun0920/platform-api/pkg/dto"

type MessageRepository interface {
	ListMessages(customerId string) ([]*dto.Message, error)
	SaveMessage(m *dto.Message) error
}

type CustomerRepository interface {
	UpsertCustomerProfile(profile *dto.CustomerProfile) error
}
