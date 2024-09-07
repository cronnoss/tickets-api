package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cronnoss/tickets-api/internal/app/domain"
	"github.com/cronnoss/tickets-api/internal/app/repository/models"
	"github.com/cronnoss/tickets-api/internal/pkg/pg"
)

type EventRepo struct {
	DB *pg.DB
}

func NewEventRepo(db *pg.DB) *EventRepo {
	return &EventRepo{
		DB: db,
	}
}

// GetEvents returns events.
func (r EventRepo) GetEvents(ctx context.Context) ([]models.Event, error) {
	var events []models.Event
	err := r.DB.NewSelect().Model(&events).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	modelsEvents := make([]models.Event, 0, len(events))
	for _, event := range events { // nolint: gosimple
		modelsEvents = append(modelsEvents, event)
	}

	return modelsEvents, nil
}

// CreateEvents creates events.
func (r EventRepo) CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error) {
	dbEvents := make([]models.Event, 0, len(events))

	var insertedEvents []models.Event
	err := r.DB.NewInsert().Model(&dbEvents).Returning("*").Scan(ctx, &insertedEvents)
	if err != nil {
		return insertedEvents, nil // nolint: nilerr
	}

	modelsEvents := make([]models.Event, 0, len(insertedEvents))
	return modelsEvents, nil
}

// CreateEvent creates a event.
func (r EventRepo) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	dbEvent := event
	err := r.DB.NewInsert().Model(&dbEvent).Returning("*").Scan(ctx, &dbEvent)
	if err != nil {
		return event, nil // nolint: nilerr
	}

	return dbEvent, nil
}
