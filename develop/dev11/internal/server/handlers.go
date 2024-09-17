package server

import (
	"dev11/internal/api"
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) MapHandlers() http.Handler {

	mux := http.NewServeMux()

	api.MapCalendarHandlers(mux, s.log, s.domain)

	loggedMux := LoggingMiddleware(mux, s.log)

	return loggedMux
}
