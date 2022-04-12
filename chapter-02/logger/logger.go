package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information. Use New to get a bespoke logger!
type Logger struct {
	level  Level
	output io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(level Level) *Logger {
	return &Logger{
		level:  level,
		output: os.Stdout,
	}
}

// WithOutput sets the output of the logger, and returns it.
// You can call logger.WithOutput(os.StdOut).Info().
// The thread safety is left to the implementation of output.
func (l Logger) WithOutput(output io.Writer) Logger {
	l.output = output
	return l
}

// Debug formats and prints a message if the log level is debug or higher
func (l Logger) Debug(format string, args ...any) {
	if l.level <= LevelDebug {
		l.log(format, args)
	}
}

// Info formats and prints a message if the log level is info or higher
func (l Logger) Info(format string, args ...any) {
	if l.level <= LevelInfo {
		l.log(format, args)
	}
}

// Error formats and prints a message if the log level is error or higher
func (l Logger) Error(format string, args ...any) {
	if l.level <= LevelError {
		l.log(format, args)
	}
}

// log prints the message to the output.
// Add decorations here, if any.
func (l Logger) log(format string, args []any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}