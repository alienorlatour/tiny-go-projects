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

func ExampleLogger_Debug_info() {
	infoLogger := logger.New(logger.LevelInfo)
	infoLogger.Debug("Hello, %s", "world")
	// Output:
}

func TestLogger_WithOutput(t *testing.T) {
	tw := &testWriter{}
	infoLogger := logger.New(logger.LevelInfo).WithOutput(tw)
	message := "I will write. All may be well enough."
	expected := message + "\n"

	infoLogger.Info(message)

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
