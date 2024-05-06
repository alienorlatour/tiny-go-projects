package log

import (
	"io"
	"log"
)

// Logger is a lightweight logger.
type Logger struct {
	lgr *log.Logger
}

// New returns a logger that prints to the specified output.
func New(w io.Writer) *Logger {
	return &Logger{lgr: log.New(w, "", log.Ldate|log.Ltime)}
}

// Logf formats the message and sends it to the output.
func (l *Logger) Logf(format string, args ...any) {
	l.lgr.Printf(format, args...)
}
