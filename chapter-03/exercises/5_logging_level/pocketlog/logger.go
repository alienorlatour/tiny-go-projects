package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(LevelError, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)
}
