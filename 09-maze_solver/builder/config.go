package main

import "image/color"

// config contains the colours of the different types of pixels in our maze.
type config struct {
	wallColour     color.RGBA
	pathColour     color.RGBA
	entranceColour color.RGBA
	treasureColour color.RGBA
}

// defaultColours returns the colour configuration of our maze.
func defaultColours() config {
	return config{
		wallColour:     color.RGBA{R: 0, G: 0, B: 0, A: 255},       // black
		pathColour:     color.RGBA{R: 255, G: 255, B: 255, A: 255}, // white
		entranceColour: color.RGBA{R: 0, G: 191, B: 255, A: 255},   // deep sky blue
		treasureColour: color.RGBA{R: 255, G: 0, B: 128, A: 255},   // pink
	}
}
