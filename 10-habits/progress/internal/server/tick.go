package server

import (
	"context"
	"log"

	"learngo-pockets/habits/api"
)

func (s *Server) TickHabit(_ context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("TickHabit request received: %s", request)

	return &api.TickHabitResponse{}, nil
}
