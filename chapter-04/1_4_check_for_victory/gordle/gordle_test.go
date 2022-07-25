package gordle

import (
	"testing"
)

func TestGordle_Play(t *testing.T) {
	expected := "hello"
	reader := testReader{
		line: []byte(expected),
	}

	g := &Gordle{
		reader:   reader,
		solution: []rune(expected),
	}

	got := g.Play()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

type testReader struct {
	line []byte
}

func (tr testReader) ReadLine() (line []byte, isPrefix bool, err error) {
	return tr.line, false, nil
}
