package validators

import (
	"errors"

	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
	"github.com/google/uuid"
)

func SaveMessageValidator(message entities.Message) error {
	if message.UserID == uuid.Nil {
		return errors.New("UserID should not be empty")
	}
	if message.Content == "" {
		return errors.New("message content should not be empty")
	}
	if message.Sender == "" {
		return errors.New("message sender should not be empty")
	}
	if message.Receiver == "" {
		return errors.New("message reciver should not be empty")
	}
	return nil
}