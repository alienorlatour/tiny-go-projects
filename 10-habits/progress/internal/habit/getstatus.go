package habit

import (
	"context"
	"fmt"
	"time"
)

//go:generate minimock -i tickFinder -s "_mock.go" -o "mocks"
type tickFinder interface {
	FindWeeklyTicks(ctx context.Context, id ID, t time.Time) ([]time.Time, error)
}

// GetStatus returns the status a habit.
func GetStatus(ctx context.Context, habitDB habitFinder, tickDB tickFinder, id ID, t time.Time) (Habit, int, error) {
	h, err := habitDB.Find(ctx, id)
	if err != nil {
		return Habit{}, 0, fmt.Errorf("cannot find habit %s: %w", id, err)
	}

	ticks, err := tickDB.FindWeeklyTicks(ctx, id, t)
	if err != nil {
		return Habit{}, 0, fmt.Errorf("cannot find weekly ticks for habit %q: %w", id, err)
	}

	return h, len(ticks), nil
}
