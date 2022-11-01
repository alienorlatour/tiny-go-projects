package pocketlog

import "fmt"

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold.
func New(level Level) *Logger {
	return &Logger{
		threshold: level,
	}
}

// Debug formats and prints a message if the log level is debug or higher.
func (l *Logger) Debug(format string, args ...any) {
	if l.threshold <= LevelDebug {
		fmt.Printf(format+"\n", args...)
	}
}

// Info formats and prints a message if the log level is info or higher.
func (l *Logger) Info(format string, args ...any) {
	// implement me
}

// Error formats and prints a message if the log level is error or higher.
func (l *Logger) Error(format string, args ...any) {
	// implement me
}
