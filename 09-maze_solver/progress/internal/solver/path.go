package solver

import "image"

type path struct {
	previousSteps *path
	at            image.Point
}
