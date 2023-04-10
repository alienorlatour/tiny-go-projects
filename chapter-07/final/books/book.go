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
		// We shouldn't end up calling Before a non-book entity.
		// The returned value here doesn't mean that books matter less than other things.
		return false
	}

	if b.Author != other.Author {
		return b.Author < other.Author
	}

	return b.Title < other.Title
}

func (b Book) String() string {
	return fmt.Sprint(b.Title + ", by " + b.Author)
}
