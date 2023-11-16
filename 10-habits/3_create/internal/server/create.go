package server

import (
	"context"
	"errors"
	"fmt"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	log.Printf("CreateHabit request received: %s", request)

	var freq uint
	if request.Habit.WeeklyFrequency != nil && uint(*request.Habit.WeeklyFrequency) > 0 {
		freq = uint(*request.Habit.WeeklyFrequency)
	}

	h := habit.Habit{
		Name:            habit.Name(request.Habit.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}

	err := habit.CreateHabit(ctx, h)
	if err != nil {
		invalidErr := habit.InvalidInputError{}
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		// other error
		return nil, fmt.Errorf("cannot save habit %v: %w", h, err)
	}

	return &api.CreateHabitResponse{
		Habit: request.Habit,
	}, nil
}
