package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold        Level
	output           io.Writer
	maxMessageLength uint
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
// There is no default maximum length - messages aren't trimmed.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold:        threshold,
		output:           os.Stdout,
		maxMessageLength: 0, // we could get rid of this line and use the zero value but let's be explicit
	}

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

// Logf formats and prints a message if the log level is high enough
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	contents := fmt.Sprintf(format, args...)

	// check the trimming is activated, and that we should apply it to this message
	// checking the length in runes, as this won't print unexpected characters
	if l.maxMessageLength != 0 && uint(len([]rune(contents))) > l.maxMessageLength {
		contents = string([]rune(contents)[:l.maxMessageLength]) + "[TRIMMED]"
	}

	msg := message{
		Level:   lvl.String(),
		Message: contents,
	}

	// encode the message
	formattedMessage, err := json.Marshal(msg)
	if err != nil {
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
		return
	}

	_, _ = fmt.Fprintln(l.output, string(formattedMessage))
}

// message represents the JSON structure of the logged messages.
type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
