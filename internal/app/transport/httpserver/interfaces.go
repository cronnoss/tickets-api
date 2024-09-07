//go:generate mockery

package httpserver

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

// ShowService is a show service.
type ShowService interface {
	GetShows(ctx context.Context) ([]models.Show, error)
	CreateShow(ctx context.Context, show models.Show) (models.Show, error)
}

// EventService is a event service.
type EventService interface {
	GetEvents(ctx context.Context) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
}

type PlaceService interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) (models.Place, error)
}
