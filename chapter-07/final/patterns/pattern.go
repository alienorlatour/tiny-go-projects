package patterns

import (
	"fmt"
	"io"
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
	sort.Slice(patternsInCommon, func(i, j int) bool {
		if patternsInCommon[i].Craft != patternsInCommon[j].Craft {
			return patternsInCommon[i].Craft < patternsInCommon[j].Craft
		}

		// Since the pair of Craft and Name create a unique pattern, we don't need to check further fields.
		return patternsInCommon[i].Name < patternsInCommon[j].Name
	})

	return patternsInCommon
}

// Display prints out the titles and authors of a list of books
func Display(w io.Writer, patterns []Pattern) {
	for _, pattern := range patterns {
		_, _ = fmt.Fprintf(w, "Craft: %10s; Name: %20s; Yardage: %d yards\n", pattern.Craft, pattern.Name, pattern.Yardage)
	}
}
