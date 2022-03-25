package logger_test

import "github.com/ablqk/tiny-go-projects/chapter-02/logger"

func ExampleLogger_Debug_debug() {
	debugLogger := logger.New(logger.LevelDebug)
	debugLogger.Debug("Hello, %s", "world")
	// Output: Hello, world
}

func ExampleLogger_Debug_info() {
	infoLogger := logger.New(logger.LevelInfo)
	infoLogger.Debug("Hello, %s", "world")
	// Output:
}
