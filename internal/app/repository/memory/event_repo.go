package memory

import (
	"context"
	"sync"

	"github.com/cronnoss/tickets-api/internal/app/domain"
)

type mapEvent map[int64]*domain.Event

type StorageEvent struct {
	data mapEvent
	mu   sync.RWMutex
}

func NewEventRepo() StorageEvent {
	return StorageEvent{data: make(mapEvent), mu: sync.RWMutex{}}
}

// GetEvents returns events.
func (s *StorageEvent) GetEvents(_ context.Context) ([]domain.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceS := []domain.Event{}
	for _, v := range s.data {
		sliceS = append(sliceS, *v)
	}
	return sliceS, nil
}

// CreateEvents creates events.
func (s *StorageEvent) CreateEvents(_ context.Context, events []domain.Event) ([]domain.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range events {
		events[i].ID = getNewIDSafe()
		s.data[events[i].ID] = &events[i]
	}
	return events, nil
}

// CreateEvent creates a event.
func (s *StorageEvent) CreateEvent(_ context.Context, event domain.Event) (domain.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event.ID = getNewIDSafe()
	s.data[event.ID] = &event
	return event, nil
}
