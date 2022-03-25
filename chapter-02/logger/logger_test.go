package logger_test

import (
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-02/logger"
)

func ExampleLogger_Debug_debug() {
	debugLogger := logger.New(logger.LevelDebug)
	debugLogger.Debug("Hello, %s", "world")
	// Output: Hello, world
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_LevelInfo(t *testing.T) {
	tw := &testWriter{}
	infoLogger := logger.New(logger.LevelInfo).WithOutput(tw)
	expected := infoMessage + "\n" + errorMessage + "\n"

	infoLogger.Debug(debugMessage)
	infoLogger.Info(infoMessage)
	infoLogger.Error(errorMessage)

	if tw.contents != expected {
		t.Errorf("invalid contents, expected %q, got %q", expected, tw.contents)
	}
}

// testWriter is a struct that implements io.Writer
type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
