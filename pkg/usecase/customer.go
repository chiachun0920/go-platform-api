package usecase

import (
	"github.com/chiachun0920/platform-api/pkg/service"
)

func UpdateCustomerProfile(repo service.CustomerRepository, messaging service.Messaging, customerId string) error {
	profile, err := messaging.GetProfile(customerId)
	if err != nil {
		return err
	}

	err = repo.UpsertCustomerProfile(profile)
	if err != nil {
		return err
	}
	return nil
}
