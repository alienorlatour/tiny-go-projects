package solver

import "image"

// path represents a route from the entrance of the maze up to a position.
type path struct {
	previousStep *path       // nolint:unused
	at           image.Point // nolint:unused
}

// isPreviousStep returns true if the given point is the previous position of the path.
// nolint:unused
func (p path) isPreviousStep(n image.Point) bool {
	return p.previousStep != nil && p.previousStep.at == n
}
