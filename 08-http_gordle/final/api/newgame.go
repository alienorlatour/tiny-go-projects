package api

type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	Solution     string  `json:"solution,omitempty"`
}

type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
