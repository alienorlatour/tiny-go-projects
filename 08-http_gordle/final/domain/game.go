package domain

type Game struct {
	ID           string  `json:"id"`
	AttemptsLeft string  `json:"attemptsLeft"`
	Guesses      []Guess `json:"guesses"`
}

type Guess struct {
}
