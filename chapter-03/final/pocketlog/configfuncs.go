package pocketlog

import "io"

// ConfigFunc defines a configuration option that will be passed to our logger via New()
type ConfigFunc func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) ConfigFunc {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
