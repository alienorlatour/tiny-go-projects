package gordle

import (
	"bufio"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		reader  *bufio.Reader
		want    []rune
		wantErr bool
	}{
		"5 characters in english": {
			reader:  bufio.NewReader(strings.NewReader("hello")),
			want:    []rune("hello"),
			wantErr: false,
		},
		"5 characters in arabic": {
			reader:  bufio.NewReader(strings.NewReader("مرحبا")),
			want:    []rune("مرحبا"),
			wantErr: false,
		},
		"5 characters in japanese": {
			reader:  bufio.NewReader(strings.NewReader("のうにゅう")),
			want:    []rune("のうにゅう"),
			wantErr: false,
		},
		"3 characters in japanese": {
			reader: bufio.NewReader(strings.NewReader("のうに\nのうにゅう")),
			want:   []rune("のうにゅう"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := Game{reader: tc.reader}

			got := g.ask()
			if !compareRunes(got, tc.want) {
				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

// compareRunes compares two slices and returns whether they have the same elements.
func compareRunes(s1, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v1 := range s1 {
		if v1 != s2[i] {
			return false
		}
	}
	return true
}
