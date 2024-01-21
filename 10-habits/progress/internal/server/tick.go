package server

import (
	"context"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// TickHabit inserts a new tick for a given habit.
func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("TickHabit request received: %s", request)
	err := habit.TickHabit(ctx, s.db, s.tickDB, habit.ID(request.Id))
	if err != nil {
		return nil, err
	}

	return &api.TickHabitResponse{}, nil
}
