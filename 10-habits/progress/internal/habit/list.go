package habit

import (
	"context"
	"fmt"
)

//go:generate minimock -i habitLister -s "_mock.go" -o "mocks"
type habitLister interface {
	FindAll(ctx context.Context) ([]Habit, error)
}

// ListHabits returns the habits that we have in the database.
// TODO: Do we want this to be deterministic?
func ListHabits(ctx context.Context, db habitLister) ([]Habit, error) {
	habits, err := db.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot list habits: %w", err)
	}

	return habits, nil
}
