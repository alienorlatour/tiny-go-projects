package tick

import (
	"time"
)

// Tick corresponds to a new event for a Habit.
type Tick struct {
	Timestamp time.Time
}
