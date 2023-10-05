package solver

import "image"

type path struct {
	previousStep *path
	at           image.Point
}
