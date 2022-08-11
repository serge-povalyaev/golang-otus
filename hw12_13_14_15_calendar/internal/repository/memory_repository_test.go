package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetEvent(t *testing.T) {
	repo := NewMemoryRepository()
	event := createNewEvent()
	_ = repo.AddEvent(event)

	foundEvent, err := repo.GetEvent(1, event.ID)
	require.NoError(t, err)
	require.Equal(t, event, foundEvent)

	foundEvent, err = repo.GetEvent(2, event.ID)
	require.Error(t, err)
	require.Nil(t, foundEvent)

	foundEvent, err = repo.GetEvent(1, uuid.New())
	require.Error(t, err)
	require.Nil(t, foundEvent)
}

func TestAddEvent(t *testing.T) {
	repo := NewMemoryRepository()
	event := createNewEvent()
	err := repo.AddEvent(event)
	foundEvent, err := repo.findEvent(event.ID)

	require.NoError(t, err)
	require.Equal(t, foundEvent, event)

	err = repo.AddEvent(event)
	require.Error(t, err)
}

func TestUpdateEvent(t *testing.T) {
	repo := NewMemoryRepository()
	event := createNewEvent()
	_ = repo.AddEvent(event)

	err := repo.UpdateEvent(2, event)
	require.Error(t, err)

	event.Description = "Проектная работа"
	err = repo.UpdateEvent(1, event)
	require.NoError(t, err)

	foundEvent, err := repo.findEvent(event.ID)
	require.Equal(t, foundEvent.Description, "Проектная работа")

	event.ID = uuid.New()
	err = repo.UpdateEvent(1, event)
	require.Error(t, err)
}

func TestRemoveEvent(t *testing.T) {
	repo := NewMemoryRepository()
	event := createNewEvent()
	_ = repo.AddEvent(event)

	err := repo.RemoveEvent(2, event)
	require.Error(t, err)

	err = repo.RemoveEvent(1, event)
	require.NoError(t, err)

	err = repo.RemoveEvent(1, event)
	require.Error(t, err)

	foundEvent, err := repo.findEvent(event.ID)
	require.Error(t, err)
	require.Nil(t, foundEvent)
}

func TestGetUserEvents(t *testing.T) {
	repo := getRepositoryWithEvents()
	from, _ := time.Parse("2006-01-02", "2022-01-01")

	to1, _ := time.Parse("2006-01-02", "2023-01-01")
	events, _ := repo.GetUserEvents(1, from, to1)
	require.Equal(t, 4, len(events))

	to2, _ := time.Parse("2006-01-02", "2022-07-16")
	events, _ = repo.GetUserEvents(1, from, to2)
	require.Equal(t, 1, len(events))

	to3, _ := time.Parse("2006-01-02", "2022-01-16")
	events, _ = repo.GetUserEvents(1, from, to3)
	require.Equal(t, 0, len(events))
}

func TestGetUserEventsByDay(t *testing.T) {
	repo := getRepositoryWithEvents()

	from1, _ := time.Parse("2006-01-02", "2022-07-15")
	events, _ := repo.GetUserEventsByDay(1, from1)
	require.Equal(t, 1, len(events))

	from2, _ := time.Parse("2006-01-02", "2022-07-17")
	events, _ = repo.GetUserEventsByDay(1, from2)
	require.Equal(t, 0, len(events))
}

func TestGetUserEventsByWeek(t *testing.T) {
	repo := getRepositoryWithEvents()

	from1, _ := time.Parse("2006-01-02", "2022-07-15")
	events, _ := repo.GetUserEventsByWeek(1, from1)
	require.Equal(t, 2, len(events))

	from2, _ := time.Parse("2006-01-02", "2022-07-17")
	events, _ = repo.GetUserEventsByWeek(1, from2)
	require.Equal(t, 0, len(events))
}

func TestGetUserEventsByMonth(t *testing.T) {
	repo := getRepositoryWithEvents()

	from1, _ := time.Parse("2006-01-02", "2022-07-15")
	events, _ := repo.GetUserEventsByMonth(1, from1)
	require.Equal(t, 3, len(events))

	from2, _ := time.Parse("2006-01-02", "2022-07-22")
	events, _ = repo.GetUserEventsByMonth(1, from2)
	require.Equal(t, 1, len(events))

	from3, _ := time.Parse("2006-01-02", "2022-07-29")
	events, _ = repo.GetUserEventsByMonth(1, from3)
	require.Equal(t, 0, len(events))
}

func createNewEvent() *Event {
	event, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang",
		"2022-07-15 20:00:00",
		"2022-07-15 21:30:00",
		"Сервис календаря",
		1,
		time.Duration(5400),
	)

	return event
}

func createEvents() []*Event {
	event1, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang 1",
		"2022-07-15 20:00:00",
		"2022-07-15 21:30:00",
		"Сервис календаря: работа с SQL",
		1,
		time.Duration(5400),
	)

	event2, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang 2",
		"2022-07-16 20:00:00",
		"2022-07-16 21:30:00",
		"Сервис календаря: конфигурирование",
		1,
		time.Duration(5400),
	)

	event3, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang 3",
		"2022-07-27 20:00:00",
		"2022-07-27 21:30:00",
		"Сервис календаря: работа с RabbitMQ",
		1,
		time.Duration(5400),
	)

	event4, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang 4",
		"2022-08-15 20:00:00",
		"2022-08-15 21:30:00",
		"Сервис календаря: интеграционное тестирование",
		1,
		time.Duration(5400),
	)

	event5, _ := NewEvent(
		uuid.New(),
		"Занятие по Golang 5",
		"2022-07-15 20:00:00",
		"2022-07-15 21:30:00",
		"Сервис календаря: разработка API",
		2,
		time.Duration(5400),
	)

	return []*Event{event1, event2, event3, event4, event5}
}

func getRepositoryWithEvents() *MemoryRepository {
	repo := NewMemoryRepository()
	events := createEvents()
	for _, event := range events {
		_ = repo.AddEvent(event)
	}

	return repo
}
