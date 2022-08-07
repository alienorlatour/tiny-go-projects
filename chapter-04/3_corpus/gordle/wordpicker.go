package gordle

import (
	"math/rand"
	"strings"
	"time"
)

// pickWord returns a random word from the corpus
func pickWord(corpus []string) []rune {
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Int() % len(corpus)

	return []rune(strings.ToUpper(corpus[index]))
}
