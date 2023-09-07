package solver

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log/slog"
	"os"
	"strings"
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

// SaveSolution saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outputPath string) error {
	_, err := os.Stat(outputPath)
	switch {
	case err == nil:
		return fmt.Errorf("output file %s already exists", outputPath)
	case !os.IsNotExist(err):
		return fmt.Errorf("unable to check output file %s: %w", outputPath, err)
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}
	defer fd.Close()

	for _, p := range s.solution {
		s.maze.Set(p.X, p.Y, s.config.solutionColour)
	}

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s", outputPath)
	}

	err = s.saveAnimation(strings.Replace(outputPath, "png", "gif", -1))
	return nil
}

func (s *Solver) saveAnimation(gifPath string) error {
	outputImage, err := os.Create(gifPath)
	if err != nil {
		return fmt.Errorf("unable to create output gif at %s: %w", outputImage, err)
	}

	defer outputImage.Close()

	// make sure the solution frame is present in the GIF
	s.drawCurrentFrameToGif()

	slog.Info(fmt.Sprintf("animation contains %d frames", len(s.animation.Image)))
	s.animation.Delay[len(s.animation.Delay)-1] = 300
	err = gif.EncodeAll(outputImage, s.animation)
	if err != nil {
		return fmt.Errorf("unable to encode gif: %w", err)
	}

	return nil
}
