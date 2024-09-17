package models

import (
	"time"
)

type Event struct {
	Date        time.Time
	Title       string
	Description string
}

func NewEvent(date time.Time, description, title string) *Event {
	return &Event{
		Date:        date,
		Description: description,
		Title:       title,
	}
}
