package main

import (
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
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)

			switch {
			case err != nil && !testCase.wantErr:
				t.Fatalf("expected an error %s, got an empty one", err.Error())
			case err == nil && testCase.wantErr:
				t.Fatalf("expected no error, got one %s", err.Error())
			case !equalBookworms(t, got, testCase.want):
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()
	if len(bookworms) != len(target) {
		// Early exit!
		return false
	}

	for i := range bookworms {
		// Verify the name of the Bookworm.
		if bookworms[i].Name != target[i].Name {
			return false
		}
		// Verify the content of the collections of Books for each Bookworm.
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	// Everything is equal!
	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()
	if len(books) != len(target) {
		// Early exit!
		return false
	}
	// Verify the content of the collections of Books for each Bookworm.
	for i := range target {
		if target[i] != books[i] {
			return false
		}
	}
	// Everything is equal!
	return true
}
