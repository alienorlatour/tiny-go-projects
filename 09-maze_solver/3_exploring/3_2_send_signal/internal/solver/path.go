package solver

import "image"

// path represents a route from the entrance of the maze up to a position.
type path struct {
	previousSteps *path
	at            image.Point
}
