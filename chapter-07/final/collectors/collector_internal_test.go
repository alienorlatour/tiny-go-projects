package collectors

import (
	"testing"
)

type item string

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input Collectors[item]
		want  map[item]uint
	}{
		"nominal use case": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "handmaidsTale", "janeEyre"}},
			},
			want: map[item]uint{"handmaidsTale": 2, "theBellJar": 1, "oryxAndCrake": 1, "janeEyre": 1},
		},
		"no colls": {
			input: Collectors[item]{},
			want:  map[item]uint{},
		},
		"coll without books": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{}},
			},
			want: map[item]uint{"handmaidsTale": 1, "theBellJar": 1},
		},
		"coll with twice the same book": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar", "handmaidsTale"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "handmaidsTale", "janeEyre"}},
			},
			want: map[item]uint{"handmaidsTale": 3, "theBellJar": 1, "oryxAndCrake": 1, "janeEyre": 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.countItems()

			if !equalCounts(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommon(t *testing.T) {
	tt := map[string]struct {
		input Collectors[item]
		want  []item
	}{
		"no common book": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "janeEyre"}},
			},
			want: nil,
		},
		"one common book": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"oryxAndCrake", "janeEyre"}},
				{Name: "Did", Items: []item{"janeEyre"}},
			},
			want: []item{"janeEyre"},
		},
		"three colls have the same books on their shelves": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"oryxAndCrake", "ilPrincipe", "janeEyre"}},
				{Name: "Did", Items: []item{"janeEyre"}},
				{Name: "Ali", Items: []item{"janeEyre", "ilPrincipe"}},
			},
			want: []item{"janeEyre", "ilPrincipe"},
		},
		"output is sorted by authors and then title": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"ilPrincipe", "janeEyre", "villette"}},
				{Name: "Did", Items: []item{"janeEyre"}},
				{Name: "Ali", Items: []item{"villette", "ilPrincipe"}},
			},
			want: []item{"janeEyre", "villette", "ilPrincipe"},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.FindCommon()
			if !equal(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// equalBooks is a helper to test the equality of two lists of Books.
func equal[T lesser](t *testing.T, books, target []T) bool {
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

// equalCounts is a helper to test the equality of two maps of books count.
func equalCounts[T lesser](t *testing.T, bookCount, target map[T]uint) bool {
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
