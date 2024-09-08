//go:generate mockery

package httpserver

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

// ShowService is a show service.
type ShowService interface {
	GetShows(ctx context.Context) ([]domain.Show, error)
	CreateShow(ctx context.Context, show domain.Show) (domain.Show, error)
}

// EventService is a event service.
type EventService interface {
	GetEvents(ctx context.Context) ([]domain.Event, error)
	CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error)
}

type PlaceService interface {
	GetPlaces(ctx context.Context) ([]domain.Place, error)
	CreatePlace(ctx context.Context, place domain.Place) (domain.Place, error)
}
