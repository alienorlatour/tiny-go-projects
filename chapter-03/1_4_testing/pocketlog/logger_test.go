package pocketlog_test

import "github.com/ablqk/tiny-go-projects/chapter-03/1_4_testing/pocketlog"

// ExampleLogger_Debug should fail, because the code is coming in the next chapter.
func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
