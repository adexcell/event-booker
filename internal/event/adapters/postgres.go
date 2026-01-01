package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/zlog"
)

type eventPostgres struct {
	db *dbpg.DB
}

func NewEventPostgres(db *dbpg.DB) *eventPostgres {
	return &eventPostgres{db: db}
}

func (p *eventPostgres) AcquireSlot(ctx context.Context, eventID uuid.UUID) error {
	const query = `update events set available_slots = available_slots - 1
	where event_id = $1 and available_slots > 0`

	res, err := p.db.ExecContext(ctx, query, eventID)
	if err != nil {
		zlog.Logger.Error().Err(err).Msgf("failed to acquire slot")
		return ErrfailedToacquier
	}
}