package main

import "sort"

// bookSuggestions are the suggested books for a given book.
// Suggestions are the neighbour-books on the shelves.
type bookSuggestions map[Book]bookCollection

// bookCollection is a collection of books.
type bookCollection map[Book]struct{}

// suggestOtherBooks returns the suggestions book from the matching-book.
func suggestOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookSuggestions)

	// Register all books on everyone's shelf.
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			// Each book on the shelf will be a suggestion for those who have read any book on the shelf.
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookSuggestions(sb, book, otherBooksOnShelves)
		}
	}

	// Suggest a list of related books to each bookworm.
	suggestions := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		suggestions[i] = Bookworm{
			Name:  bookworm.Name,
			Books: suggestBooks(sb, bookworm.Books),
		}
	}

	return suggestions
}

// listOtherBooksOnShelves returns the list of my books except the one at the given index.
func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	// Initialise the first array of books with a length until the given index.
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	// Copy the slice of books up to the given index into the initialised index.
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	// Append with the rest of myBooks after the index of the discarded book into the created slice.
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1:]...)

	return otherBooksOnShelves
}

// registerBookSuggestions registers the books on the same shelf.
func registerBookSuggestions(suggestions bookSuggestions, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		// Check if this reference has already been added to the map.
		if suggestions[reference] == nil {
			// Create a new bookCollection.
			suggestions[reference] = make(bookCollection)
		}

		// Fill the associated books for the book reference.
		suggestions[reference][book] = struct{}{}
	}
}

// suggestBooks returns the list of suggested books for a reference book.
func suggestBooks(suggestions bookSuggestions, myBooks []Book) []Book {
	bc := make(bookCollection)

	// Register all the books on shelves.
	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	// Fill suggestions with all the neighbour-books.
	for _, myBook := range myBooks {
		// Find suggestions in other bookworm's shelves.
		for suggestion := range suggestions[myBook] {
			if myShelf[suggestion] {
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

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}
		return bookList[i].Title < bookList[j].Title
	})

	return bookList
}
