package memory

import (
	"context"
	"sync"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

type mapPlace map[int64]*domain.NewPlaceData

type StoragePlace struct {
	data mapPlace
	mu   sync.RWMutex
}

func NewPlaceRepo() StoragePlace {
	return StoragePlace{data: make(mapPlace), mu: sync.RWMutex{}}
}

// GetPlaces returns places.
func (s *StoragePlace) GetPlaces(_ context.Context) ([]domain.NewPlaceData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceS := []domain.NewPlaceData{}
	for _, v := range s.data {
		sliceS = append(sliceS, *v)
	}
	return sliceS, nil
}

// CreatePlaces creates places.
func (s *StoragePlace) CreatePlaces(_ context.Context, places []domain.NewPlaceData) ([]domain.NewPlaceData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range places {
		places[i].ID = getNewIDSafe()
		s.data[places[i].ID] = &places[i]
	}
	return places, nil
}

// CreatePlace creates a place.
func (s *StoragePlace) CreatePlace(_ context.Context, place domain.NewPlaceData) (domain.NewPlaceData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	place.ID = getNewIDSafe()
	s.data[place.ID] = &place
	return place, nil
}
