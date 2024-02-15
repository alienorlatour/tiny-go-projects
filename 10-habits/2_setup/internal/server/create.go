package server

import (
	"context"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/log"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	log.Infof("CreateHabit request received: %s", request)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{},
	}, nil
}
