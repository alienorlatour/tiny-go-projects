package collectors

import (
	"fmt"
	"io"
	"sort"
)

// A Collector contains the list of books on a collector's shelf.
type Collector[T collectible] struct {
	Name  string `json:"name"`
	Items []T    `json:"items"`
}

type collectible interface {
	comparable
	Sortable
}

// Sortable exposes a way of telling if an item should appear before another.
// Sortable needs to be exposed so that we can use Before elsewhere.
type Sortable interface {
	Before(Sortable) bool
}

// Collectors is a list of collectors.
type Collectors[T collectible] []Collector[T]

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

	// Sort allows us to be deterministic.
	return sortItems(itemsInCommon)
}

// Display prints the collectible items nicely.
func Display[T collectible](w io.Writer, items []T) {
	for _, item := range items {
		_, _ = fmt.Fprintf(w, "- %s\n", item)
	}
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

func sortItems[T collectible](items []T) []T {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Before(items[j])
	})

	return items
}
