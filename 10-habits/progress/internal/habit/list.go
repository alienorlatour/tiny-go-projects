package habit

import (
	"context"
	"fmt"
)

type habitLister interface {
	FindAll(ctx context.Context) ([]Habit, error)
}

func ListHabits(ctx context.Context, db habitLister) ([]Habit, error) {
	habits, err := db.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot list habits: %w", err)
	}

	return habits, nil
}
