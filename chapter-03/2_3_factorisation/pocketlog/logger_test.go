package pocketlog_test

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-03/2_2_adding_a_writer/pocketlog"
)

func ExampleLogger_Debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
