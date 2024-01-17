package tick

import (
	"time"

	"learngo-pockets/habits/internal/habit"
)

// Tick corresponds to a new event for a Habit.
type Tick struct {
	HabitID   habit.ID
	Timestamp time.Time
}
