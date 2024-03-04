package log

import (
	"io"
	"log"
	"sync"
)

// A Logger that can log messages
type Logger struct {
	mutex  sync.Mutex
	logger *log.Logger
	level  Level
}

// New returns a logger.
func New(output io.Writer, level Level) *Logger {
	if level < Debug || level > Error {
		level = Info
	}

	return &Logger{
		logger: log.New(output, "", log.Ldate|log.Ltime),
		level:  level,
	}
}

// Level is used to specify both the threshold of our logger, and the severity of a message
type Level byte

const (
	// Unset is the zero-value. Don't use it.
	Unset Level = iota
	// Debug messages, used for debugging.
	Debug
	// Info messages, contain valuable and not-flooding information.
	Info
	// Warn is used for non-blocking errors.
	Warn
	// Error is used when a blocking error was faced.
	Error
)

// Logf sends a message to the log if the severity is high enough.
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if lvl >= l.level {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		l.logger.Printf(format, args...)
	}
}
