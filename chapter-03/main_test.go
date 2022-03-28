package main

import (
	"testing"
)

func Test_input(t *testing.T) {
	expected := "hello"
	reader := testReader{
		line: []byte(expected),
	}

	got := input(reader)

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
