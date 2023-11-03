package server

import (
	"context"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

func (s *Server) ListHabits(ctx context.Context, request *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	habits, err := habit.ListHabits(ctx, s.db)
	if err != nil {
		return nil, err // todo wrap
	}

	return response(habits), nil
}

func response(habits []habit.Habit) *api.ListHabitsResponse {
	hts := make([]*api.Habit, len(habits))

	for i := range habits {
		frequency := int32(habits[i].WeeklyFrequency)
		hts[i] = &api.Habit{
			Name:            string(habits[i].Name),
			WeeklyFrequency: &frequency,
		}
	}

	return &api.ListHabitsResponse{
		Habits: hts,
	}
}
