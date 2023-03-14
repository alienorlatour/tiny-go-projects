package collectors

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
	tests := map[string]struct {
		collsFile string
		want      []Collector
		wantErr   bool
	}{
		"file exists": {
			collsFile: "testdata/colls.json",
			want: []Collector{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			collsFile: "testdata/no_file_here.json",
			want:      nil,
			wantErr:   true,
		},
		"invalid JSON": {
			collsFile: "testdata/invalid.json",
			want:      nil,
			wantErr:   true,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Load(testCase.collsFile)

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

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Collector
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Collector{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no colls": {
			input: []Collector{},
			want:  map[Book]uint{},
		},
		"coll without books": {
			input: []Collector{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{handmaidsTale: 1, theBellJar: 1},
		},
		"coll with twice the same book": {
			input: []Collector{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 3, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := count(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Collector
		want  []Book
	}{
		"no common book": {
			input: []Collector{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Collector{
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre}},
			},
			want: []Book{janeEyre},
		},
		"three colls have the same books on their shelves": {
			input: []Collector{
				{Name: "Peggy", Books: []Book{oryxAndCrake, ilPrincipe, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{janeEyre, ilPrincipe}},
			},
			want: []Book{janeEyre, ilPrincipe},
		},
		"output is sorted by authors and then title": {
			input: []Collector{
				{Name: "Peggy", Books: []Book{ilPrincipe, janeEyre, villette}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{villette, ilPrincipe}},
			},
			want: []Book{janeEyre, villette, ilPrincipe},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := FindCommonBooks(tc.input)
			if !equalBooks(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, colls, target []Collector) bool {
	t.Helper()

	if len(colls) != len(target) {
		// Early exit!
		return false
	}

	for i := range colls {
		// Verify the name of the Collector.
		if colls[i].Name != target[i].Name {
			return false
		}
		// Verify the content of the collections of Books for each Collector.
		if !equalBooks(t, colls[i].Books, target[i].Books) {
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
		// Early exit
		return false
	}

	// Verify the content of the collections of Books for each Collector.
	for i := range target {
		if target[i] != books[i] {
			return false
		}
	}
	// Everything is equal!
	return true
}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(t *testing.T, bookCount, target map[Book]uint) bool {
	t.Helper()

	// Ranging over the target to retrieve all the keys.
	for book, targetCount := range target {
		// Verify the book in present in the map we check against.
		count, ok := bookCount[book]
		// Book is not found or if found, counts are different.
		if !ok || targetCount != count {
			return false
		}
	}

	// Everything is equal!
	return true
}
