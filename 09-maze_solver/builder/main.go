package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
)

// go run main.go WIDTH HEIGHT
func main() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("Something horrible happened: %s\n", e)
		}
	}()

	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])

	maze := generateMaze(width, height)
	saveToPNG(maze, "maze.png")
}

func generateMaze(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// I see a red door...
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, wall)
		}
	}

	entry := pos{0, height / 2}
	img.Set(entry.x, entry.y, path)

	// draw the path
	p := posWithCount{entry, 0}

	const complexity = 2 // change this for easier / creatable mazes, or harder ones.
	// create a massive channel, because I don't want to start a listener right now.
	b := &builder{ps: make(chan posWithCount, width*height), width: width - 1, height: height - 1, complexity: complexity * (width + height)}

	for {
		// look for eligible places
		nextPositions := b.candidates(img, p)
		if len(nextPositions) == 0 {
			break
		}
		p = nextPositions[rand.Intn(len(nextPositions))]
		img.Set(p.x, p.y, path)
		b.ps <- p

		if p.x == 0 || p.x == width-1 || p.y == 0 || p.y == height-1 {
			// we've reached the border - this is the exit now
			b.exit = &p
			break
		}
	}

	b.completeMaze(img)

	img.Set(entry.x, entry.y, color.RGBA{0, 255, 0, 255})
	img.Set(b.exit.x, b.exit.y, color.RGBA{255, 0, 0, 255})
	slog.Info(fmt.Sprintf("Start at %v\n", entry))
	slog.Info(fmt.Sprintf("End at %v\n", b.exit))
	slog.Info(fmt.Sprintf("total length: %d\n", b.exit.count))
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

type pos struct {
	x int
	y int
}

type posWithCount struct {
	pos
	count int
}

type builder struct {
	ps            chan posWithCount
	exit          *posWithCount
	width, height int
	complexity    int
}

func (bldr *builder) allowExit(p posWithCount) bool {
	if bldr.exit != nil {
		return false
	}

	if p.count < bldr.complexity {
		return false
	}
	return true
}

func (bldr *builder) isInside(p pos) bool {
	if p.x <= 0 {
		return false
	}
	if p.y <= 0 {
		return false
	}
	if p.x >= bldr.width {
		return false
	}
	if p.y >= bldr.height {
		return false
	}
	return true
}

func (bldr *builder) candidates(img image.Image, pwc posWithCount) []posWithCount {
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
	// - b, c, d, g, and i are black
	eligible := make([]posWithCount, 0)
	width := bldr.width
	height := bldr.height

	b := pos{pwc.x - 1, pwc.y - 2}
	c := pos{pwc.x, pwc.y - 2}
	d := pos{pwc.x + 1, pwc.y - 2}
	f := pos{pwc.x - 2, pwc.y - 1}
	g := pos{pwc.x - 1, pwc.y - 1}
	h := pos{pwc.x, pwc.y - 1}
	i := pos{pwc.x + 1, pwc.y - 1}
	j := pos{pwc.x + 2, pwc.y - 1}
	k := pos{pwc.x - 2, pwc.y}
	l := pos{pwc.x - 1, pwc.y}
	m := pos{pwc.x + 1, pwc.y}
	n := pos{pwc.x + 2, pwc.y}
	o := pos{pwc.x - 2, pwc.y + 1}
	p := pos{pwc.x - 1, pwc.y + 1}
	q := pos{pwc.x, pwc.y + 1}
	r := pos{pwc.x + 1, pwc.y + 1}
	s := pos{pwc.x + 2, pwc.y + 1}
	u := pos{pwc.x - 1, pwc.y + 2}
	v := pos{pwc.x, pwc.y + 2}
	w := pos{pwc.x + 1, pwc.y + 2}

	if /* h */ img.At(h.x, h.y) == wall {
		if /* g */ img.At(g.x, g.y) == wall &&
			/* i */ img.At(i.x, i.y) == wall &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.y == 1) || (bldr.isInside(h) &&
				/* c */ img.At(c.x, c.y) == wall &&
				/* b */ img.At(b.x, b.y) == wall &&
				/* d */ img.At(d.x, d.y) == wall)) {
			eligible = append(eligible, posWithCount{h, pwc.count + 1})
		}
	}

	if /* q */ img.At(q.x, q.y) == wall {
		if /* p */ img.At(p.x, p.y) == wall &&
			/* r */ img.At(r.x, r.y) == wall &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.y == height-1) || (bldr.isInside(q) &&
				/* v */ img.At(v.x, v.y) == wall &&
				/* u */ img.At(u.x, u.y) == wall &&
				/* w */ img.At(w.x, w.y) == wall)) {
			eligible = append(eligible, posWithCount{q, pwc.count + 1})
		}
	}

	if /* l */ img.At(l.x, l.y) == wall {
		if /* g */ img.At(g.x, g.y) == wall &&
			/* p */ img.At(p.x, p.y) == wall &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.x == 1) || (bldr.isInside(l) &&
				/* k */ img.At(k.x, k.y) == wall &&
				/* f */ img.At(f.x, f.y) == wall &&
				/* o */ img.At(o.x, o.y) == wall)) {
			eligible = append(eligible, posWithCount{l, pwc.count + 1})
		}
	}

	if /* m */ img.At(m.x, m.y) == wall {
		if /* i */ img.At(i.x, i.y) == wall &&
			/* r */ img.At(r.x, r.y) == wall &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.x == width-1) || (bldr.isInside(m) &&
				/* n */ img.At(n.x, n.y) == wall &&
				/* j */ img.At(j.x, j.y) == wall &&
				/* s */ img.At(s.x, s.y) == wall)) {
			eligible = append(eligible, posWithCount{m, pwc.count + 1})
		}
	}

	return eligible
}

func (bldr *builder) completeMaze(img *image.RGBA) {
	for p := range bldr.ps {
		newPos := p
		for {
			nextPositions := bldr.candidates(img, newPos)
			if len(nextPositions) == 0 {
				break
			}
			newPos = nextPositions[rand.Intn(len(nextPositions))]
			img.Set(newPos.x, newPos.y, path)
			bldr.ps <- newPos

			if newPos.x == 0 || newPos.x == bldr.width || newPos.y == 0 || newPos.y == bldr.height {
				// we've reached the border - this is the exit now
				bldr.exit = &newPos
				break
			}
		}

		if len(bldr.ps) == 0 {
			close(bldr.ps)
		}
	}
}

var (
	wall = color.RGBA{0, 0, 0, 255}
	path = color.RGBA{255, 255, 255, 255}
)
