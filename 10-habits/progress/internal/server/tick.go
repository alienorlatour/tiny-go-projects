package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	r "learngo-pockets/habits/internal/repository"
)

// TickHabit inserts a new tick for a given habit.
func (s *Server) TickHabit(ctx context.Context, request *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	s.lgr.Logf("Tick request received: %s", request)

	err := validateTickHabitRequest(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: "+err.Error())
	}

	// if empty, the timestamp is set to the current time
	var t time.Time
	if request.Timestamp == nil {
		t = time.Now()
	}

	err = habit.Tick(ctx, s.db, s.db, habit.ID(request.HabitId), t)
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

func validateTickHabitRequest(request *api.TickHabitRequest) error {
	switch {
	case request == nil:
		return fmt.Errorf("empty request")
	case request.HabitId == "":
		return fmt.Errorf("missing habit id")
	}
	return nil
}
