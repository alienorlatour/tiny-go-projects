package pocketlog_test

import (
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-03/final/pocketlog"
)

func ExampleLogger_Debug_debug() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debug("Hello, %s", "world")
	// Output: [DEBUG] Hello, world
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_LevelInfo(t *testing.T) {
	tt := map[string]struct {
		level    pocketlog.Level
		expected string
	}{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debug(debugMessage)
			testedLogger.Info(infoMessage)
			testedLogger.Error(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
