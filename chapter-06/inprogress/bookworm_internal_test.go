package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"reflect"
	"testing"
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
			want:          bookwormsWithNoCommonBooks,
			checkError:    noError,
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

var (
	bookwormsWithNoCommonBooks = []Bookworm{
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
	twoBookwormsWithACommonBook = []Bookworm{
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
		{
			Name: "Did",
			Books: []Book{
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
			},
		},
	}
	threeBookwormsWithACommonBook = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Authors: "Margaret Atwood",
					Title:   "Oryx and Crake",
				},
				{
					Authors: "Niccolò Machiavelli",
					Title:   "Il Principe",
				},
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
			},
		},
		{
			Name: "Ali",
			Books: []Book{
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
				{
					Authors: "Niccolò Machiavelli",
					Title:   "Il Principe",
				},
			},
		},
	}
	bookwormsWithTwoBooksByTheSameAuthorInCommon = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Authors: "Niccolò Machiavelli",
					Title:   "Il Principe",
				},
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
				{
					Authors: "Charlotte Brontë",
					Title:   "Villette",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Authors: "Charlotte Brontë",
					Title:   "Jane Eyre",
				},
			},
		},
		{
			Name: "Ali",
			Books: []Book{
				{
					Authors: "Charlotte Brontë",
					Title:   "Villette",
				},
				{
					Authors: "Niccolò Machiavelli",
					Title:   "Il Principe",
				},
			},
		},
	}
)

func TestFindMatchingBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: bookwormsWithNoCommonBooks,
			want:  []Book{},
		},
		"one common book": {
			input: twoBookwormsWithACommonBook,
			want:  []Book{{Authors: "Charlotte Brontë", Title: "Jane Eyre"}},
		},
		"three bookworms have the same books on their shelves": {
			input: threeBookwormsWithACommonBook,
			want: []Book{
				{Authors: "Charlotte Brontë", Title: "Jane Eyre"},
				{Authors: "Niccolò Machiavelli", Title: "Il Principe"},
			},
		},
		"output is sorted by authors and then title": {
			input: bookwormsWithTwoBooksByTheSameAuthorInCommon,
			want: []Book{
				{Authors: "Charlotte Brontë", Title: "Jane Eyre"},
				{Authors: "Charlotte Brontë", Title: "Villette"},
				{Authors: "Niccolò Machiavelli", Title: "Il Principe"},
			},
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
