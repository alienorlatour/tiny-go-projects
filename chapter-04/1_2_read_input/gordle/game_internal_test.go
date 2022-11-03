package gordle

import (
	"bufio"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "hello",
			want:  []rune("hello"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := Game{reader: bufio.NewReader(strings.NewReader(tc.input))}

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
