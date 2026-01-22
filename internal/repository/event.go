package repository

import (
	"context"
	"fmt"

	"github.com/adexcell/event-booker/internal/domain"
	"github.com/google/uuid"
)

func (p *postgres) AcquireSlot(ctx context.Context, eventID uuid.UUID) error {
	const query = `update events set available_slots = available_slots - 1
	where event_id = $1 and available_slots > 0`

	res, err := p.db.ExecContext(ctx, query, eventID)
	if err != nil {
		return fmt.Errorf("failed to acquire slot: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return domain.ErrNoAvailableSlots
	}

	return nil
}

func (p *postgres) ReleaseSlot(ctx context.Context, eventID uuid.UUID) error {
	const query = `update events set available_slots = available_slots + 1
	where event_id = $1 and available_slots < total_slots`

	_, err := p.db.ExecContext(ctx, query, eventID)
	return err
}
