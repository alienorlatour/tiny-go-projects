package solver

import "image"

type Path struct {
	PreviousSteps *Path
	At            image.Point
}
