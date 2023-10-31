package server

import (
	"context"
	"fmt"
	"log/slog"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	slog.Info(fmt.Sprintf("CreateHabit request received: %s", request))

	var freq uint
	if request.Habit.Frequency != nil && uint(*request.Habit.Frequency) > 0 {
		freq = uint(*request.Habit.Frequency)
	} else {
		freq = 1
	}

	h := habit.Habit{
		Name:      habit.Name(request.Habit.Name),
		Frequency: habit.WeeklyFrequency(freq),
	}

	err := habit.CreateHabit(ctx, s.db, h)
	if err != nil {
		return nil, fmt.Errorf("cannot save habit %v: %w", h, err)
	}

	return &api.CreateHabitResponse{
		Habit: request.Habit,
	}, nil
}
