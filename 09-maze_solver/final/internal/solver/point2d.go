package solver

type point2d struct {
	x int
	y int
}

func (p point2d) neighbours() []point2d {
	return []point2d{
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x - 1, p.y},
	}
}
