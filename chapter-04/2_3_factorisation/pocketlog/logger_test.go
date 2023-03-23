package pocketlog_test

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-03/2_3_factorisation/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
