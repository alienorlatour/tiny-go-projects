package server

import (
	"context"
	"log"
	"time"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("TickHabit request received: %s", request)

	s.tickDB.Add(ctx, habit.ID(request.Id), tick.Tick{Timestamp: time.Now()})

	return &api.TickHabitResponse{}, nil
}
