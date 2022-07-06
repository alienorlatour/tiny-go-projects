package gordle

import (
	_ "embed" // required to use the embedding mechanism
	"math/rand"
	"strings"
	"time"
)

//go:embed corpus/english_5letters.txt
var corpus string

// randomWord returns a random word from the corpus
func randomWord() []rune {
	list := strings.Split(corpus, "\n")

	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Int() % len(list)

	return []rune(strings.ToUpper(list[index]))
}
