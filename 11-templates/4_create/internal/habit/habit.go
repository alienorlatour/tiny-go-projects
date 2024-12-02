package habit

// ID is the identifier of the Habit.
type ID string

// Name is a short string that represents the name of a Habit.
type Name string

// TickCount defines a number of weekly ticks.
type TickCount uint

// Habit to track.
type Habit struct {
	ID              ID
	Name            Name
	WeeklyFrequency TickCount
	Ticks           TickCount
}

// IsDone returns whether a habit has been fully completed.
func (h *Habit) IsDone() bool {
	return h.Ticks >= h.WeeklyFrequency
}
