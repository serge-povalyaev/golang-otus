package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	"github.com/jmoiron/sqlx"
	"time"
)

type PgsqlRepository struct {
	conn *sqlx.DB
}

func NewPgsqlRepository(connection *sqlx.DB) *PgsqlRepository {
	return &PgsqlRepository{connection}
}

func (eventRepository *PgsqlRepository) GetEvent(userID int, eventID uuid.UUID) (*Event, error) {
	var foundEvent *Event
	sql := `SELECT * FROM events WHERE id = $1 AND user_id = $2 LIMIT 1`
	err := eventRepository.conn.Get(&foundEvent, sql, userID, eventID)

	if err != nil {
		return nil, err
	}

	if foundEvent == nil {
		return nil, ErrEventNotFound
	}

	return foundEvent, nil
}

func (eventRepository *PgsqlRepository) AddEvent(event *Event) error {
	_, err := eventRepository.findEvent(event.ID)
	if errors.Is(err, ErrEventNotFound) {
		sql := `INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6, $7)`
		eventRepository.conn.MustExec(
			sql,
			event.ID.String(),
			event.DateStart,
			event.DateFinish,
			event.Title,
			event.Description,
			event.UserID,
			event.NotifyBefore,
		)

		return nil
	}

	if err != nil {
		return err
	}

	return ErrEventExists
}

func (eventRepository *PgsqlRepository) UpdateEvent(userID int, event *Event) error {
	_, err := eventRepository.GetEvent(userID, event.ID)
	if err != nil {
		return err
	}

	sql := `UPDATE events SET 
                  date_start = $1, 
                  date_finish = $2, 
                  title = $3, 
                  description = $4, 
                  notify_before = $5
                  WHERE id = $6`
	eventRepository.conn.MustExec(
		sql,
		event.DateStart,
		event.DateFinish,
		event.Title,
		event.Description,
		event.NotifyBefore,
		event.ID.String(),
	)

	return nil
}

func (eventRepository *PgsqlRepository) RemoveEvent(userID int, event *Event) error {
	_, err := eventRepository.GetEvent(userID, event.ID)
	if err != nil {
		return err
	}

	eventRepository.conn.MustExec(`DELETE FROM events WHERE id = $1`, event.ID.String())

	return nil
}

func (eventRepository *PgsqlRepository) GetUserEvents(userID int, dateStart, dateFinish time.Time) ([]*Event, error) {
	events := make([]*Event, 0)
	sql := `SELECT * 
					 FROM events 
					 WHERE user_id = $1 
					   AND (date_start BETWEEN $2 AND $3 OR date_finish BETWEEN $2 AND $3)`
	err := eventRepository.conn.Select(
		&events,
		sql,
		userID,
		dateStart.Format("2006-01-02"),
		dateFinish.Format("2006-01-02"),
	)

	return events, err
}

func (eventRepository *PgsqlRepository) GetUserEventsByDay(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfDay())
}

func (eventRepository *PgsqlRepository) GetUserEventsByWeek(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfWeek())
}

func (eventRepository *PgsqlRepository) GetUserEventsByMonth(userID int, dateStart time.Time) ([]*Event, error) {
	return eventRepository.GetUserEvents(userID, dateStart, now.With(dateStart).EndOfMonth())
}

func (eventRepository *PgsqlRepository) findEvent(eventID uuid.UUID) (*Event, error) {
	var foundEvent Event
	sql := `SELECT * FROM events WHERE id = $1 LIMIT 1`
	err := eventRepository.conn.Get(&foundEvent, sql, eventID)
	if err != nil {
		return nil, err
	}

	if &foundEvent != nil {
		return &foundEvent, nil
	}

	return nil, ErrEventNotFound
}
