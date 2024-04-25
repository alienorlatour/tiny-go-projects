//go:build !exercise

package habit

import (
	"time"
)

// FormattedWeek defines the start and end of a week, and formats them.
// Start to end will be from Sunday 00:00 to Saturday 23:59, rounded to the minute.
type FormattedWeek struct {
	start, end time.Time
	layout     string
}

// NewWeek builds an immutable week from one moment inside it.
func NewWeek(include time.Time, layout string) FormattedWeek {
	start := startOfDay(include).AddDate(0, 0, -int(include.Weekday()))
	return FormattedWeek{
		start:  start,
		end:    start.AddDate(0, 0, 7).Add(-1 * time.Minute),
		layout: layout,
	}
}

// startOfDay returns the same date as given, at midnight.
func startOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Start returns the formatted first day of the week.
func (w FormattedWeek) Start() string {
	return w.start.Format(w.layout)
}

// End returns the formatted last day of the week.
func (w FormattedWeek) End() string {
	return w.end.Format(w.layout)
}

func (w FormattedWeek) Next() int64 {
	return w.start.Add(time.Hour * 24 * 7).Unix()
}

func (w FormattedWeek) Previous() int64 {
	return w.start.Add(-time.Hour * 24 * 7).Unix()
}
