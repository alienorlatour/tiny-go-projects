package isoweek

import (
	"time"
)

// ISO8601 holds the number of the week and the year.
type ISO8601 struct {
	Year int
	Week int
}

// At returns the ISO8601 week and year at the given timestamp.
// The definition for week 01 is the week with the first Thursday of January in it.
func At(t time.Time) ISO8601 {
	y, w := t.ISOWeek()
	return ISO8601{Year: y, Week: w}
}
