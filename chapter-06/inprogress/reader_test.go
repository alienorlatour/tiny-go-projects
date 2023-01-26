package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"reflect"
	"testing"
)

func TestLoadReaders(t *testing.T) {
	noError := func(err error) bool { return err == nil }

	tests := map[string]struct {
		readersFile string
		want        []Reader
		checkError  func(err error) bool
	}{
		"no common book": {
			readersFile: "testdata/no_common_book.json",
			want:        readersWithNoCommonBooks,
			checkError:  noError,
		},
		"file doesn't exist": {
			readersFile: "testdata/no_file_here.json",
			checkError: func(err error) bool {
				return errors.Is(err, fs.ErrNotExist)
			},
		},
		"invalid JSON": {
			readersFile: "testdata/invalid.json",
			checkError: func(err error) bool {
				var expectedErr *json.SyntaxError
				return errors.As(err, &expectedErr)
			},
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadReaders(testCase.readersFile)
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
	readersWithNoCommonBooks = []Reader{
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
	twoReadersWithACommonBook = []Reader{
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
	threeReadersWithACommonBook = []Reader{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Authors: "Margaret Atwood",
					Title:   "Oryx and Crake",
				},
				{
					Authors: "Niccolo Machiavelli",
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
					Authors: "Niccolo Machiavelli",
					Title:   "Il Principe",
				},
			},
		},
	}
	readersWithTwoBooksByTheSameAuthorInCommon = []Reader{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Authors: "Niccolo Machiavelli",
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
					Authors: "Niccolo Machiavelli",
					Title:   "Il Principe",
				},
			},
		},
	}
)

func TestFindMatchingBooks(t *testing.T) {
	tt := map[string]struct {
		input []Reader
		want  []Book
	}{
		"no common book": {
			input: readersWithNoCommonBooks,
			want:  []Book{},
		},
		"one common book": {
			input: twoReadersWithACommonBook,
			want:  []Book{{Authors: "Charlotte Brontë", Title: "Jane Eyre"}},
		},
		"three readers have the same books on their shelves": {
			input: threeReadersWithACommonBook,
			want: []Book{
				{Authors: "Charlotte Brontë", Title: "Jane Eyre"},
				{Authors: "Niccolo Machiavelli", Title: "Il Principe"},
			},
		},
		"output is sorted by authors and then title": {
			input: readersWithTwoBooksByTheSameAuthorInCommon,
			want: []Book{
				{Authors: "Charlotte Brontë", Title: "Jane Eyre"},
				{Authors: "Charlotte Brontë", Title: "Villette"},
				{Authors: "Niccolo Machiavelli", Title: "Il Principe"},
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
