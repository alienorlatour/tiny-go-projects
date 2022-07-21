package pocketlog_test

import "tiny-go-projects/chapter03/2_1_first_implementation/pocketlog"

func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debug("Hello, %s", "world")
	// Output: Hello, world
}
