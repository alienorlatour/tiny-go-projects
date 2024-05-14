package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Server serves all the HTML routes on this service.
type Server struct {
	lgr Logger
}

// New builds a new server.
func New(lgr Logger) *Server {
	return &Server{lgr: lgr}
}

// Router returns an HTTP handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.index)

	return r
}

type Logger interface {
	Logf(format string, args ...any)
}
