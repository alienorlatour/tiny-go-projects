package pocketlog_test

import "github.com/ablqk/tiny-go-projects/chapter-03/2_4_functional_options/pocketlog"

func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
