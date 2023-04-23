package domain

type Game struct {
	ID           string
	AttemptsLeft byte
	Guesses      []Guess
	Solution     string
}

type Guess struct {
	Word     string
	Feedback string
}
