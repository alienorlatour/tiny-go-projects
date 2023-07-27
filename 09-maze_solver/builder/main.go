package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"math/rand"
	"os"
)

func main() {
	maze := generateMaze(60, 40)
	saveToPNG(maze, "maze.png")
}

func generateMaze(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	entry := getRandomEdgeNotCorner(width, height)
	img.Set(entry.x, entry.y, color.White)

	// draw the path
	p := entry
	// create a massive channel, because I don't want to start a listener right now.
	b := &builder{allowEdge: true, ps: make(chan pos, 2000)}

	slog.Info(fmt.Sprintf("Start at %v\n", p))
	for {
		// look for eligible places
		nextPositions := b.candidates(img, p)
		if len(nextPositions) == 0 {
			break
		}
		p = nextPositions[rand.Intn(len(nextPositions))]
		img.Set(p.x, p.y, color.White)
		b.ps <- p

		if p.x == 0 || p.x == width-1 || p.y == 0 || p.y == height-1 {
			// we've reached the border - this is the exit now
			b.allowEdge = false
			b.exit = p
			break
		}
	}

	b.completeMaze(img)

	if b.exit.x == entry.x || b.exit.y == entry.y {
		return generateMaze(width, height)
	}

	return img
}

func saveToPNG(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}

func getRandomEdgeNotCorner(width, height int) pos {
	side := rand.Intn(4)
	switch side {
	case 0:
		// top
		return pos{rand.Intn(width-2 /*ignore corners*/) + 1, 0}
	case 1:
		// right
		return pos{width - 1, rand.Intn(height-2 /*ignore corners*/) + 1}
	case 2:
		// bottom
		return pos{rand.Intn(width-2 /*ignore corners*/) + 1, height - 1}
	case 3:
		// left
		return pos{0, rand.Intn(height-2 /*ignore corners*/) + 1}
	default:
		return pos{0, 0}
	}
}

type pos struct {
	x int
	y int
}

type builder struct {
	allowEdge bool
	ps        chan pos
	exit      pos
}

func (b *builder) candidates(img image.Image, p pos) []pos {
	// we are in the center of 5x5 grid. We can't go to a neighbour of a white pixel.
	// Since we reached this position from a pixel, we only need to evaluate the exterior ring
	/*
	   a b c d e
	   f g h i j
	   k l X m n
	   o p q r s
	   t u v w x
	*/
	// In order to go from X to h, we need:
	// - h is black
	// - b, c, d, g, i are black
	eligible := make([]pos, 0)
	width := img.Bounds().Dx() - 1
	height := img.Bounds().Dy() - 1

	if /* h */ p.y > 0 && img.At(p.x, p.y-1) == rgbaBlack {
		if /* g */ (p.x > 0 && img.At(p.x-1, p.y-1) == rgbaBlack) &&
			/* i */ (p.x < width && img.At(p.x+1, p.y-1) == rgbaBlack) &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			(b.allowEdge && (p.y == 1) ||
				/* c */ (p.y > 1 && (img.At(p.x, p.y-2) == rgbaBlack) &&
					/* b */ (p.x > 0 && img.At(p.x-1, p.y-2) == rgbaBlack) &&
					/* d */ (p.x < width && img.At(p.x+1, p.y-2) == rgbaBlack))) {
			eligible = append(eligible, pos{p.x, p.y - 1})
		}
	}

	if /* q */ p.y < height && img.At(p.x, p.y+1) == rgbaBlack {
		if /* p */ (p.x > 0 && img.At(p.x-1, p.y+1) == rgbaBlack) &&
			/* r */ (p.x < width && img.At(p.x+1, p.y+1) == rgbaBlack) &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			(b.allowEdge && (p.y == height-1) ||
				/* v */ (p.y < height-1 && (img.At(p.x, p.y+2) == rgbaBlack) &&
					/* u */ (p.x > 0 && img.At(p.x-1, p.y+2) == rgbaBlack) &&
					/* w */ (p.x < width && img.At(p.x+1, p.y+2) == rgbaBlack))) {
			eligible = append(eligible, pos{p.x, p.y + 1})
		}
	}

	if /* l */ p.x > 0 && img.At(p.x-1, p.y) == rgbaBlack {
		if /* g */ (p.y > 0 && img.At(p.x-1, p.y-1) == rgbaBlack) &&
			/* p */ (p.y < height && img.At(p.x-1, p.y+1) == rgbaBlack) &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			(b.allowEdge && (p.x == 1) ||
				/* k */ (p.x > 1 && (img.At(p.x-2, p.y) == rgbaBlack) &&
					/* f */ (p.y > 0 && img.At(p.x-2, p.y-1) == rgbaBlack) &&
					/* o */ (p.y < height && img.At(p.x-2, p.y+1) == rgbaBlack))) {
			eligible = append(eligible, pos{p.x - 1, p.y})
		}
	}

	if /* m */ p.x < width && img.At(p.x+1, p.y) == rgbaBlack {
		if /* i */ (p.y > 0 && img.At(p.x+1, p.y-1) == rgbaBlack) &&
			/* r */ (p.y < height && img.At(p.x+1, p.y+1) == rgbaBlack) &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			(b.allowEdge && (p.x == width-1) ||
				/* n */ (p.x < width-1 && (img.At(p.x+2, p.y) == rgbaBlack) &&
					/* j */ (p.y > 0 && img.At(p.x+2, p.y-1) == rgbaBlack) &&
					/* s */ (p.y < height && img.At(p.x+2, p.y+1) == rgbaBlack))) {
			eligible = append(eligible, pos{p.x + 1, p.y})
		}
	}

	return eligible
}

func (b *builder) completeMaze(img *image.RGBA) {
	for p := range b.ps {
		newPos := p
		for {
			nextPositions := b.candidates(img, newPos)
			if len(nextPositions) == 0 {
				break
			}
			newPos = nextPositions[rand.Intn(len(nextPositions))]
			img.Set(newPos.x, newPos.y, color.White)
			b.ps <- newPos

			if newPos.x == 0 || newPos.x == img.Rect.Dx()-1 || newPos.y == 0 || newPos.y == img.Rect.Dy()-1 {
				// we've reached the border - this is the exit now
				b.allowEdge = false
				b.exit = newPos
				break
			}
		}

		if len(b.ps) == 0 {
			close(b.ps)
		}
	}
}

var rgbaBlack = color.RGBA{}
