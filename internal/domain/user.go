package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Surname      string
	Email        string
	PasswordHash string
	TelegramID   int
	CreatedAt    time.Time
	DeletedAt    time.Time
}
