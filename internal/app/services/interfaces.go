package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

type ShowRepo interface {
	GetShows(ctx context.Context) ([]models.Show, error)
	CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error)
	CreateShow(ctx context.Context, show models.Show) (models.Show, error)
}

type EventRepo interface {
	GetEvents(ctx context.Context) ([]models.Event, error)
	CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
}

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
	CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) (models.Place, error)
}
