package solver

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log/slog"
	"os"
	"time"
)

func (s *Solver) draw() {
	ticker := time.NewTicker(50 * time.Nanosecond)
	for {
		if s.solution != nil {
			break
		}

		// wait for a tick
		select {
		case <-ticker.C:
		}
		go s.drawCurrentFrameToGif()
	}

	// add solution
	s.drawCurrentFrameToGif()
}

func (s *Solver) drawCurrentFrameToGif() {
	bounds := s.maze.Bounds()
	frame := image.NewPaletted(image.Rect(0, 0, bounds.Dx(), bounds.Dy()), palette.Plan9)

	// Convert RGBA to paletted
	draw.FloydSteinberg.Draw(frame, s.maze.Bounds(), s.maze, image.Point{0, 0})

	s.gif.Image = append(s.gif.Image, frame)
}

func (s *Solver) saveGif(gifPath string) error {
	outputImage, err := os.Create(gifPath)
	if err != nil {
		return fmt.Errorf("unable to create output gif at %s: %w", outputImage, err)
	}

	defer outputImage.Close()

	s.gif.Delay = make([]int, len(s.gif.Image))
	for i := 0; i < len(s.gif.Delay)-1; i++ {
		s.gif.Delay[i] = 5
	}
	slog.Info(fmt.Sprintf("gif contains %d frames", len(s.gif.Image)))
	s.gif.Delay[len(s.gif.Delay)-1] = 50
	err = gif.EncodeAll(outputImage, s.gif)
	if err != nil {
		return fmt.Errorf("unable to encode gif: %w", err)
	}

	return nil
}
