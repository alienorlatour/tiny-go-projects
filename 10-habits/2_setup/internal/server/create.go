package server

import (
	"context"
	"log"

	"learngo-pockets/habits/api"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	log.Printf("CreateHabit request received: %s", request)

	return &api.CreateHabitResponse{
		Habit: request.Habit,
	}, nil
}
