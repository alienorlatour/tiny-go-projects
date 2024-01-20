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
	Week int
	Year int
}

// GetISOWeek returns the ISOWeek of the current date.
func GetISOWeek() ISOWeek {
	t := time.Now()
	w, y := t.ISOWeek()
	return ISOWeek{
		Week: w,
		Year: y,
	}
}
