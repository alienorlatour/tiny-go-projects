package log

import (
	"io"
	"log"
)

var logger log.Logger

// Set sets the logger output with the given writer.
func Set(writer io.Writer) {
	l := log.New(writer, "habit tracker", log.Ldate+log.Ltime)
	logger = *l
}

// Debugf formats and prints a message.
func Debugf(format string, args ...any) {
	logger.Printf(format, args...)
}

// Infof formats and prints a message.
func Infof(format string, args ...any) {
	logger.Printf(format, args...)
}

// Errorf formats and prints a message.
func Errorf(format string, args ...any) {
	logger.Printf(format, args...)
}

// Fatalf formats and prints a message.
func Fatalf(format string, args ...any) {
	logger.Printf(format, args...)
	return
}
