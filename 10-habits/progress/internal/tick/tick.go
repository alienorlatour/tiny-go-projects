package tick

import (
	"time"
)

// Tick corresponds to a new event for a Habit.
type Tick struct {
	Timestamp time.Time
}

// ISOWeek holds the number of the week and the year.
type ISOWeek struct {
	Year int
	Week int
}

// GetISOWeek returns the ISOWeek corresponding to the current time.
func GetISOWeek() ISOWeek {
	t := time.Now()
	y, w := t.ISOWeek()

	return ISOWeek{Year: y, Week: w}
}
