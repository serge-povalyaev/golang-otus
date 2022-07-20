package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	"sync"
	"time"
)

type MemoryRepository struct {
	events map[uuid.UUID]*Event
	mu     *sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		events: make(map[uuid.UUID]*Event),
		mu:     new(sync.RWMutex),
	}
}

func (eventRepository *MemoryRepository) GetEvent(userID int, eventID uuid.UUID) (*Event, error) {
	eventRepository.mu.Lock()
	defer eventRepository.mu.Unlock()

	foundEvent, ok := eventRepository.events[eventID]
	if ok && foundEvent.UserID == userID {
		return foundEvent, nil
	}

	return nil, ErrEventNotFound
}

func (eventRepository *MemoryRepository) AddEvent(event *Event) error {
	_, err := eventRepository.findEvent(event.ID)
	if errors.Is(err, ErrEventNotFound) {
		eventRepository.mu.Lock()
		defer eventRepository.mu.Unlock()

		eventRepository.events[event.ID] = event

		return nil
	}

	return ErrEventExists
}

func (eventRepository *MemoryRepository) UpdateEvent(userID int, event *Event) error {
	_, err := eventRepository.GetEvent(userID, event.ID)
	if err != nil {
		return err
	}

	eventRepository.mu.Lock()
	defer eventRepository.mu.Unlock()

	eventRepository.events[event.ID] = event

	return nil
}

func (eventRepository *MemoryRepository) RemoveEvent(userID int, event *Event) error {
	_, err := eventRepository.GetEvent(userID, event.ID)
	if err != nil {
		return err
	}

	eventRepository.mu.Lock()
	defer eventRepository.mu.Unlock()

	delete(eventRepository.events, event.ID)

	return nil
}

func (eventRepository *MemoryRepository) GetUserEvents(userID int, dateStart, dateFinish time.Time) ([]*Event, error) {
	eventRepository.mu.Lock()
	defer eventRepository.mu.Unlock()

	events := make([]*Event, 0)

	for _, event := range eventRepository.events {
		if event.UserID == userID &&
			(event.DateStart.After(dateStart) || event.DateStart.Equal(dateStart)) &&
			(event.DateFinish.Before(dateFinish) || event.DateFinish.Equal(dateFinish)) {
			events = append(events, event)
		}
	}

	return events, nil
}

func (eventRepository *MemoryRepository) GetUserEventsByDay(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfDay())
}

func (eventRepository *MemoryRepository) GetUserEventsByWeek(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfWeek())
}

func (eventRepository *MemoryRepository) GetUserEventsByMonth(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfMonth())
}

func (eventRepository *MemoryRepository) findEvent(eventID uuid.UUID) (*Event, error) {
	eventRepository.mu.Lock()
	defer eventRepository.mu.Unlock()

	foundEvent, ok := eventRepository.events[eventID]
	if ok {
		return foundEvent, nil
	}

	return nil, ErrEventNotFound
}
