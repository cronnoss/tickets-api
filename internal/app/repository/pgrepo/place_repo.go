package pgrepo

import (
	"context"
	"fmt"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
	"github.com/cronnoss/tickets-api/internal/pkg/pg"
)

type PlaceRepo struct {
	DB *pg.DB
}

func NewPlaceRepo(db *pg.DB) *PlaceRepo {
	return &PlaceRepo{
		DB: db,
	}
}

// GetPlaces returns places.
func (r PlaceRepo) GetPlaces(ctx context.Context) ([]models.Place, error) {
	var places []models.Place
	err := r.DB.NewSelect().Model(&places).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get places: %w", err)
	}

	modelsPlaces := make([]models.Place, 0, len(places))
	for _, place := range places { // nolint: gosimple
		modelsPlaces = append(modelsPlaces, place)
	}

	return modelsPlaces, nil
}

// CreatePlaces creates places.
func (r PlaceRepo) CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error) {
	dbPlaces := make([]models.Place, 0, len(places))

	var insertedPlaces []models.Place
	err := r.DB.NewInsert().Model(&dbPlaces).Returning("*").Scan(ctx, &insertedPlaces)
	if err != nil {
		return insertedPlaces, nil // nolint: nilerr
	}

	modelsPlaces := make([]models.Place, 0, len(insertedPlaces))
	return modelsPlaces, nil
}

// CreatePlace creates a place.
func (r PlaceRepo) CreatePlace(ctx context.Context, place models.Place) (models.Place, error) {
	dbPlace := place
	err := r.DB.NewInsert().Model(&dbPlace).Returning("*").Scan(ctx, &dbPlace)
	if err != nil {
		return place, nil // nolint: nilerr
	}

	return dbPlace, nil
}
