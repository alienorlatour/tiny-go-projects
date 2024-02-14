package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_neighbours(t *testing.T) {
	testCases := map[string]struct {
		p    image.Point
		want []image.Point
	}{
		"1,1": {
			p: image.Point{X: 1, Y: 1},
			want: []image.Point{
				{0, 1}, {1, 0}, {2, 1}, {1, 2},
			},
		},
		"8, -6": {
			p: image.Point{X: 8, Y: -6},
			want: []image.Point{
				{8, -5}, {8, -7}, {9, -6}, {7, -6},
			},
		},
	}
	for name, tt := range testCases {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.ElementsMatch(t, tt.want, neighbours(tt.p), "neighbours(%v)", tt.p)
		})
	}
}
