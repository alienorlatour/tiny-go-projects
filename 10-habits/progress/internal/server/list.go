package server

import (
	"context"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

func (s *Server) ListHabits(ctx context.Context, request *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	habits, err := s.db.FindAll(ctx)
	if err != nil {
		return nil, err // todo wrap
	}

	return response(habits), nil
}

func response(habits []habit.Habit) *api.ListHabitsResponse {
	hts := make([]*api.Habit, len(habits))

	for i := range habits {
		frequency := int32(habits[i].Frequency)
		hts[i] = &api.Habit{
			Name:      habits[i].Name,
			Frequency: &frequency,
		}
	}

	return &api.ListHabitsResponse{
		Habits: hts,
	}
}
