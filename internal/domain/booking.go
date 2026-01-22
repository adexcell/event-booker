package domain

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus int

const (
	StatusPending BookingStatus = iota // 0: ожидает оплаты
	StatusConfirmed
	StatusCancelled
)

type Booking struct {
	ID        uuid.UUID
	EventID   uuid.UUID
	UserID    uuid.UUID
	Status    BookingStatus
	ExpiresAt time.Time
	CreatedAt time.Time
}
