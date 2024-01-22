package habit

import (
	"context"
	"fmt"
	"time"

	"learngo-pockets/habits/internal/isoweek"
)

//go:generate minimock -i tickFinder -s "_mock.go" -o "mocks"
type tickFinder interface {
	FindWeeklyTicks(ctx context.Context, id ID, w isoweek.ISO8601) ([]time.Time, error)
}

// GetStatus returns the status a habit.
func GetStatus(ctx context.Context, habitDB habitFinder, tickDB tickFinder, id ID, isoWeek isoweek.ISO8601) (Habit, int, error) {
	h, err := habitDB.Find(ctx, id)
	if err != nil {
		return Habit{}, 0, fmt.Errorf("cannot find habit %q: %w", id, err)
	}

	if len(h.ID) == 0 {
		return Habit{}, 0, fmt.Errorf("habit ID %q not found", id)
	}

	ticks, err := tickDB.FindWeeklyTicks(ctx, id, isoWeek)
	if err != nil {
		return Habit{}, 0, fmt.Errorf("cannot find weekly ticks for habit %q: %w", id, err)
	}

	return h, len(ticks), nil
}
