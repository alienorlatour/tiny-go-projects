package handlers

import (
	"context"
	"net/http"
	"time"

	"learngo-pockets/templates/internal/habit"

	"github.com/go-chi/chi/v5"
)

// HabitsClient is the dependency towards the Habits service.
//
//go:generate minimock -s "_mock.go" -o "mocks"
type HabitsClient interface {
	ListHabits(ctx context.Context, t time.Time) ([]habit.Habit, error)
}

// Server serves all the HTML routes on this service.
type Server struct {
	client HabitsClient
	lgr    Logger
}

// New builds a new server.
func New(cli HabitsClient, lgr Logger) *Server {
	return &Server{
		client: cli,
		lgr:    lgr,
	}
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
