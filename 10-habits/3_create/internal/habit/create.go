package habit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CreateHabit adds a habit into the DB.
func CreateHabit(_ context.Context, h Habit) (Habit, error) {
	h, err := completeHabit(h)
	if err != nil {
		return h, err
	}

	// log h
	fmt.Println(h)

	return h, nil
}

// completeHabit fills the habit with values that we want in our database.
// Returns InvalidInputError. #D
func completeHabit(h Habit) (Habit, error) {
	// name cannot be empty
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return h, InvalidInputError{field: "name", reason: "cannot be empty"}
	}

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
