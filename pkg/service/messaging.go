package service

import (
	"github.com/chiachun0920/platform-api/pkg/dto"
	"github.com/chiachun0920/platform-api/pkg/dto/schema"
)

type Messaging interface {
	Send(event schema.MessagingRequest) error
	GetProfile(customerId string) (*dto.CustomerProfile, error)
}
