package handlers

import (
	"context"
	"net/http"
	"time"

	"learngo-pockets/templates/internal/habit"
)

// Client is the dependency towards the Habits service.
//
//go:generate minimock -i habitsClient -s "_mock.go" -o "mocks"
type habitsClient interface {
	ListHabits(ctx context.Context, t time.Time) ([]habit.Habit, error)
	TickHabit(ctx context.Context, id habit.ID) error
	CreateHabit(ctx context.Context, h habit.Habit) error
}

// Server serves all the HTML routes on this service.
type Server struct {
	client habitsClient
	lgr    Logger
}

// New builds a new server.
func New(cli habitsClient, lgr Logger) *Server {
	return &Server{
		client: cli,
		lgr:    lgr,
	}
}

const (
	indexPath = "/"
)

// Router returns an HTTP handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodGet+" "+indexPath, s.index)
	r.HandleFunc(http.MethodGet+" "+"/tick/{habitID}", s.tick)
	r.HandleFunc(http.MethodPost+" "+"/create", s.create)

	// serve static files
	r.HandleFunc(http.MethodGet+" "+"/assets/styles.css", s.styles)

	return r
}

func (s *Server) logAndHideError(w http.ResponseWriter, endpoint string, err error, httpStatus int) {
	s.lgr.Logf("Error in %s: %s", endpoint, err.Error())
	http.Error(w, "Error while rendering - please retry.", httpStatus)
}

type Logger interface {
	Logf(format string, args ...any)
}
