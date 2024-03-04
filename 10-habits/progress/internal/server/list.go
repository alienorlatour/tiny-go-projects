package server

import (
	"context"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

func (s *Server) ListHabits(ctx context.Context, _ *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	s.lgr.Logf("ListHabits request received")
	habits, err := habit.ListHabits(ctx, s.db)
	if err != nil {
		return nil, err // todo wrap
	}

	return convertHabitsToAPI(habits), nil
}

func convertHabitsToAPI(habits []habit.Habit) *api.ListHabitsResponse {
	hts := make([]*api.Habit, len(habits))

	for i := range habits {
		hts[i] = &api.Habit{
			Id:              string(habits[i].ID),
			Name:            string(habits[i].Name),
			WeeklyFrequency: int32(habits[i].WeeklyFrequency),
		}
	}

	return &api.ListHabitsResponse{
		Habits: hts,
	}
}
