package habit

// ID is the identifier of the Habit.
type ID string

// Name is a short string that represents the name of a Habit.
type Name string

// WeeklyFrequency is the number of times a Habit should happen every week.
type WeeklyFrequency uint

// TickCount defines a number of weekly ticks.
type TickCount uint

// Habit to track.
type Habit struct {
	ID              ID
	Name            Name
	WeeklyFrequency WeeklyFrequency
	Ticks           TickCount
}
