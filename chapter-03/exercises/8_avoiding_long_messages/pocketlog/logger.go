package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	level            Level
	output           io.Writer
	maxMessageLength uint
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
// There is no default maximum length - messages aren't trimmed.
func New(level Level, configFuncs ...ConfigFunc) *Logger {
	lgr := &Logger{
		level:            level,
		output:           os.Stdout,
		maxMessageLength: 0,
	}
	for _, configFunc := range configFuncs {
		configFunc(lgr)
	}
	return lgr
}

// Debug formats and prints a message if the log level is debug or higher.
func (l Logger) Debug(format string, args ...any) {
	if l.level <= LevelDebug {
		l.log(LevelDebug, format, args...)
	}
}

// Info formats and prints a message if the log level is info or higher.
func (l Logger) Info(format string, args ...any) {
	if l.level <= LevelInfo {
		l.log(LevelInfo, format, args...)
	}
}

// Error formats and prints a message if the log level is error or higher.
func (l Logger) Error(format string, args ...any) {
	if l.level <= LevelError {
		l.log(LevelError, format, args...)
	}
}

// Log formats and prints a message if the log level is high enough
func (l Logger) Log(lvl Level, format string, args ...any) {
	if l.level <= lvl {
		l.log(lvl, format, args...)
	}
}

// log prints the message to the output.
// Add decorations here, if any.
func (l Logger) log(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	// check the trimming is activated, and that we should apply it to this message
	// checking the length in runes, as this won't print unexpected characters
	if l.maxMessageLength != 0 && uint(len([]rune(message))) > l.maxMessageLength {
		message = string([]rune(message)[:l.maxMessageLength]) + "[TRIMMED]"
	}
	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)
}
