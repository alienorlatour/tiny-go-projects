package habit

// ID is the identifier of the Habit.
type ID string

// Habit to track.
type Habit struct {
	ID        ID
	Name      string
	Frequency uint
}
