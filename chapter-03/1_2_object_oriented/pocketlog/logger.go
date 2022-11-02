package pocketlog

// Logger is used to log information.
type Logger struct {
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	// implement me
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// implement me
}
