package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

type ShowRepo interface {
	GetShows(ctx context.Context) ([]domain.Show, error)
	CreateShows(ctx context.Context, shows []domain.Show) ([]domain.Show, error)
	CreateShow(ctx context.Context, show domain.Show) (domain.Show, error)
}

type EventRepo interface {
	GetEvents(ctx context.Context) ([]domain.Event, error)
	CreateEvents(ctx context.Context, events []domain.Event) ([]domain.Event, error)
	CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error)
}

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]domain.Place, error)
	CreatePlaces(ctx context.Context, places []domain.Place) ([]domain.Place, error)
	CreatePlace(ctx context.Context, place domain.Place) (domain.Place, error)
}
