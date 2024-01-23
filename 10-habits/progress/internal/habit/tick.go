package habit

import (
	"context"
	"fmt"
	"time"
)

//go:generate minimock -i habitFinder -s "_mock.go" -o "mocks"
type habitFinder interface {
	Find(ctx context.Context, id ID) (Habit, error)
}

//go:generate minimock -i tickAdder -s "_mock.go" -o "mocks"
type tickAdder interface {
	AddTick(ctx context.Context, id ID, t time.Time) error
}

// Tick inserts a new tick for a habit.
func Tick(ctx context.Context, habitDB habitFinder, tickDB tickAdder, id ID, t time.Time) error {
	// Check if the habit exists.
	h, err := habitDB.Find(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot find habit %q: %w", id, err)
	}

	if len(h.ID) == 0 {
		return fmt.Errorf("habit ID %q not found", id)
	}

	// AddTick adds a new tick for the habit.
	err = tickDB.AddTick(ctx, id, t)
	if err != nil {
		return fmt.Errorf("cannot insert tick for habit %q: %w", id, err)
	}

	return nil
}
