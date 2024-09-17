package api

import (
	"dev11/internal/domain"
	"net/http"

	"github.com/sirupsen/logrus"
)

func MapCalendarHandlers(mux *http.ServeMux, log *logrus.Logger, domain *domain.Domain) {

	h := NewCalendarHandler(log, domain)

	mux.HandleFunc("GET /events_for_day", h.GetEventsForDay)
	mux.HandleFunc("GET /events_for_week", h.GetEventsForWeek)
	mux.HandleFunc("GET /events_for_month", h.GetEventsForMonth)
	mux.HandleFunc("POST /create_event", h.CreateEvent)
	mux.HandleFunc("POST /update_event", h.UpdateEvent)
	mux.HandleFunc("POST /delete_event", h.DeleteEvent)
}
