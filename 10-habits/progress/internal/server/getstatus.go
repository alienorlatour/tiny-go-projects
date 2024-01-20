package server

import (
	"context"
	"learngo-pockets/habits/internal/habit"
	"log"

	"learngo-pockets/habits/api"
)

// GetHabitStatus is the endpoint that retrieve the status of a habit.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	log.Printf("GetHabitStatus request received: %s", request)

	id := habit.ID(request.Id)
	h, err := s.db.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	ticks, err := s.tickDB.FindAll(ctx, id)
	if err != nil {
		return nil, err
	}

	return &api.GetHabitStatusResponse{
		Habit: &api.Habit{
			Id:              string(h.ID),
			Name:            string(h.Name),
			WeeklyFrequency: int32(h.WeeklyFrequency),
		},
		TicksCount: int32(len(ticks)),
	}, nil
}
