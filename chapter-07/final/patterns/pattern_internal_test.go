package patterns_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/genericworms/collectors"
	"learngo-pockets/genericworms/patterns"
)

func TestUnmarshalPattern(t *testing.T) {
	encoding := []byte(`{
        "craft": "crochet",
        "name": "Rainbow in my pocket",
        "yardage": 2370
      }`)

	var pattern patterns.Pattern

	err := json.Unmarshal(encoding, &pattern)

	assert.NoError(t, err)
	assert.Equal(t, patterns.Pattern{
		Craft:   "crochet",
		Name:    "Rainbow in my pocket",
		Yardage: 2370,
	}, pattern)
}

func TestPatternBefore(t *testing.T) {
	testCases := map[string]struct {
		left  patterns.Pattern
		right collectors.Sortable
		want  bool
	}{
		"Patterns of different craft": {
			left:  patterns.Pattern{Craft: "knit", Name: "Sparky", Yardage: 2620},
			right: patterns.Pattern{Craft: "crochet", Name: "Rainbow in my pocket", Yardage: 2370},
			want:  false,
		},
		"Patterns of different craft, reversed": {
			left:  patterns.Pattern{Craft: "crochet", Name: "Rainbow in my pocket", Yardage: 2370},
			right: patterns.Pattern{Craft: "knit", Name: "Sparky", Yardage: 2620},
			want:  true,
		},
		"Patterns of same craft, but different name": {
			left:  patterns.Pattern{Craft: "knit", Name: "Sparky", Yardage: 2620},
			right: patterns.Pattern{Craft: "knit", Name: "Plaid Pocket Socks", Yardage: 450},
			want:  false,
		},
		"Patterns of same craft, same name, but different yardage": {
			left:  patterns.Pattern{Craft: "knit", Name: "Hat", Yardage: 280},
			right: patterns.Pattern{Craft: "knit", Name: "Hat", Yardage: 300},
			want:  true,
		},
		"Trying to compare a pattern with a non-pattern": {
			left:  patterns.Pattern{Craft: "sewing", Name: "Rosalie Skirt", Yardage: 3},
			right: nonPattern{},
			want:  false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.left.Before(testCase.right)
			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestPatternString(t *testing.T) {
	p := patterns.Pattern{
		Craft:   "knit",
		Name:    "Plaid Pocket Socks",
		Yardage: 450,
	}

	want := `Craft: knit      ; Name: Plaid Pocket Socks  ; Yardage:   450 yards`
	assert.Equal(t, want, p.String())
}

type nonPattern struct {
}

func (n nonPattern) Before(_ collectors.Sortable) bool {
	panic("This is a test utility, it shouldn't be called")
}
