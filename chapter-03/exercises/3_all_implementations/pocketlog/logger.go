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

// Debugf formats and prints a message if the log level is debug or higher.
func (l Logger) Debugf(format string, args ...any) {
	if l.threshold <= LevelDebug {
		fmt.Printf(format+"\n", args...)
	}
}

// Infof formats and prints a message if the log level is info or higher.
func (l Logger) Infof(format string, args ...any) {
	if l.threshold <= LevelInfo {
		fmt.Printf(format+"\n", args...)
	}
}

// Errorf formats and prints a message if the log level is error or higher.
func (l Logger) Errorf(format string, args ...any) {
	if l.threshold <= LevelError {
		fmt.Printf(format+"\n", args...)
	}
}
