package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// TickHabit inserts a new tick for a given habit.
func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("Tick request received: %s", request)
	err := habit.Tick(ctx, s.db, s.db, habit.ID(request.HabitId), time.Now())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot tick habit %q: %s", request.HabitId, err.Error())
	}

	return &api.TickHabitResponse{}, nil
}
