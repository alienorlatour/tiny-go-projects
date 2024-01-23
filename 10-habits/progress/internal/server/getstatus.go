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
	"learngo-pockets/habits/internal/isoweek"
	r "learngo-pockets/habits/internal/repository"
)

// GetHabitStatus is the endpoint that retrieves the status of a habit per week.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	log.Printf("GetStatus request received: %s", request)
	h, ticksCount, err := habit.GetStatus(ctx, s.db, s.db, habit.ID(request.HabitId), isoweek.At(time.Now()))
	if err != nil {
		switch {
		case errors.Is(err, r.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "couldn't find habit %q in repository", request.HabitId)
		default:
			return nil, status.Errorf(codes.Internal, "cannot get status %q: %s", h.ID, err.Error())
		}
	}

	return &api.GetHabitStatusResponse{
		Habit: &api.Habit{
			Id:              string(h.ID),
			Name:            string(h.Name),
			WeeklyFrequency: int32(h.WeeklyFrequency),
		},
		TicksCount: int32(ticksCount),
	}, nil
}
