package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestLoadReaders(t *testing.T) {
	tests := map[string]struct {
		readersFile string
		want        []Reader
		wantError   error
	}{
		"no common book": {
			readersFile: "testdata/no_common_book.json",
			want:        noCommonBookContents,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadReaders(testCase.readersFile)
			if !errors.Is(err, testCase.wantError) {
				t.Fatalf("unexpected error: %s", err.Error())
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

var (
	noCommonBookContents = []Reader{
		{
			Name: "Fadi",
			Books: []Book{
				{
					Authors: "Margaret Atwood",
					Title:   "The Handmaid's Tale",
				},
				{
					Authors: "Sylvia Plath",
					Title:   "The Bell Jar",
				},
			},
		},
		{
			Name: "Peggy",
			Books: []Book{
				{
					Authors: "Margaret Atwood",
					Title:   "Oryx and Crake",
				},
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
			},
		},
	}
)
