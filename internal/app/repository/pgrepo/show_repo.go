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

type ShowRepo struct {
	DB *pg.DB
}

func NewShowRepo(db *pg.DB) *ShowRepo {
	return &ShowRepo{
		DB: db,
	}
}

// GetShows returns shows.
func (r ShowRepo) GetShows(ctx context.Context) ([]models.Show, error) {
	var shows []models.Show
	err := r.DB.NewSelect().Model(&shows).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get shows: %w", err)
	}

	modelsShows := make([]models.Show, 0, len(shows))
	for _, show := range shows { // nolint: gosimple
		modelsShows = append(modelsShows, show)
	}

	return modelsShows, nil
}

// CreateShows creates shows.
func (r ShowRepo) CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error) {
	dbShows := make([]models.Show, 0, len(shows))
	var insertedShows []models.Show
	err := r.DB.NewInsert().Model(&dbShows).Ignore().Returning("*").Scan(ctx, &insertedShows)
	if err != nil {
		return insertedShows, nil // nolint: nilerr
	}

	modelsShows := make([]models.Show, 0, len(insertedShows))
	return modelsShows, nil
}

// CreateShow creates a show.
func (r ShowRepo) CreateShow(ctx context.Context, show models.Show) (models.Show, error) {
	dbShow := show
	var insertedShow models.Show
	err := r.DB.NewInsert().Model(&dbShow).Ignore().Returning("*").Scan(ctx, &insertedShow)
	if err != nil {
		return show, nil // nolint: nilerr
	}

	return insertedShow, nil
}
