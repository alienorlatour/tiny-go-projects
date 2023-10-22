package server

import (
	"context"
	"fmt"
	"log/slog"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"

	"github.com/google/uuid"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	slog.Info(fmt.Sprintf("CreateHabit request received: %s", request))

	if request.Habit.Frequency == nil || uint(*request.Habit.Frequency) == 0 {
		return nil, fmt.Errorf("invalid frequency")
	}
	freq := *request.Habit.Frequency

	habit := habit.Habit{
		ID:        habit.ID(uuid.NewString()),
		Name:      request.Habit.Name,
		Frequency: uint(freq),
	}

	err := s.db.Add(habit)
	if err != nil {
		return nil, fmt.Errorf("cannot save habit %v: %w", habit, err)
	}

	return &api.CreateHabitResponse{
		Habit: request.Habit,
	}, nil
}
