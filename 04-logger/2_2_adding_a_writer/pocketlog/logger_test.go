package pocketlog_test

import (
	"os"

	"learngo-pockets/logger/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
