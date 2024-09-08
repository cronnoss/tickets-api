package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

type ShowRepo interface {
	GetShows(ctx context.Context) ([]domain.NewShowData, error)
	CreateShows(ctx context.Context, shows []domain.NewShowData) ([]domain.NewShowData, error)
	CreateShow(ctx context.Context, show domain.NewShowData) (domain.NewShowData, error)
}

type EventRepo interface {
	GetEvents(ctx context.Context) ([]domain.NewEventData, error)
	CreateEvents(ctx context.Context, events []domain.NewEventData) ([]domain.NewEventData, error)
	CreateEvent(ctx context.Context, event domain.NewEventData) (domain.NewEventData, error)
}

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]domain.NewPlaceData, error)
	CreatePlaces(ctx context.Context, places []domain.NewPlaceData) ([]domain.NewPlaceData, error)
	CreatePlace(ctx context.Context, place domain.NewPlaceData) (domain.NewPlaceData, error)
}
