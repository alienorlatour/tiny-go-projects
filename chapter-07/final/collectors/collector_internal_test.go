package collectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountItems(t *testing.T) {
	tt := map[string]struct {
		input Collectors[item]
		want  map[item]uint
	}{
		"nominal use case": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "handmaidsTale", "janeEyre"}},
			},
			want: map[item]uint{"handmaidsTale": 2, "theBellJar": 1, "oryxAndCrake": 1, "janeEyre": 1},
		},
		"no colls": {
			input: Collectors[item]{},
			want:  map[item]uint{},
		},
		"coll without books": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{}},
			},
			want: map[item]uint{"handmaidsTale": 1, "theBellJar": 1},
		},
		"coll with twice the same book": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar", "handmaidsTale"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "handmaidsTale", "janeEyre"}},
			},
			want: map[item]uint{"handmaidsTale": 3, "theBellJar": 1, "oryxAndCrake": 1, "janeEyre": 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.countItems()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFindCommon(t *testing.T) {
	tt := map[string]struct {
		input Collectors[item]
		want  []item
	}{
		"no common book": {
			input: Collectors[item]{
				{Name: "Fadi", Items: []item{"handmaidsTale", "theBellJar"}},
				{Name: "Peggy", Items: []item{"oryxAndCrake", "janeEyre"}},
			},
			want: nil,
		},
		"one common book": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"oryxAndCrake", "janeEyre"}},
				{Name: "Did", Items: []item{"janeEyre"}},
			},
			want: []item{"janeEyre"},
		},
		"three colls have the same books on their shelves": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"oryxAndCrake", "ilPrincipe", "janeEyre"}},
				{Name: "Did", Items: []item{"janeEyre"}},
				{Name: "Ali", Items: []item{"janeEyre", "ilPrincipe"}},
			},
			want: []item{"janeEyre", "ilPrincipe"},
		},
		"output is sorted by authors and then title": {
			input: Collectors[item]{
				{Name: "Peggy", Items: []item{"ilPrincipe", "janeEyre", "villette"}},
				{Name: "Did", Items: []item{"janeEyre"}},
				{Name: "Ali", Items: []item{"villette", "ilPrincipe"}},
			},
			want: []item{"janeEyre", "villette", "ilPrincipe"},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.FindCommon()
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
