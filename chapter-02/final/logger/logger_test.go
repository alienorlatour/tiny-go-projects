package logger_test

import (
	"testing"

	"tiny-go-projects/chapter-02/final/logger"
)

func ExampleLogger_Debug_debug() {
	debugLogger := logger.New(logger.LevelDebug)
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
		level          logger.Level
		expectedOutput string
	}{
		"debug": {
			level:          logger.LevelDebug,
			expectedOutput: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"info": {
			level:          logger.LevelInfo,
			expectedOutput: "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"error": {
			level:          logger.LevelError,
			expectedOutput: "[ERROR] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := logger.New(tc.level, logger.WithOutput(tw))

			testedLogger.Debug(debugMessage)
			testedLogger.Info(infoMessage)
			testedLogger.Error(errorMessage)

			if tw.contents != tc.expectedOutput {
				t.Errorf("invalid contents, expected %q, got %q", tc.expectedOutput, tw.contents)
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
