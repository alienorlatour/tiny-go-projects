package solver

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log/slog"
	"os"
	"strings"
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

	defer func() {
		if closeErr := fd.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()

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
func (s *Solver) SaveSolution(outputPath string) error {
	err := s.saveSolution(outputPath)
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

func (s *Solver) saveSolution(outputPath string) (err error) {
	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}

	defer func() {
		if closeErr := fd.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("unable to close file: %w", closeErr))
		}
	}()

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s: %w", outputPath, err)
	}

	return nil
}

// saveAnimation writes the gif file.
func (s *Solver) saveAnimation(gifPath string) (err error) {
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

	slog.Info(fmt.Sprintf("animation contains %d frames", len(s.animation.Image)))
	err = gif.EncodeAll(outputImage, s.animation)
	if err != nil {
		return fmt.Errorf("unable to encode gif: %w", err)
	}

	return nil
}
