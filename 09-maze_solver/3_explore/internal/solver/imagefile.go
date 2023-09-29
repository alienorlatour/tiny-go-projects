package solver

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// openMaze opens a RGBA png image from a path.
func openMaze(inputPath string) (*image.RGBA, error) {
	_, err := os.Stat(inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to check input file %s: %w", inputPath, err)
	}

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

	return rgbaImage, nil
}

// SaveSolution saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outputPath string) error {
	return nil
}
