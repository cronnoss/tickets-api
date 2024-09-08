package services

import (
	"context"

	"github.com/cronnoss/tickets-api/internal/app/domain"
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
func (s ShowService) GetShows(ctx context.Context) ([]domain.Show, error) {
	return s.repo.GetShows(ctx)
}

// CreateShows creates shows.
func (s ShowService) CreateShows(ctx context.Context, shows []domain.Show) ([]domain.Show, error) {
	return s.repo.CreateShows(ctx, shows)
}

// CreateShow creates a show.
func (s ShowService) CreateShow(ctx context.Context, show domain.Show) (domain.Show, error) {
	return s.repo.CreateShow(ctx, show)
}
