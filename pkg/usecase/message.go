package usecase

import (
	"github.com/chiachun0920/platform-api/pkg/dto"
	"github.com/chiachun0920/platform-api/pkg/service"
)

func SaveMessage(repo service.MessageRepository, msg *dto.Message) error {
	err := repo.SaveMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
