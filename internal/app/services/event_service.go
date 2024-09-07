package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

// EventService is a event service.
type EventService struct {
	repo EventRepo
}

// NewEventService creates a new event service.
func NewEventService(repo EventRepo) EventService {
	return EventService{
		repo: repo,
	}
}

// GetEvents returns events.
func (s EventService) GetEvents(ctx context.Context) ([]models.Event, error) {
	return s.repo.GetEvents(ctx)
}

// CreateEvents creates events.
func (s EventService) CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error) {
	return s.repo.CreateEvents(ctx, events)
}

// CreateEvent creates a event.
func (s EventService) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	return s.repo.CreateEvent(ctx, event)
}
