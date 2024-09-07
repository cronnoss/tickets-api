package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

// ShowService is a show service.
type ShowService struct {
	repo ShowRepo
}

// NewShowService creates a new show service.
func NewShowService(repo ShowRepo) ShowService {
	return ShowService{
		repo: repo,
	}
}

// GetShows returns shows.
func (s ShowService) GetShows(ctx context.Context) ([]models.Show, error) {
	return s.repo.GetShows(ctx)
}

// CreateShows creates shows.
func (s ShowService) CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error) {
	return s.repo.CreateShows(ctx, shows)
}

// CreateShow creates a show.
func (s ShowService) CreateShow(ctx context.Context, show models.Show) (models.Show, error) {
	return s.repo.CreateShow(ctx, show)
}
