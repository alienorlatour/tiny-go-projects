package books_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/genericworms/books"
)

func TestDecodeBook(t *testing.T) {
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
