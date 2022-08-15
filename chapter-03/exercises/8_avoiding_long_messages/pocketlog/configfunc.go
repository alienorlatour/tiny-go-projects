package pocketlog

import "io"

// ConfigFunc defines what a configuration function, an optional parameter to New, to change the behaviour of the Logger
type ConfigFunc func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) ConfigFunc {
	return func(lgr *Logger) {
		lgr.output = output
	}
}

// WithMaxMessageLength sets the maximum length, in characters, of a message.
// Use 0 for no maximum length.
func WithMaxMessageLength(maxMessageLength uint) ConfigFunc {
	return func(l *Logger) {
		l.maxMessageLength = maxMessageLength
	}
}
