package handlers

import (
	"context"
	"net/http"

	habit "learngo-pockets/templates/internal/habits"

	chi "github.com/go-chi/chi/v5"
)

// Client is the dependency towards the Habits service.
//
//go:generate minimock -i habitsClient -s "_mock.go" -o "mocks"
type habitsClient interface {
	ListHabits(ctx context.Context) ([]habit.Habit, error)
}

// Server serves all the HTML routes on this service.
type Server struct {
	client habitsClient
	router chi.Router
}

// New builds a new server.
func New(cli habitsClient) *Server {
	return &Server{
		client: cli,
	}
}

// Router returns an http handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.index)

	return r
}
