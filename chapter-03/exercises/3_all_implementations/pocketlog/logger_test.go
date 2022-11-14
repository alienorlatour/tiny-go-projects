package pocketlog_test

import "github.com/ablqk/tiny-go-projects/chapter-03/exercises/3_all_implementations/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
