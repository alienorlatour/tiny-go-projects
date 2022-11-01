package pocketlog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Logger is used to log information. Use New to get a bespoke logger!
type Logger struct {
	threshold   Level
	output      io.Writer
	outputMutex sync.Mutex
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will
// The default output is Stdout.
func New(threshold Level, configFuncs ...ConfigFunc) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}
	for _, configFunc := range configFuncs {
		configFunc(lgr)
	}
	return lgr
}

// Debug formats and prints a message if the log level is debug or higher
func (lgr *Logger) Debug(format string, args ...any) {
	lgr.Log(LevelDebug, format, args...)
}

// Info formats and prints a message if the log level is info or higher
func (lgr *Logger) Info(format string, args ...any) {
	lgr.Log(LevelInfo, format, args...)
}

// Error formats and prints a message if the log level is error or higher
func (lgr *Logger) Error(format string, args ...any) {
	lgr.Log(LevelError, format, args...)
}

// Log formats and prints a message if the log level is high enough
func (lgr *Logger) Log(lvl Level, format string, args ...any) {
	if lgr.threshold <= lvl {
		lgr.log(lvl, format, args...)
	}
}

// log prints the message to the output.
// Add decorations here, if any.
func (lgr *Logger) log(lvl Level, format string, args ...any) {
	lgr.outputMutex.Lock()
	defer lgr.outputMutex.Unlock()
	decoratedMessage := fmt.Sprintf("%s %s\n", lvl, format)
	_, _ = fmt.Fprintf(lgr.output, decoratedMessage, args...)
}
