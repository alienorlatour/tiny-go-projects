package config

import "image/color"

// Config contains the colours of the different types of pixels in our maze
type Config struct {
	WallColour     color.RGBA
	PathColour     color.RGBA
	StartColour    color.RGBA
	EndColour      color.RGBA
	SolutionColour color.RGBA
	ExploredColour color.RGBA
}

// Get returns the configuration of our maze
func Get() Config {
	return Config{
		WallColour:     color.RGBA{R: 0, G: 0, B: 0, A: 255},
		PathColour:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
		StartColour:    color.RGBA{R: 0, G: 255, B: 0, A: 255},
		EndColour:      color.RGBA{R: 255, G: 0, B: 0, A: 255},
		SolutionColour: color.RGBA{R: 255, G: 128, B: 0, A: 255},
		ExploredColour: color.RGBA{R: 0, G: 128, B: 255, A: 255},
	}
}
