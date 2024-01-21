package habit

import (
	"context"
	"fmt"

	"learngo-pockets/habits/internal/tick"
)

//go:generate minimock -i tickFinder -s "_mock.go" -o "mocks"
type tickFinder interface {
	FindWeeklyTicks(ctx context.Context, id ID, w tick.ISOWeek) ([]tick.Tick, error)
}

func GetHabitStatus(ctx context.Context, habitDB habitFinder, tickDB tickFinder, id ID) (Habit, int, error) {
	h, err := habitDB.Find(ctx, id)
	if err != nil {
		return Habit{}, 0, err
	}

	if len(h.ID) == 0 {
		return Habit{}, 0, fmt.Errorf("habit ID \"%s\" not found", id)
	}

	ticks, err := tickDB.FindWeeklyTicks(ctx, id, tick.GetISOWeek())
	if err != nil {
		return Habit{}, 0, err
	}

	return h, len(ticks), nil
}
