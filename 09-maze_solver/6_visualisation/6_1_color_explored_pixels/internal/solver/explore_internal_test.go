package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver_explore(t *testing.T) {
	tests := map[string]struct {
		inputImage string
		wantSize   int
	}{
		"cross": {
			inputImage: "testdata/explore_cross.png",
			wantSize:   2,
		},
		"dead end": {
			inputImage: "testdata/explore_deadend.png",
			wantSize:   0,
		},
		"double": {
			inputImage: "testdata/explore_double.png",
			wantSize:   1,
		},
		"treasure": {
			inputImage: "testdata/explore_treasure.png",
			wantSize:   0,
		},
		"treasure only": {
			inputImage: "testdata/explore_treasureonly.png",
			wantSize:   0,
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			maze, err := openMaze(tt.inputImage)
			require.NoError(t, err)

			s := &Solver{
				maze:           maze,
				palette:        defaultPalette(),
				pathsToExplore: make(chan *path, 3),
				quit:           make(chan struct{}),
			}

			// All our tests have the entrance at the same position.
			s.explore(&path{previousStep: nil, at: image.Point{X: 0, Y: 2}})

			assert.Equal(t, tt.wantSize, len(s.pathsToExplore))
		})
	}
}
