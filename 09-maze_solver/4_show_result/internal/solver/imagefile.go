package solver

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
)

// openMaze opens a RGBA png image from a path.
func openMaze(imagePath string) (*image.RGBA, error) {
	_, err := os.Stat(imagePath)
	if err != nil {
		return nil, fmt.Errorf("unable to check input file %s: %w", imagePath, err)
	}

	fd, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open input image at %s: %w", imagePath, err)
	}
	defer fd.Close()

	img, err := png.Decode(fd)
	if err != nil {
		return nil, fmt.Errorf("unable to load input image from %s: %w", imagePath, err)
	}

	// Using RGBAAt() instead of At() saves a lot of time, but it requires a *image.RGBA
	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("this isn't a RGBA image")
	}

	return rgbaImage, nil
}

// SaveSolution saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outputPath string) (err error) {
	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}
	defer func() {
		if closeErr := fd.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()

	stepsFromTreasure := s.solution
	// Paint the path from last position (treasure) back to first position (entrance).
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.at.X, stepsFromTreasure.at.Y, s.config.solutionColour)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s: %w", outputPath, err)
	}

	return nil
}
