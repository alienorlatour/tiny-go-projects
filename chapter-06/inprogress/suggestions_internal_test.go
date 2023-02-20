package main

import (
	"reflect"
	"testing"
)

var (
	theHelp   = Book{Author: "Kathryn Stockett", Title: "The Help"}
	fairyTale = Book{Author: "Stephen King", Title: "Fairy Tale"}
)

func TestSuggestOtherBooks(t *testing.T) {
	tt := map[string]struct {
		bookworms []Bookworm
		want      []Bookworm
	}{
		"No common books": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
			},
		},
		"Two bookworms with a common book and several books on their respective shelves": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, ilPrincipe}},
				{Name: "Did", Books: []Book{janeEyre, theBellJar}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{theBellJar}},
				{Name: "Did", Books: []Book{oryxAndCrake, ilPrincipe}},
			},
		},
		"Two bookworms with a common book and one has only one book on shelf": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{janeEyre}},
				{Name: "Did", Books: []Book{janeEyre, theBellJar}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{theBellJar}},
				{Name: "Did", Books: []Book{}},
			},
		},
		"Three bookworms with two common books and several books on their respective shelves": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{theHelp, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre, theHelp, fairyTale}},
				{Name: "Ali", Books: []Book{janeEyre, ilPrincipe, theHelp}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{fairyTale, ilPrincipe}},
				{Name: "Did", Books: []Book{oryxAndCrake, ilPrincipe}},
				{Name: "Ali", Books: []Book{oryxAndCrake, fairyTale}},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := suggestOtherBooks(tc.bookworms)
			if !equals(got, tc.want) {
				t.Errorf("suggestOtherBooks() = %v, want %v", got, tc.want)
			}
		})
	}
}

// equals compares two list of Bookworms.
func equals(bookwormA, bookwormB []Bookworm) bool {
	for i := range bookwormA {
		if !reflect.DeepEqual(bookwormB[i], bookwormB[i]) {
			return false
		}
	}
	return true
}
