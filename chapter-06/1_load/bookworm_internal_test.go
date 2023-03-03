package main

import (
	"reflect"
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		err           error
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			err: nil,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := loadBookworms(testCase.bookwormsFile)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}
