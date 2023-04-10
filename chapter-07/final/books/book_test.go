package books_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/genericworms/books"
	"learngo-pockets/genericworms/collectors"
)

func TestUnmarshalBook(t *testing.T) {
	bookJson := []byte(`
{
  "author": "Sylvia Plath",
  "title": "The Bell Jar"
}`)

	var book books.Book
	err := json.Unmarshal(bookJson, &book)

	assert.NoError(t, err, "unexpected error while unmarshalling book")
	assert.Equal(t, books.Book{Author: "Sylvia Plath", Title: "The Bell Jar"}, book)
}

func TestBookBefore(t *testing.T) {
	testCases := map[string]struct {
		left  books.Book
		right collectors.Sortable
		want  bool
	}{
		"Different authors": {
			left:  books.Book{Author: "Orhan Pamuk", Title: "Kırmızı Saçlı Kadın"},
			right: books.Book{Author: "Sylvia Plath", Title: "The Bell Jar"},
			want:  true,
		},
		"Different authors, reversed": {
			left:  books.Book{Author: "Sylvia Plath", Title: "The Bell Jar"},
			right: books.Book{Author: "Orhan Pamuk", Title: "Kırmızı Saçlı Kadın"},
			want:  false,
		},
		"Same author, different titles": {
			left:  books.Book{Author: "Sylvia Plath", Title: "The Bell Jar"},
			right: books.Book{Author: "Sylvia Plath", Title: "Lady Lazarus"},
			want:  false,
		},
		"can't compare a book to a non-book": {
			left:  books.Book{Author: "Orhan Pamuk", Title: "Kırmızı Saçlı Kadın"},
			right: nonBook{},
			want:  false,
		},
	}

	for name, testCase := range testCases {
		// go < 1.20
		// name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			got := testCase.left.Before(testCase.right)
			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestBookString(t *testing.T) {
	book := books.Book{
		Author: "Orhan Pamuk",
		Title:  "Kırmızı Saçlı Kadın",
	}

	want := `Kırmızı Saçlı Kadın, by Orhan Pamuk`

	assert.Equal(t, want, book.String())
}

type nonBook struct{}

func (n nonBook) Before(_ collectors.Sortable) bool {
	panic("This is a test utility, it shouldn't be called")
}
