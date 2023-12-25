package habit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

//go:generate minimock -i habitCreator -s "_mock.go" -o "mocks"
type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

// CreateHabit adds a habit into the DB.
func CreateHabit(ctx context.Context, db habitCreator, h Habit) (Habit, error) {
	h, err := completeHabit(h)
	if err != nil {
		return Habit{}, err
	}

	err = db.Add(ctx, h)
	if err != nil {
		return Habit{}, fmt.Errorf("cannot save habit: %w", err)
	}

	return h, nil
}

// completeHabit fills the habit with values that we want in our database.
// Returns InvalidInputError.
func completeHabit(h Habit) (Habit, error) {
	// name cannot be empty
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return h, InvalidInputError{field: "name", reason: "cannot be empty"}
	}

	// default to 1
	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1
	}

	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h, nil
}
