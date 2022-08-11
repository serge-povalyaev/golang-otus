package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type EventRepository interface {
	AddEvent(event *Event) error
	GetEvent(userId int, eventID uuid.UUID) (*Event, error)
	UpdateEvent(userID int, event *Event) error
	RemoveEvent(userID int, event *Event) error
	GetUserEvents(userID int, dateStart, dateFinish time.Time) ([]*Event, error)
	GetUserEventsByDay(userID int, dateStart time.Time) ([]*Event, error)
	GetUserEventsByWeek(userID int, dateStart time.Time) ([]*Event, error)
	GetUserEventsByMonth(userID int, dateStart time.Time) ([]*Event, error)
}

func GetEventRepository(repositoryType string, conn *sqlx.DB) (eventRepository EventRepository) {
	switch repositoryType {
	case "memory":
		eventRepository = NewMemoryRepository()
	case "pgsql":
		eventRepository = NewPgsqlRepository(conn)
	}

	return eventRepository
}
