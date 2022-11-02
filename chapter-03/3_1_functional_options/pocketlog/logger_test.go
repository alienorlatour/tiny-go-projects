package pocketlog_test

import "github.com/ablqk/tiny-go-projects/chapter-03/3_1_functional_options/pocketlog"

func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
