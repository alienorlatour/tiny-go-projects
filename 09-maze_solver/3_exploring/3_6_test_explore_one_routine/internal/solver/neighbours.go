package solver

import "image"

// neighbours returns an array of the 4 neighbours of a pixel.
// Some returned positions may be outside the image.
func neighbours(p image.Point) []image.Point {
	return []image.Point{
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
	}
}
