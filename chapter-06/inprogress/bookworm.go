package main

import (
	"encoding/json"
	"io"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Authors string `json:"authors"`
	Title   string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	bookworms := make([]Bookworm, 0)
	err = json.Unmarshal(contents, &bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// findMatchingBooks returns books that are on more than one bookworm's shelf.
func findMatchingBooks(bookworms []Bookworm) []Book {
	booksOnShelves := make(map[Book]uint)

	// Register all books on shelves.
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			booksOnShelves[book]++
		}
	}

	matchingBooks := make([]Book, 0)

	// Find books that were added to shelves more than once.
	for book, count := range booksOnShelves {
		if count > 1 {
			matchingBooks = append(matchingBooks, book)
		}
	}

	sort.Slice(matchingBooks, func(i, j int) bool {
		if matchingBooks[i].Authors != matchingBooks[j].Authors {
			return matchingBooks[i].Authors < matchingBooks[j].Authors
		}
		return matchingBooks[i].Title < matchingBooks[j].Title
	})

	return matchingBooks
}

// bookSuggestions references all the related books to one book.
type bookSuggestions map[Book]bookList

// bookList is a collection of books.
type bookList map[Book]struct{}

// suggestOtherBooks from the matching-book shelf.
func suggestOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookSuggestions)

	// register
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelf := make([]Book, i, len(bookworm.Books)-1)
			copy(otherBooksOnShelf, bookworm.Books[:i])
			otherBooksOnShelf = append(otherBooksOnShelf, bookworm.Books[i+1:]...)

			registerBookSuggestions(sb, book, otherBooksOnShelf)
		}
	}

	// suggest books
	suggestions := make([]Bookworm, len(bookworms))
	for i, b := range bookworms {
		suggestions[i] = Bookworm{
			Name:  b.Name,
			Books: findSimilarBooks(sb, b.Books),
		}
	}

	return suggestions
}

func registerBookSuggestions(sb bookSuggestions, reference Book, books []Book) {
	for _, book := range books {
		if sb[reference] == nil {
			sb[reference] = make(bookList)
		}

		sb[reference][book] = struct{}{}
	}
}

func findSimilarBooks(sb bookSuggestions, myBooks []Book) []Book {
	bl := make(bookList)

	readBooks := make(map[Book]bool)
	for _, book := range myBooks {
		readBooks[book] = true
	}

	for _, book := range myBooks {
		for suggestion := range sb[book] {
			if readBooks[suggestion] {
				continue
			}

			bl[suggestion] = struct{}{}
		}
	}

	suggestions := make([]Book, 0, len(bl))
	for book := range bl {
		suggestions = append(suggestions, book)
	}

	return suggestions
}
