package solver

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Solver struct {
	maze *image.RGBA

	solution []point2d
}

type point2d struct {
	x int
	y int
}

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

	s := &Solver{maze: rgbaImage}

	return s, nil
}

// Solve finds the path from one end to the other.
func (s *Solver) Solve() error {
	return nil
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

	var pathColour = color.RGBA{255, 0, 255, 255}

	for _, p := range s.solution {
		s.maze.Set(p.x, p.y, pathColour)
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s", outputPath)
	}

	return nil
}
