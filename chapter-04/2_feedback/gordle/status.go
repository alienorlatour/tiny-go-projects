package gordle

import "strings"

// status describes the validity of a letter in a word.
type status int

const (
	absentCharacter status = iota
	wrongPosition
	correctPosition
)

// toString returns a string that represents the status.
// We didn't implement Stringer here, as we don't plan on calling fmt.Print on
func (s status) toString() string {
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

// feedback is a list of status, one per letter of the word
type feedback []status

// String implements the Stringer interface for a slice of status.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for i, s := range fb {
		if i != 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(s.toString())
	}
	return sb.String()
}
