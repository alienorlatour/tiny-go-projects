package habit

import (
	"context"
	"fmt"
	"time"

	"learngo-pockets/habits/internal/tick"
)

//go:generate minimock -i habitFinder -s "_mock.go" -o "mocks"
type habitFinder interface {
	Find(ctx context.Context, id ID) (Habit, error)
}

//go:generate minimock -i tickAdder -s "_mock.go" -o "mocks"
type tickAdder interface {
	Add(ctx context.Context, id ID, t tick.Tick, w tick.ISOWeek) error
}

func TickHabit(ctx context.Context, habitDB habitFinder, tickDB tickAdder, id ID) error {
	// Check if the habit exists.
	h, err := habitDB.Find(ctx, id)
	if err != nil {
		return err
	}

	if len(h.ID) == 0 {
		return fmt.Errorf("habit ID \"%s\" not found", id)
	}

	// Add a new tick for the habit.
	err = tickDB.Add(ctx, id, tick.Tick{Timestamp: time.Now()}, tick.GetISOWeek())
	if err != nil {
		return err
	}

	return nil
}
