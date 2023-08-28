package solver

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"09-maze_solver/final/internal/config"
)

type Solver struct {
	maze *image.RGBA

	solution []point2d
	config   config.Config

	pathsToExplore chan []point2d

	quit chan struct{}
	b    broadcaster[struct{}]

	exploredCount int
	toPaint       chan point2d
	mutex         sync.Mutex

	gif *gif.GIF
}

// New returns a solver on a RGBA png image
func New(inputPath string) (*Solver, error) {
	// Check input file
	_, err := os.Stat(inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to check input file %s: %w", inputPath, err)
	}

	// load image
	fd, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open input image at %s: %w", inputPath, err)
	}
	defer fd.Close()

	img, err := png.Decode(fd)
	if err != nil {
		return nil, fmt.Errorf("unable to load input image from %s: %w", inputPath, err)
	}

	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("this isn't a RGBA image")
	}

	s := &Solver{
		maze:           rgbaImage,
		config:         config.Get(),
		pathsToExplore: make(chan []point2d, 10),
		toPaint:        make(chan point2d),
		gif:            &gif.GIF{},
		quit:           make(chan struct{}, poolSize),
	}

	s.b = broadcaster[struct{}]{
		out:   s.quit,
		count: poolSize,
	}

	return s, nil
}

// Solve finds the path from one end to the other.
func (s *Solver) Solve() error {
	now := time.Now()
	start, end, err := s.findExtremities()
	if err != nil {
		return fmt.Errorf("unable to find extremities: %w", err)
	}

	slog.Info(fmt.Sprintf("starting at %v, ending at %v", start, end))

	// We know the first pixel is on the left edge.
	s.pathsToExplore <- []point2d{start, {1, start.y}}

	go s.paintExplored()
	s.listenToBranches()

	slog.Info(fmt.Sprintf("It took %d nanoseconds to solve the maze", time.Since(now).Nanoseconds()))
	return nil
}

// findExtremities returns the position of the extremities on the image.
func (s *Solver) findExtremities() (start, end point2d, err error) {
	// We know the extremities are on the edge.

	width, height := s.maze.Bounds().Dx()-1, s.maze.Bounds().Dy()-1

	// Scan the vertical edges
	for y := 1; y <= height-1; y++ {
		// check the left edge
		switch s.maze.RGBAAt(0, y) {
		case s.config.StartColour:
			start = point2d{0, y}
		case s.config.EndColour:
			end = point2d{0, y}
		}

		// check the right edge
		switch s.maze.RGBAAt(width, y) {
		case s.config.StartColour:
			start = point2d{height, y}
		case s.config.EndColour:
			end = point2d{height, y}
		}
	}

	// Scan the horizontal edges
	for x := 1; x <= width-1; x++ {
		// check the top edge
		switch s.maze.RGBAAt(x, 0) {
		case s.config.StartColour:
			start = point2d{x, 0}
		case s.config.EndColour:
			end = point2d{x, 0}
		}

		// check the bottom edge
		switch s.maze.RGBAAt(x, height) {
		case s.config.StartColour:
			start = point2d{x, height}
		case s.config.EndColour:
			end = point2d{x, height}
		}
	}

	origin := point2d{}
	switch {
	case start == end:
		return start, end, fmt.Errorf("start and end at same positions: %v", start)
	case start == origin:
		return start, end, fmt.Errorf("start position not found")
	case end == origin:
		return start, end, fmt.Errorf("end position not found")
	}

	return
}

// SaveSolution saves the image as a PNG file with the solution path in red.
func (s *Solver) SaveSolution(outputPath string) error {
	// check output file
	_, err := os.Stat(outputPath)
	switch {
	case err == nil:
		return fmt.Errorf("output file %s already exists", outputPath)
	case !os.IsNotExist(err):
		return fmt.Errorf("unable to check output file %s: %w", outputPath, err)
	}

	for _, p := range s.solution {
		s.maze.Set(p.x, p.y, s.config.SolutionColour)
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s", outputPath)
	}

	err = s.saveGif(strings.Replace(outputPath, "png", "gif", -1))
	if err != nil {
		return fmt.Errorf("can't save gif")
	}

	return nil
}