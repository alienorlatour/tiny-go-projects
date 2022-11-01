package pocketlog

import "io"

// Option defines a configuration option that will be passed to our logger via New()
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
