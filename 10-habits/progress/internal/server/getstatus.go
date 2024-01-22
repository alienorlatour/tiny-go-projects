package server

import (
	"context"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// GetHabitStatus is the endpoint that retrieves the status of a habit per week.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	log.Printf("GetHabitStatus request received: %s", request)
	h, ticksCount, err := habit.GetHabitStatus(ctx, s.db, s.tickDB, habit.ID(request.HabitId))
	if err != nil {
		return nil, err
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
