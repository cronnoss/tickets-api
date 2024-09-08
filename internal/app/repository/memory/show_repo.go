package memory

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

type mapShow map[int64]*domain.NewShowData

type StorageShow struct {
	data mapShow
	mu   sync.RWMutex
}

var GenID int64

func getNewIDSafe() int64 {
	return atomic.AddInt64(&GenID, 1)
}

func NewShowRepo() StorageShow {
	return StorageShow{data: make(mapShow), mu: sync.RWMutex{}}
}

// GetShows returns shows.
func (s *StorageShow) GetShows(_ context.Context) ([]domain.NewShowData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceS := []domain.NewShowData{}
	for _, v := range s.data {
		sliceS = append(sliceS, *v)
	}
	return sliceS, nil
}

// CreateShows creates shows.
func (s *StorageShow) CreateShows(_ context.Context, shows []domain.NewShowData) ([]domain.NewShowData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range shows {
		shows[i].ID = getNewIDSafe()
		s.data[shows[i].ID] = &shows[i]
	}
	return shows, nil
}

// CreateShow creates a show.
func (s *StorageShow) CreateShow(_ context.Context, show domain.NewShowData) (domain.NewShowData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	show.ID = getNewIDSafe()
	s.data[show.ID] = &show
	return show, nil
}
