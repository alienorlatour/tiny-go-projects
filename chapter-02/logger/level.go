package logger

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

// String implements the Stringer interface
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
