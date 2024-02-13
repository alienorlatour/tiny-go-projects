package client

import (
	"context"
	"time"

	"learngo-pockets/habits/api"
	"learngo-pockets/templates/internal/habit"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListHabits lists the habits available.
func (hc *HabitsClient) ListHabits(ctx context.Context, t time.Time) ([]habit.Habit, error) {
	resp, err := hc.cli.ListHabits(ctx, &api.ListHabitsRequest{})
	if err != nil {
		return nil, err
	}

	list := make([]habit.Habit, len(resp.Habits))
	for i, h := range resp.Habits {
		// get status at time t
		status, err := hc.cli.GetHabitStatus(ctx, &api.GetHabitStatusRequest{
			HabitId:   h.Id,
			Timestamp: timestamppb.New(t),
		})
		if err != nil {
			return nil, err
		}

		// build habit struct
		list[i] = habit.Habit{
			ID:              habit.ID(h.Id),
			Name:            habit.Name(h.Name),
			WeeklyFrequency: habit.TickCount(h.WeeklyFrequency),
			Ticks:           habit.TickCount(status.TicksCount),
		}
	}
	return list, nil
}
