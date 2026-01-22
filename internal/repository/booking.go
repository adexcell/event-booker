package repository

import (
	"context"
	"time"

	"github.com/adexcell/event-booker/internal/domain"
	"github.com/google/uuid"
)

type DTOBooking struct {
	EventID   uuid.UUID
	UserID    uuid.UUID
	Status    domain.BookingStatus
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (p *postgres) Create(ctx context.Context, b DTOBooking) error {
	const query = `
		insert into booking (booking_id, event_id, user_id, status, expires_at, created_at)
		values ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.ExecContext(ctx, query, b)
	return err
}
