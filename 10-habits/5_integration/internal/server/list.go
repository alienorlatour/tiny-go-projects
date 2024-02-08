package server

import (
	"context"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

func (s *Server) ListHabits(ctx context.Context, _ *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	log.Println("ListHabits request received")
	habits, err := habit.ListHabits(ctx, s.db)
	if err != nil {
		return nil, err // todo wrap
	}

	return convertHabitsToAPI(habits), nil
}

func convertHabitsToAPI(habits []habit.Habit) *api.ListHabitsResponse {
	hts := make([]*api.Habit, len(habits))

	for i := range habits {
		frequency := int32(habits[i].WeeklyFrequency)
		hts[i] = &api.Habit{
			Id:              string(habits[i].ID),
			Name:            string(habits[i].Name),
			WeeklyFrequency: frequency,
		}
	}

	// TODO: Do we want to redo a sort here? Deterministicity matters.
	return &api.ListHabitsResponse{
		Habits: hts,
	}
}
