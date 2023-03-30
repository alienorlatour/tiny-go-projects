package patterns

import (
	"fmt"

	"learngo-pockets/genericworms/collectors"
)

// Pattern describes a pattern on a collector's shelf.
// A pattern is fully identified by its craft and its name. Yardage is merely informative.
type Pattern struct {
	Craft   string `json:"craft"`
	Name    string `json:"name"`
	Yardage int    `json:"yardage"`
}

func (pattern Pattern) Before(sortable collectors.Sortable) bool {
	other, ok := sortable.(Pattern)
	if !ok {
		return false
	}

	if pattern.Craft != other.Craft {
		return pattern.Craft < other.Craft
	}

	return pattern.Name < other.Name
}

// String implements the Stringer interface.
func (pattern Pattern) String() string {
	return fmt.Sprintf("Craft: %10s; Name: %20s; Yardage: %d yards\n", pattern.Craft, pattern.Name, pattern.Yardage)
}
