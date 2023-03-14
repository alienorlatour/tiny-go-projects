package collectors

import (
	"golang.org/x/exp/constraints"
)

// A Collector contains the list of books on a collector's shelf.
type Collector[T comparable] struct {
	Name  string `json:"name"`
	Items []T    `json:"items"`
}

type Collectors[T constraints.Ordered] []Collector[T]

// FindCommon returns items that are on more than one collector's shelf.
func (c Collectors[T]) FindCommon() []T {
	// Register all items on shelves.
	itemsCount := c.countItems()

	// List containing all the items that were sitting on at least two shelves.
	var itemsInCommon []T

	// Find items that were added to shelve more than once.
	for item, count := range itemsCount {
		if count > 1 {
			itemsInCommon = append(itemsInCommon, item)
		}
	}

	// Sort allows us to be deterministic, sorted by the
	return sortItems(itemsInCommon)
}

// count registers all the books and their occurrences from the collectors shelves.
func (c Collectors[C]) countItems() map[C]uint {
	itemCount := make(map[C]uint)

	for _, coll := range c {
		for _, item := range coll.Items {
			// If a collector has two copies, that counts as two.
			itemCount[item]++
		}
	}

	return itemCount
}

func sortItems[T comparable](items []T) {
	func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	}
}
