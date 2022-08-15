package pocketlog_test

import "tiny-go-projects/chapter03/2_2_adding_a_writer/pocketlog"

func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debug("Hello, %s", "world")
	// Output: Hello, world
}
