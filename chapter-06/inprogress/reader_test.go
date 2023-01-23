package main

import (
	"errors"
	"testing"

	"golang.org/x/exp/slices"
)

func TestLoadReaders(t *testing.T) {
	tests := map[string]struct {
		readersFile string
		want        []Reader
		wantError   error
	}{
		"no common book": {
			readersFile: "testdata/no_common_book.json",
			want:        []Reader{},
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			readers, err := loadReaders(testCase.readersFile)
			if !errors.Is(err, testCase.wantError) {
				t.Errorf("unexpected error: %s", err.Error())
			}
			if !slices.Equal[Reader](readers, testCase.want) {

			}
		})
	}
}

func (r Reader) Equal(other Reader) bool {
	if r.Name != other.Name {
		return false
	}

	return slices.Equal[Book](r.Books, other.Books)
}
