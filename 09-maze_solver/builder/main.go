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
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])

	maze := generateMaze(width, height)
	saveToPNG(maze, "maze.png")
}

func generateMaze(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//	entry := getRandomEdgeNotCorner(width, height)
	entry := pos{0, width / 2}
	img.Set(entry.x, entry.y, color.White)

	// draw the path
	p := posWithCount{entry, 0}
	// create a massive channel, because I don't want to start a listener right now.
	b := &builder{ps: make(chan posWithCount, width*height), width: width - 1, height: height - 1, complexity: max(width, height)}

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
			b.exit = &p
			break
		}
	}
	saveToPNG(img, "tmp.png")
	b.completeMaze(img)

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
	if p.x >= bldr.width-1 {
		return false
	}
	if p.y >= bldr.height-1 {
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

	if /* h */ img.At(h.x, h.y) == rgbaBlack {
		if /* g */ img.At(g.x, g.y) == rgbaBlack &&
			/* i */ img.At(i.x, i.y) == rgbaBlack &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.y == 1) || (bldr.isInside(h) &&
				/* c */ img.At(c.x, c.y) == rgbaBlack &&
				/* b */ img.At(b.x, b.y) == rgbaBlack &&
				/* d */ img.At(d.x, d.y) == rgbaBlack)) {
			eligible = append(eligible, posWithCount{h, pwc.count + 1})
		}
	}

	if /* q */ img.At(q.x, q.y) == rgbaBlack {
		if /* p */ img.At(p.x, p.y) == rgbaBlack &&
			/* r */ img.At(r.x, r.y) == rgbaBlack &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.y == height-1) || (bldr.isInside(q) &&
				/* v */ img.At(v.x, v.y) == rgbaBlack &&
				/* u */ img.At(u.x, u.y) == rgbaBlack &&
				/* w */ img.At(w.x, w.y) == rgbaBlack)) {
			eligible = append(eligible, posWithCount{q, pwc.count + 1})
		}
	}

	if /* l */ img.At(l.x, l.y) == rgbaBlack {
		if /* g */ img.At(g.x, g.y) == rgbaBlack &&
			/* p */ img.At(p.x, p.y) == rgbaBlack &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.x == 1) || (bldr.isInside(l) &&
				/* k */ img.At(k.x, k.y) == rgbaBlack &&
				/* f */ img.At(f.x, f.y) == rgbaBlack &&
				/* o */ img.At(o.x, o.y) == rgbaBlack)) {
			eligible = append(eligible, posWithCount{l, pwc.count + 1})
		}
	}

	if /* m */ img.At(m.x, m.y) == rgbaBlack {
		if /* i */ img.At(i.x, i.y) == rgbaBlack &&
			/* r */ img.At(r.x, r.y) == rgbaBlack &&
			// if we still allow edge, then we can venture in there. otherwise, it's OK to ignore it
			((bldr.allowExit(pwc) && pwc.x == width-1) || (bldr.isInside(m) &&
				/* n */ img.At(n.x, n.y) == rgbaBlack &&
				/* j */ img.At(j.x, j.y) == rgbaBlack &&
				/* s */ img.At(s.x, s.y) == rgbaBlack)) {
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
			img.Set(newPos.x, newPos.y, color.White)
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

var rgbaBlack = color.RGBA{}
