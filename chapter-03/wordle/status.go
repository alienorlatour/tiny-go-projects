package wordle

type status int

const (
	absentCharacter status = iota
	wrongPosition
	correctPosition
)

func (s status) String() string {
	switch s {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		// This should never happen.
		return "💔"
	}
}
