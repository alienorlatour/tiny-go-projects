package habit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//go:generate minimock -i habitCreator -s "_mock.go" -o "mocks"
type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

// CreateHabit adds a habit into the DB.
func CreateHabit(ctx context.Context, db habitCreator, h Habit) error {
	h = completeHabit(h)

	err := db.Add(ctx, h)
	if err != nil {
		return fmt.Errorf("cannot save habit: %w", err)
	}

	return nil
}

// completeHabit fills the habit with values that we want in our database.
func completeHabit(h Habit) Habit {
	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h
}
