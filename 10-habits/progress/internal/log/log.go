package log

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "habit tracker", log.Ldate+log.Ltime)
}

// Set sets the logger output with the given writer.
func Set(writer io.Writer) {
	logger.SetOutput(writer)
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
	os.Exit(1)
}
