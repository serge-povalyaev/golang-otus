package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"github.com/serge.povalyaev/calendar/internal/repository"
)

type App struct {
	logger          logger.CalendarLogger
	eventRepository repository.EventRepository
	connection      *sqlx.DB
}

func New(logger logger.CalendarLogger, eventRepository repository.EventRepository, connection *sqlx.DB) *App {
	return &App{
		logger:          logger,
		eventRepository: eventRepository,
		connection:      connection,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.repository.CreateEvent(repository.Event{ID: id, Title: title})
}

// TODO
