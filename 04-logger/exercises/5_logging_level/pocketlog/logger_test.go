package pocketlog_test

import (
	"testing"

	"learngo-pockets/logger/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: [DEBUG] Hello, world
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugInfoError(t *testing.T) {
	tt := map[string]struct {
		level    pocketlog.Level
		expected string
	}{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] Why write I still all one, ever the same,\n[INFO] And keep invention in a noted weed,\n[ERROR] That every word doth almost tell my name,\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] And keep invention in a noted weed,\n[ERROR] That every word doth almost tell my name,\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] That every word doth almost tell my name,\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
