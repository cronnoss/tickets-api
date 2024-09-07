package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

// PlaceService is a place service.
type PlaceService struct {
	repo PlaceRepo
}

// NewPlaceService creates a new place service.
func NewPlaceService(repo PlaceRepo) PlaceService {
	return PlaceService{
		repo: repo,
	}
}

// GetPlaces returns places.
func (s PlaceService) GetPlaces(ctx context.Context) ([]models.Place, error) {
	return s.repo.GetPlaces(ctx)
}

// CreatePlaces creates places.
func (s PlaceService) CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error) {
	return s.repo.CreatePlaces(ctx, places)
}

// CreatePlace creates a place.
func (s PlaceService) CreatePlace(ctx context.Context, place models.Place) (models.Place, error) {
	return s.repo.CreatePlace(ctx, place)
}
