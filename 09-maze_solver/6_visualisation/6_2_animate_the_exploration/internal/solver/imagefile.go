package solver

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log"
	"os"
	"strings"
)

// openMaze opens a RGBA png image from a path.
func openMaze(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open image %s: %w", imagePath, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("unable to load input image from %s: %w", imagePath, err)
	}

	// Using RGBAAt() instead of At() saves a lot of time, but it requires a *image.RGBA
	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("expected RGBA image, got %T", img)
	}

	return rgbaImage, nil
}

// SaveSolution saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outputPath string) (err error) {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()

	stepsFromTreasure := s.solution
	// Paint the path from last position (treasure) back to first position (entrance).
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.at.X, stepsFromTreasure.at.Y, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(f, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s: %w", outputPath, err)
	}

	gifPath := strings.Replace(outputPath, "png", "gif", -1)
	err = s.saveAnimation(gifPath)
	if err != nil {
		return fmt.Errorf("unable to write output animation at %s", gifPath)
	}

	return nil
}

// saveAnimation writes the gif file.
func (s *Solver) saveAnimation(gifPath string) error {
	outputImage, err := os.Create(gifPath)
	if err != nil {
		return fmt.Errorf("unable to create output gif at %s: %w", gifPath, err)
	}

	defer func() {
		if closeErr := outputImage.Close(); closeErr != nil {
			// Return err and closeErr, in worst case scenario.
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()

	log.Printf("animation contains %d frames\n", len(s.animation.Image))
	err = gif.EncodeAll(outputImage, s.animation)
	if err != nil {
		return fmt.Errorf("unable to encode gif: %w", err)
	}

	return nil
}
