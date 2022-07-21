package pocketlog

// Logger is used to log information.
type Logger struct {
	level Level
}

// New returns you a logger, ready to log at the required threshold.
func New(level Level) *Logger {
	return &Logger{
		level: level,
	}
}

// Debug formats and prints a message if the log level is debug or higher.
func (l Logger) Debug(format string, args ...any) {
	// implement me
}

// Info formats and prints a message if the log level is info or higher.
func (l Logger) Info(format string, args ...any) {
	// implement me
}

// Error formats and prints a message if the log level is error or higher.
func (l Logger) Error(format string, args ...any) {
	// implement me
}
