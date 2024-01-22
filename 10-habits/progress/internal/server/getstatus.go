package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learngo-pockets/habits/internal/isoweek"
	"log"
	"time"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// GetHabitStatus is the endpoint that retrieves the status of a habit per week.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	log.Printf("GetStatus request received: %s", request)
	h, ticksCount, err := habit.GetStatus(ctx, s.db, s.db, habit.ID(request.HabitId), isoweek.At(time.Now()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot get status %q: %s", h.ID, err.Error())
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
