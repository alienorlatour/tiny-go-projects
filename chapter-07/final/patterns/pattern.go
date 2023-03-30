package patterns

import (
	"sort"

	"learngo-pockets/genericworms/collectors"
)

// Pattern describes a pattern on a collector's shelf.
// A pattern is fully identified by its craft and its name. Yardage is merely informative.
type Pattern struct {
	Craft   string `json:"craft"`
	Name    string `json:"name"`
	Yardage int    `json:"yardage"`
}

// Collectors describe a list of pattern collectors and their pattern
type Collectors collectors.Collectors[Pattern]

// FindCommon return the patterns in common, sorted first by craft, title, needle size, and then yardage.
func (colls Collectors) FindCommon() []Pattern {
	// We need a Collectors[T] here
	patternsInCommon := collectors.Collectors[Pattern](colls).FindCommon()

	// sort.Sort sorts the slice in place.
	// We can instantiate a slice of the type byCraft, which implements sort.Interface.
	sort.Sort(byCraft(patternsInCommon))

	return patternsInCommon
}

// byCraft implements sort.Interface for a list of patterns.
type byCraft []Pattern

// Len implements sort.Interface by returning the length of Books.
func (b byCraft) Len() int { return len(b) }

// Less returns true if Craft_i is before Craft_j in alphabetical order.
func (b byCraft) Less(i, j int) bool {
	if b[i].Craft != b[j].Craft {
		return b[i].Craft < b[j].Craft
	}

	// Since the pair of Craft and Name create a unique pattern, we don't need to check further fields.
	return b[i].Name < b[j].Name
}

// Swap implements sort.Interface and swaps two books.
func (b byCraft) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
