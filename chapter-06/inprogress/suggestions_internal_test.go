package main

import (
	"reflect"
	"testing"
)

func TestSuggestOtherBooks(t *testing.T) {
	tt := map[string]struct {
		bookworms []Bookworm
		want      []Bookworm
	}{
		"No common books": {
			bookworms: bookwormsWithNoCommonBooks,
			want:      suggestionsForBookwormsWithNoCommonBooks,
		},
		"Two bookworms with a common book and several books on their respective shelves": {
			bookworms: twoBookwormsWithACommonBookAndSeveralBooksOnShelves,
			want:      suggestionsForTwoBookwormsWithACommonBookAndSeveralBooksOnShelves,
		},
		"Two bookworms with a common book and one has only one book on shelf": {
			bookworms: twoBookwormsWithACommonBookAndOneBooksOnShelves,
			want:      suggestionsForTwoBookwormsWithACommonBookAndOneBooksOnShelves,
		},
		"Three bookworms with two common books and several books on their respective shelves": {
			bookworms: threeBookwormsWithTwoCommonBooksAndSeveralBooksOnShelves,
			want:      suggestionForThreeBookwormsWithTwoCommonBooksAndSeveralBooksOnShelves,
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

var (
	suggestionsForBookwormsWithNoCommonBooks = []Bookworm{
		{
			Name: "Fadi", Books: []Book{},
		},
		{
			Name: "Peggy", Books: []Book{},
		},
	}

	twoBookwormsWithACommonBookAndSeveralBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Margaret Atwood",
					Title:  "Oryx and Crake",
				},
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
				{
					Author: "Niccolò Machiavelli",
					Title:  "Il Principe",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
				{
					Author: "Sylvia Plath",
					Title:  "The Bell Jar",
				},
			},
		},
	}

	suggestionsForTwoBookwormsWithACommonBookAndSeveralBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Sylvia Plath",
					Title:  "The Bell Jar",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Author: "Margaret Atwood",
					Title:  "Oryx and Crake",
				},
				{
					Author: "Niccolò Machiavelli",
					Title:  "Il Principe",
				},
			},
		},
	}

	twoBookwormsWithACommonBookAndOneBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
				{
					Author: "Sylvia Plath",
					Title:  "The Bell Jar",
				},
			},
		},
	}

	suggestionsForTwoBookwormsWithACommonBookAndOneBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Sylvia Plath",
					Title:  "The Bell Jar",
				},
			},
		},
		{
			Name:  "Did",
			Books: []Book{},
		},
	}

	threeBookwormsWithTwoCommonBooksAndSeveralBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Kathryn Stockett",
					Title:  "The Help",
				},
				{
					Author: "Margaret Atwood",
					Title:  "Oryx and Crake",
				},
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
				{
					Author: "Kathryn Stockett",
					Title:  "The Help",
				},
				{
					Author: "Stephen King",
					Title:  "Fairy Tale",
				},
			},
		},
		{
			Name: "Ali",
			Books: []Book{
				{
					Author: "Charlotte Brontë",
					Title:  "Jane Eyre",
				},
				{
					Author: "Niccolò Machiavelli",
					Title:  "Il Principe",
				},
				{
					Author: "Kathryn Stockett",
					Title:  "The Help",
				},
			},
		},
	}

	suggestionForThreeBookwormsWithTwoCommonBooksAndSeveralBooksOnShelves = []Bookworm{
		{
			Name: "Peggy",
			Books: []Book{
				{
					Author: "Stephen King",
					Title:  "Fairy Tale",
				},
				{
					Author: "Niccolò Machiavelli",
					Title:  "Il Principe",
				},
			},
		},
		{
			Name: "Did",
			Books: []Book{
				{
					Author: "Margaret Atwood",
					Title:  "Oryx and Crake",
				},
				{
					Author: "Niccolò Machiavelli",
					Title:  "Il Principe",
				},
			},
		},
		{
			Name: "Ali",
			Books: []Book{
				{
					Author: "Margaret Atwood",
					Title:  "Oryx and Crake",
				},
				{
					Author: "Stephen King",
					Title:  "Fairy Tale",
				},
			},
		},
	}
)

// equals compares two list of Bookworms.
func equals(bookwormA, bookwormB []Bookworm) bool {
	for i := range bookwormA {
		if !reflect.DeepEqual(bookwormB[i], bookwormB[i]) {
			return false
		}
	}
	return true
}
