package solver

import "image/color"

// config contains the colours of the different types of pixels in our maze.
type config struct {
	wallColour     color.RGBA
	pathColour     color.RGBA
	entranceColour color.RGBA
	treasureColour color.RGBA
	solutionColour color.RGBA
	exploredColour color.RGBA
}

// Get returns the configuration of our maze
func defaultColours() config {
	return config{
		wallColour:     color.RGBA{R: 0, G: 0, B: 0, A: 255},
		pathColour:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
		entranceColour: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		treasureColour: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		solutionColour: color.RGBA{R: 255, G: 128, B: 0, A: 255},
		exploredColour: color.RGBA{R: 0, G: 128, B: 255, A: 255},
	}
}
