package books

import (
	"testing"

	"learngo-pockets/genericworms/collectors"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	villette      = Book{Author: "Charlotte Brontë", Title: "Villette"}
	ilPrincipe    = Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}
)

func TestLoad(t *testing.T) {
	tests := map[string]struct {
		collsFile string
		want      []collectors.Collector[Book]
		wantErr   bool
	}{
		"file exists": {
			collsFile: "testdata/colls.json",
			want: []collectors.Collector[Book]{
				{Name: "Fadi", Items: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Items: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
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

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, colls, target []collectors.Collector[Book]) bool {
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
		if !equalBooks(t, colls[i].Items, target[i].Items) {
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
