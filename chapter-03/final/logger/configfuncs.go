package logger

import "io"

type ConfigFunc func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) ConfigFunc {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
