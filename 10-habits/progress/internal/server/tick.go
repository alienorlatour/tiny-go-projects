package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

// TickHabit inserts a new tick for a given habit.
func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("TickHabit request received: %s", request)

	id := habit.ID(request.Id)

	h, err := s.db.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(h.ID) == 0 {
		return nil, fmt.Errorf("habit ID \"%s\" not found", id)
	}

	err = s.tickDB.Add(ctx, habit.ID(request.Id), tick.Tick{Timestamp: time.Now()}, tick.GetISOWeek())
	if err != nil {
		return nil, err
	}

	return &api.TickHabitResponse{}, nil
}
