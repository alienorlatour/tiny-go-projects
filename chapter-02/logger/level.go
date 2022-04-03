package logger

// Level describes the criticalness of a log message
type Level int

const (
	// LevelDebug messages are used to debug the application
	LevelDebug Level = iota
	// LevelInfo messages are used to log meaningful information about the processes going on
	LevelInfo
	// LevelError messages are used to highlight unexpected behaviours caught by the application
	LevelError
)
