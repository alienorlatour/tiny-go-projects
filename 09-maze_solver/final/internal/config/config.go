package config

import "image/color"

// Config contains the colours of the different types of pixels in our maze
type Config struct {
	WallColour  color.RGBA
	PathColour  color.RGBA
	StartColour color.RGBA
	EndColour   color.RGBA
}

// Get returns the configuration of our maze
func Get() Config {
	return Config{
		WallColour:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
		PathColour:  color.RGBA{R: 255, G: 255, B: 255, A: 255},
		StartColour: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		EndColour:   color.RGBA{R: 255, G: 0, B: 0, A: 255},
	}
}

func (c Config) IsStart(pixel color.RGBA) bool {
	return pixel == c.StartColour
}

func (c Config) IsEnd(pixel color.RGBA) bool {
	return pixel == c.EndColour
}

func (c Config) IsWall(pixel color.RGBA) bool {
	return pixel == c.WallColour
}

func (c Config) IsPath(pixel color.RGBA) bool {
	return pixel == c.PathColour
}
