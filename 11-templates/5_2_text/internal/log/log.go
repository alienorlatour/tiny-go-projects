package log

import (
	"io"
	"log"
	"sync"
)

// Logger is a lightweight logger.
type Logger struct {
	lgr   *log.Logger
	mutex sync.Mutex
}

// New returns a logger that prints to the specified output.
func New(w io.Writer) *Logger {
	return &Logger{lgr: log.New(w, "", log.Ldate|log.Ltime)}
}

// Logf formats the message and sends it to the output. This is concurrent-safe.
func (l *Logger) Logf(format string, args ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.lgr.Printf(format, args...)
}
