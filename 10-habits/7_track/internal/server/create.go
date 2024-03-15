package server

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("Create request received: %s", request)

	err := validateCreateHabitRequest(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: "+err.Error())
	}

	var freq uint
	if request.WeeklyFrequency != nil && uint(*request.WeeklyFrequency) > 0 {
		freq = uint(*request.WeeklyFrequency)
	}

	h := habit.Habit{
		Name:            habit.Name(request.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}

	createdHabit, err := habit.Create(ctx, s.db, h)
	if err != nil {
		var invalidErr habit.InvalidInputError
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		// other error
		return nil, status.Errorf(codes.Internal, "cannot save habit %v: %s", h, err.Error())
	}

	s.lgr.Logf("Habit %s successfully registered", createdHabit.ID)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id:              string(createdHabit.ID),
			Name:            string(createdHabit.Name),
			WeeklyFrequency: int32(createdHabit.WeeklyFrequency),
		},
	}, nil
}

func validateCreateHabitRequest(request *api.CreateHabitRequest) error {
	switch {
	case request == nil:
		return fmt.Errorf("empty request")
	case request.Name == "":
		return fmt.Errorf("missing name of habit")
	}
	return nil
}
