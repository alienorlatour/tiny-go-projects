package server

import (
	"context"
	"fmt"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

// GetHabitStatus is the endpoint that retrieves the status of a habit per week.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	log.Printf("GetHabitStatus request received: %s", request)

	id := habit.ID(request.Id)

	h, err := s.db.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(h.ID) == 0 {
		return nil, fmt.Errorf("habit ID \"%s\" not found", id)
	}

	ticks, err := s.tickDB.FindWeeklyTicks(ctx, id, tick.GetISOWeek())
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
