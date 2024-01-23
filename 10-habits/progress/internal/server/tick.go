package server

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	r "learngo-pockets/habits/internal/repository"
)

// TickHabit inserts a new tick for a given habit.
func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	log.Printf("Tick request received: %s", request)
	err := habit.Tick(ctx, s.db, s.db, habit.ID(request.HabitId), time.Now())
	if err != nil {
		switch {
		case errors.Is(err, r.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "couldn't find habit %q in repository", request.HabitId)
		default:
			return nil, status.Errorf(codes.Internal, "cannot tick habit %q: %s", request.HabitId, err.Error())
		}
	}

	return &api.TickHabitResponse{}, nil
}
