package server

import (
	"dev11/internal/domain"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	log    *logrus.Logger
	domain *domain.Domain
	config *Config
}

func New(config *Config, domain *domain.Domain, log *logrus.Logger) *Server {
	return &Server{
		log:    log,
		domain: domain,
		config: config,
	}
}

func (s *Server) Run() error {
	handler := s.MapHandlers()
	return http.ListenAndServe(s.config.BindAddr, handler)
}
