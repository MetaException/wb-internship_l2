package storage

import (
	"dev11/internal/domain/models"
	"errors"
	"time"
)

type Storage struct {
	events map[int][]models.Event
}

func NewStorage() *Storage {
	return &Storage{
		events: make(map[int][]models.Event),
	}
}

func (s *Storage) Add(user_id int, event models.Event) error {

	if _, ok := s.events[user_id]; !ok {
		s.events[user_id] = []models.Event{}
	}

	s.events[user_id] = append(s.events[user_id], event)
	return nil
}

func (s *Storage) Delete(user_id int, date time.Time) error {

	if events, ok := s.events[user_id]; ok {
		if idx := FindEvent(events, date); idx != -1 {
			s.events[user_id] = append(events[:idx], events[idx+1:]...)
			if len(s.events[user_id]) == 0 {
				delete(s.events, user_id)
			}
			return nil
		}
		return errors.New("event not found for the given date")
	}
	return errors.New("user_id not found")
}

func (s *Storage) Set(user_id int, event models.Event) error {

	if events, ok := s.events[user_id]; ok {
		if idx := FindEvent(events, event.Date); idx != -1 {
			s.events[user_id][idx] = event
			return nil
		}
		return errors.New("event not found for the given date")
	}
	return errors.New("not found")
}

func (s *Storage) GetEventsForDay(user_id int, date time.Time) ([]models.Event, error) {

	if events, ok := s.events[user_id]; ok {
		var result []models.Event
		for _, v := range events {
			if v.Date == date {
				result = append(result, v)
			}
		}
		return result, nil
	}
	return nil, errors.New("user_id not found")
}

func (s *Storage) GetEventsForWeek(user_id int, date time.Time) ([]models.Event, error) {

	if events, ok := s.events[user_id]; ok {
		var result []models.Event
		for _, v := range events {
			y1, w1 := v.Date.ISOWeek()
			y2, w2 := date.ISOWeek()
			if y1 == y2 && w1 == w2 {
				result = append(result, v)
			}
		}
		return result, nil
	}
	return nil, errors.New("user_id not found")
}

func (s *Storage) GetEventsForMonth(user_id int, date time.Time) ([]models.Event, error) {
	if events, ok := s.events[user_id]; ok {
		var result []models.Event
		for _, v := range events {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
				result = append(result, v)
			}
		}
		return result, nil
	}
	return nil, errors.New("user_id not found")
}

func FindEvent(events []models.Event, date time.Time) int {

	for i, v := range events {
		if v.Date == date {
			return i
		}
	}
	return -1
}
