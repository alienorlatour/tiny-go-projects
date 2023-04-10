package pocketlog_test

import "learngo-pockets/logger/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// We don't have any output yet.
	// Output:
}
