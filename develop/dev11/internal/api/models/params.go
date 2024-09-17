package models

import (
	"strconv"
	"time"
)

type Params struct {
	Date        time.Time `json:"date"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func NewParams(userId, date, title, description string) (*Params, error) {

	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return &Params{
		UserId:      id,
		Date:        dateTime,
		Title:       title,
		Description: description,
	}, nil
}
