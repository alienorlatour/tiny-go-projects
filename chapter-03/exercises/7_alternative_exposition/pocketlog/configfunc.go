package pocketlog

import "io"

// ConfigFunc defines what a configuration function, an optional parameter to New, to change the behaviour of the Logger
type ConfigFunc func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) ConfigFunc {
	return func(l *Logger) {
		l.output = output
	}
}
