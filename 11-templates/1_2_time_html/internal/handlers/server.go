package handlers

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

type Server struct {
	router chi.Router
}

func New() *Server {
	return &Server{}
}

func (s Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.index)

	return r
}
