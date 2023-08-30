package solver

// point2d represents the position of one pixel.
type point2d struct {
	x int
	y int
}

// neighbours returns an array of the 4 neighbours of a pixel.
// Some of the returned positions may be outside the image.
func (p point2d) neighbours() []point2d {
	return []point2d{
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x - 1, p.y},
	}
}
