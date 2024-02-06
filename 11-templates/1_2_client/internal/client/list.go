package client

import (
	"context"

	"learngo-pockets/habits/api"
	habit "learngo-pockets/templates/internal/habits"
)

// ListHabits lists the habits available.
func (hc HabitsClient) ListHabits(ctx context.Context) ([]habit.Habit, error) {
	resp, err := hc.cli.ListHabits(ctx, &api.ListHabitsRequest{})
	if err != nil {
		return nil, err
	}

	list := make([]habit.Habit, len(resp.Habits))
	for i, h := range resp.Habits {
		list[i] = habit.Habit{
			ID:              habit.ID(h.Id),
			Name:            habit.Name(h.Name),
			WeeklyFrequency: habit.WeeklyFrequency(h.WeeklyFrequency),
		}
	}
	return list, nil
}
