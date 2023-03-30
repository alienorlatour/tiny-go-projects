package patterns_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

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
	// TODO
}

func TestPatternString(t *testing.T) {
	// TODO
}
