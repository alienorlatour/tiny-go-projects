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

func StatusesToString(statuses []status) string {
	var s string
	for _, st := range statuses {
		s += st.String() + " "
	}
	return s
}
