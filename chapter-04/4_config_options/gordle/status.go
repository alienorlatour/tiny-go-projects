package gordle

import "strings"

// status describes the validity of a character in a word.
type status int

const (
	absentCharacter status = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface.
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

// feedback is a list of statuses, one per character of the word.
type feedback []status

// String implements the Stringer interface for a slice of statuses.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, s := range fb {
		sb.WriteString(s.String())
	}
	return sb.String()
}
