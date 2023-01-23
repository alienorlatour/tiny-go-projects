package main

// A Reader contains the list of books on a reader's shelf.
type Reader struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a reader's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}
