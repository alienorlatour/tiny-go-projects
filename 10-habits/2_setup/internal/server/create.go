package server

import (
	"context"

	"learngo-pockets/habits/api"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("CreateHabit request received: %s", request)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{},
	}, nil
}
