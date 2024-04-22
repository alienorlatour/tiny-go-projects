//go:build !exercise

package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	const layout = "02 January 2006"

	tests := map[string]struct {
		include time.Time
		want    FormattedWeek
	}{
		"Mon. 01 Jan 2024": {
			include: time.Date(2024, time.January, 1, 12, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, time.January, 6, 23, 59, 0, 0, time.UTC),
			},
		},
		"Wed. 14 Feb. 2024": {
			include: time.Date(2024, time.February, 14, 15, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2024, time.February, 11, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, time.February, 17, 23, 59, 0, 0, time.UTC),
			},
		},
		"Sun. 01 Jan 2023": {
			include: time.Date(2023, time.January, 1, 12, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, time.January, 7, 23, 59, 0, 0, time.UTC),
			},
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewWeek(tt.include, layout)
			assert.Equal(t, tt.want.start.String(), got.start.String())
			assert.Equal(t, tt.want.end.String(), got.end.String())
		})
	}
}
