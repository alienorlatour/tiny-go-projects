package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"reflect"
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	villette      = Book{Author: "Charlotte Brontë", Title: "Villette"}
	ilPrincipe    = Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}
)

func TestLoadBookworms(t *testing.T) {
	noError := func(err error) bool { return err == nil }

	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		checkError    func(err error) bool
	}{
		"no common book": {
			bookwormsFile: "testdata/no_common_book.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			checkError: noError,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			checkError: func(err error) bool {
				return errors.Is(err, fs.ErrNotExist)
			},
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			checkError: func(err error) bool {
				var expectedErr *json.SyntaxError
				return errors.As(err, &expectedErr)
			},
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)
			if !testCase.checkError(err) {
				t.Fatalf("unexpected error: %s", err.Error())
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func TestFindMatchingBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		// TODO FIX ME
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Book{},
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre}},
			},
			want: []Book{janeEyre},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, ilPrincipe, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{janeEyre, ilPrincipe}},
			},
			want: []Book{janeEyre, ilPrincipe},
		},
		"output is sorted by authors and then title": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{ilPrincipe, janeEyre, villette}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{villette, ilPrincipe}},
			},
			want: []Book{janeEyre, villette, ilPrincipe},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findMatchingBooks(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}
