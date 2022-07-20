package repository

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	DateStart    time.Time `json:"start_date" db:"start_date"`
	DateFinish   time.Time `json:"end_date" db:"end_date"`
	Description  string    `json:"description" db:"description"`
	UserID       int       `json:"user_id" db:"user_id"`
	NotifyBefore int       `json:"notify_before" db:"notify_before"`
}

const timeTemplate = "2006-01-02 15:04:05"

func NewEvent(
	ID uuid.UUID,
	title string,
	dateStart string,
	dateFinish string,
	description string,
	userID int,
	notifyBefore time.Duration,
) (*Event, error) {
	start, err := time.Parse(timeTemplate, dateStart)
	if err != nil {
		return nil, err
	}

	finish, err := time.Parse(timeTemplate, dateFinish)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:           ID,
		Title:        title,
		DateStart:    start,
		DateFinish:   finish,
		Description:  description,
		UserID:       userID,
		NotifyBefore: int(notifyBefore.Seconds()),
	}, nil
}
