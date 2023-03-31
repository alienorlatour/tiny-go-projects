package books

import (
	"fmt"

	"learngo-pockets/genericworms/collectors"
)

// Book describes a book on a collector's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func (b Book) Before(sortable collectors.Sortable) bool {
	other, ok := sortable.(Book)
	if !ok {
		return false
	}

	if b.Title != other.Title {
		return b.Title < other.Title
	}

	return b.Author < other.Author
}

func (b Book) String() string {
	return fmt.Sprint(b.Title + ", by " + b.Author)
}
