package collectors

import (
	"sort"
)

// A Collector contains the list of books on a collector's shelf.
type Collector[T lesser] struct {
	Name  string `json:"name"`
	Items []T    `json:"items"`
}

type Collectors[T lesser] []Collector[T]

type lesser interface {
	comparable
	// Less(other lesser) bool
}

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

func sortItems[T lesser](items []T) []T {
	sort.Slice(items, func(i, j int) bool {
		// return items[i].Less(items[j])
		return true
	})

	return items
}
