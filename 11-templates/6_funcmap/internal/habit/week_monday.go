//go:build exercise

package habit

import (
	"time"
)

// FormattedWeek defines the start and end of a week, and formats them.
// If weekStart is Monday, start to end will be
// from Monday 00:00 to Sunday 23:59, rounded to the minute.
type FormattedWeek struct {
	start, end time.Time
	layout     string
}

func NewWeek(include time.Time, layout string) FormattedWeek {
	const weekStart = int(time.Monday)
	sinceLastWeek := int(include.Weekday()) // number of days since the end of last Sunday
	if sinceLastWeek == 0 {
		// last Sunday was a full week ago
		sinceLastWeek += 7
	}

	start := startOfDay(include).AddDate(0, 0, -sinceLastWeek+weekStart)
	return FormattedWeek{
		start:  start,
		end:    start.AddDate(0, 0, 7).Add(-1 * time.Minute),
		layout: layout,
	}
}

func startOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (w FormattedWeek) Start() string {
	return w.start.Format(w.layout)
}

func (w FormattedWeek) End() string {
	return w.end.Format(w.layout)
}
