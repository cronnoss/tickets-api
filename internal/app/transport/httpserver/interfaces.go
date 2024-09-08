//go:generate mockery

package httpserver

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

// ShowService is a show service.
type ShowService interface {
	GetShows(ctx context.Context) ([]domain.NewShowData, error)
	CreateShow(ctx context.Context, show domain.NewShowData) (domain.NewShowData, error)
}

// EventService is a event service.
type EventService interface {
	GetEvents(ctx context.Context) ([]domain.NewEventData, error)
	CreateEvent(ctx context.Context, event domain.NewEventData) (domain.NewEventData, error)
}

type PlaceService interface {
	GetPlaces(ctx context.Context) ([]domain.NewPlaceData, error)
	CreatePlace(ctx context.Context, place domain.NewPlaceData) (domain.NewPlaceData, error)
}
