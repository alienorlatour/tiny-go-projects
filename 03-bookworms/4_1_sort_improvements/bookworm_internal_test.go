package main

import (
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
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
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

			if testCase.wantErr {
				if err == nil {
					t.Fatal("expected err, got nothing")
				}
				return
			}

			// we aren't expecting errors here, this should be the happy path
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !equalBookworms(t, got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  map[Book]uint
	}

	tt := map[string]testCase{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{handmaidsTale: 1, theBellJar: 1},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 3, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  []Book
	}

	tt := map[string]testCase{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
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
			got := findCommonBooks(tc.input)
			if !equalBooks(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
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

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	// Ranging over the want to retrieve all the keys.
	for book, targetCount := range want {
		// Verify the book in present in the map we check against.
		count, ok := got[book]
		// Book is not found or if found, counts are different.
		if !ok || targetCount != count {
			return false
		}
	}

	// Everything is equal!
	return true
}
