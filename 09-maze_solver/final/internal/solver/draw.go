package solver

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"log/slog"
	"os"

	"golang.org/x/image/draw"
)

func (s *Solver) paint() {
	for pos := range s.explored {
		s.paintAt(pos)
	}
}

func (s *Solver) paintAt(pos point2d) {
	slog.Info(fmt.Sprintf("painting at %v", pos))
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.maze.RGBAAt(pos.x, pos.y) == s.config.PathColour {
		s.maze.SetRGBA(pos.x, pos.y, s.config.ExploredColour)
	}

	s.exploredCount++

	if s.exploredCount%20 == 0 {
		s.drawCurrentFrameToGif()
	}
}

func (s *Solver) drawCurrentFrameToGif() {
	frame := image.NewPaletted(image.Rect(0, 0, 500, 500), palette.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.gif.Image = append(s.gif.Image, frame)
}

func (s *Solver) saveGif(gifPath string) error {
	outputImage, err := os.Create(gifPath)
	if err != nil {
		return fmt.Errorf("unable to create output gif at %s: %w", outputImage, err)
	}

	defer outputImage.Close()

	// add solution
	s.drawCurrentFrameToGif()
	s.gif.Delay = make([]int, len(s.gif.Image))

	slog.Info(fmt.Sprintf("gif contains %d frames", len(s.gif.Image)))
	s.gif.Delay[len(s.gif.Delay)-1] = 500
	err = gif.EncodeAll(outputImage, s.gif)
	if err != nil {
		return fmt.Errorf("unable to encode gif: %w", err)
	}

	return nil
}
