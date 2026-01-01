package event

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Surname      string
	Login        string
	Email        string
	PasswordHash string
	TelegramID   int
}
