package server

import (
	"context"
	"errors"
	"fmt"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	var freq uint
	if request.WeeklyFrequency != nil {
		freq = uint(*request.WeeklyFrequency)
	}

	h := habit.Habit{
		Name:            habit.Name(request.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}

	got, err := habit.CreateHabit(ctx, h)
	if err != nil {
		invalidErr := habit.InvalidInputError{}
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		// other error
		return nil, fmt.Errorf("cannot save habit %v: %w", h, err)
	}

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id:              string(got.ID),
			Name:            string(got.Name),
			WeeklyFrequency: int32(got.WeeklyFrequency),
		},
	}, nil
}
