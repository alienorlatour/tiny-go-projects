package main

// bookSuggestions are the suggested books for a given book.
// Suggestions are the neighbour-books on the shelves.
type bookSuggestions map[Book]bookCollection

// bookCollection is a collection of books.
type bookCollection map[Book]struct{}

// suggestOtherBooks returns the suggestions book from the matching-book.
func suggestOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookSuggestions)

	// Register all books with their suggestions meaning the others books on the same shelves' owner.
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := buildOtherBooksOnShelves(i, bookworm.Books)
			registerBookSuggestions(sb, book, otherBooksOnShelves)
		}
	}

	// Suggest a list of related books to each bookworm.
	suggestions := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		suggestions[i] = Bookworm{
			Name:  bookworm.Name,
			Books: findSuggestionsBook(sb, bookworm.Books),
		}
	}

	return suggestions
}

// buildOtherBooksOnShelves returns the list of my books except the one at the given index.
func buildOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	// Initialise the first array of books with a length until the given index.
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	// Copy the array of books until the given index into the initialised array.
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	// Append with the rest of myBooks after the index to remove into the created array.
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1:]...)

	return otherBooksOnShelves
}

// registerBookSuggestions registers the books on the same shelf.
func registerBookSuggestions(suggestions bookSuggestions, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		// Book reference is not registered yet.
		if suggestions[reference] == nil {
			// Create a new bookCollection.
			suggestions[reference] = make(bookCollection)
		}

		// Fill the associated books for the book reference.
		suggestions[reference][book] = struct{}{}
	}
}

// findSuggestionsBook returns the list of suggested books for a reference book.
func findSuggestionsBook(suggestions bookSuggestions, myBooks []Book) []Book {
	bc := make(bookCollection)

	// Register all the books on shelves.
	booksOnShelves := make(map[Book]bool)
	for _, myBook := range myBooks {
		booksOnShelves[myBook] = true
	}

	//  Fill suggestions with all the neighbour-books.
	for _, myBook := range myBooks {
		for suggestion := range suggestions[myBook] {
			if booksOnShelves[suggestion] {
				// Book already on the shelf.
				continue
			}

			// Add the book as a suggestion in the collection of books.
			bc[suggestion] = struct{}{}
		}
	}

	// Transform the map of books into an array.
	suggestionsForABook := bookCollectionToListOfBooks(bc)

	return suggestionsForABook
}

// bookCollectionToListOfBooks transforms a bookCollection entities into a list of Book.
func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, 0, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}

	return bookList
}
