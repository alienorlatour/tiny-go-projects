package pocketlog

// Level represents one available logging level
type Level int

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors
	LevelError
)

// String implements the Stringer interface for the Level structure
func (lvl Level) String() string {
	switch lvl {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelError:
		return "[ERROR]"
	default:
		// a very unlikely possibility
		return ""
	}
}
