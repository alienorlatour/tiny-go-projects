package solver

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// openImage returns a RGBA png image.
func openImage(inputPath string) (*image.RGBA, error) {
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

// SaveSolution saves the image as a PNG file with the solution path in red.
func (s *Solver) SaveSolution(outputPath string) error {
	_, err := os.Stat(outputPath)
	switch {
	case err == nil:
		return fmt.Errorf("output file %s already exists", outputPath)
	case !os.IsNotExist(err):
		return fmt.Errorf("unable to check output file %s: %w", outputPath, err)
	}

	for _, p := range s.solution {
		s.maze.Set(p.x, p.y, s.config.solutionColour)
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}
	defer fd.Close()

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s", outputPath)
	}

	return nil
}
