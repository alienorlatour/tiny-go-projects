package wordle

type status int

const (
	unknown status = iota
	correctPosition
	wrongPosition
	absentCharacter
)
