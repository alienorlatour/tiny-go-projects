package handlers

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

// Server serves all the HTML routes on this service.
type Server struct {
	router chi.Router
}

// New builds a new server.
func New() *Server {
	return &Server{}
}

// Router returns an http handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/", s.index)
	return r
}