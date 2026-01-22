package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID uuid.UUID
	Title string
	Description string
	TotalSlots int
	AvailableSlots int
	BookTimeout time.Duration
	EventAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}


