package main

import "sort"

// bookRecommendations are the recommended books for a given book.
// Recommendations are the neighbour-books on the shelves.
type bookRecommendations map[Book]bookCollection

// bookCollection is a collection of books.
type bookCollection map[Book]struct{}

// newCollection initialises a new bookCollection.
func newCollection() bookCollection {
	return make(bookCollection)
}

// recommendOtherBooks returns the recommendations book from the common-book.
func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)

	// Register all books on everyone's shelf.
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			// Each book on the shelf will be a recommendation for those who have read any book on the shelf.
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	// Recommend a list of related books to each bookworm.
	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

// listOtherBooksOnShelves returns the list of my books except the one at the given index.
func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	// Initialise the first slice: its capacity is the input slice's capacity reduced by 1, and its starting length is the number of items up to the index to discard.
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	// Copy the slice of books up to the given index into the initialised index.
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	// Append with the rest of myBooks after the index of the discarded book into the created slice.
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove+1:]...)

	return otherBooksOnShelves
}

// registerBookRecommendations registers the books on the same shelf.
func registerBookRecommendations(recommendations bookRecommendations, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		// Check if this reference has already been added to the map.
		collection, ok := recommendations[reference]
		if !ok {
			// Create a new bookCollection.
			collection = newCollection()
			recommendations[reference] = collection
		}

		// Fill the associated books for the book reference.
		collection[book] = struct{}{}
	}
}

// recommendBooks returns the list of recommended books for a reference book.
func recommendBooks(recommendations bookRecommendations, myBooks []Book) []Book {
	bc := make(bookCollection)

	// Register all the books on shelves.
	// This step helps us to not recommend a book that has already been read.
	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	// Fill recommendations with all the neighbour-books.
	for _, myBook := range myBooks {
		// Find recommendations in other bookworm's shelves.
		for recommendation := range recommendations[myBook] {
			// Find recommendations in other bookworms' shelves.
			if myShelf[recommendation] {
				// Book already on the shelf.
				continue
			}

			// Add the book as a recommendation in the collection of books.
			bc[recommendation] = struct{}{}
		}
	}

	// Transform the map of books into an array.
	recommendationsForABook := bookCollectionToListOfBooks(bc)

	return recommendationsForABook
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
