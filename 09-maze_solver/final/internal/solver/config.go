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

// defaultColours returns the colour configuration of our maze.
func defaultColours() config {
	return config{
		wallColour:     color.RGBA{R: 0, G: 0, B: 0, A: 255},       // Black
		pathColour:     color.RGBA{R: 255, G: 255, B: 255, A: 255}, // White
		entranceColour: color.RGBA{R: 0, G: 191, B: 255, A: 255},   // Deep sky blue
		treasureColour: color.RGBA{R: 255, G: 0, B: 128, A: 255},   // Pink
		solutionColour: color.RGBA{R: 225, G: 140, B: 0, A: 255},   // Orange
		exploredColour: color.RGBA{R: 0, G: 128, B: 255, A: 255},   // Bright Blue
	}
}
