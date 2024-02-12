package client

import (
	"context"

	"learngo-pockets/habits/api"
	"learngo-pockets/templates/internal/habit"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TickHabit adds a tick now.
func (hc *HabitsClient) TickHabit(ctx context.Context, id habit.ID) error {
	_, err := hc.cli.TickHabit(ctx, &api.TickHabitRequest{
		HabitId:   string(id),
		Timestamp: timestamppb.Now(),
	})
	return err
}
