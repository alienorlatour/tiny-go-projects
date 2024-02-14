// nolint:unused
package solver

import "image"

// path represents a route from the entrance of the maze up to a position.
type path struct {
	previousStep *path
	at           image.Point
}
