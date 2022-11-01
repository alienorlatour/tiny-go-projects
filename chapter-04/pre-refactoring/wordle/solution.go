package wordle

// Solution holds the positions of the valid characters
// since a single character can appear several times, we store these times as a slice of indexes
type Solution struct {
	word      []rune
	positions map[rune][]int
}

// NewSolution builds the solution to a game.
func NewSolution(word []rune) Solution {
	sol := Solution{
		word:      word,
		positions: make(map[rune][]int),
	}

	return sol
}

// Feedback prints out hints on how to find the solution.
func (s *Solution) Feedback(attempt []rune) []status {
	// reset the positions map
	s.positions = make(map[rune][]int)

	for i, character := range s.word {
		// appending to a nil slice will return a slice, this is safe
		s.positions[character] = append(s.positions[character], i)
	}

	feedback := make([]status, len(s.word))

	// scan the attempts and check if they are in the solution
	for i, character := range attempt {
		// keep track of already seen characters
		correctness := s.checkCharacterAtPosition(character, i)
		if correctness == correctPosition {
			// remove found character from positions
			s.markCharacterAsSeen(character, i)
			feedback[i] = correctPosition
		}
	}

	for i, character := range attempt {
		if feedback[i] == correctPosition {
			continue
		}

		correctness := s.checkCharacterAtPosition(character, i)

		if correctness == wrongPosition {
			// remove the left-most occurrence
			s.positions[character] = s.positions[character][1:]
		}

		feedback[i] = correctness
	}

	return feedback
}

// markCharacterAsSeen removes one occurrence of the character from the positions map.
func (s *Solution) markCharacterAsSeen(character rune, positionInWord int) {
	positions := s.positions[character]

	if len(positions) == 0 {
		s.positions[character] = nil
	}

	for i, pos := range positions {
		if pos == positionInWord {
			// remove the seen character from the list
			s.positions[character] = append(positions[:i], positions[i+1:]...)
			// we found it
			return
		}
	}
}

// checkCharacterAtPosition returns the correctness of a character
// at the specified index in the solution.
func (s *Solution) checkCharacterAtPosition(character rune, index int) status {
	positions, ok := s.positions[character]
	if !ok || len(positions) == 0 {
		return absentCharacter
	}

	for _, pos := range positions {
		if pos == index {
			return correctPosition
		}
	}

	return wrongPosition
}

// IsWord returns whether the attempt is the solution.
func (s *Solution) IsWord(attempt []rune) bool {
	return string(s.word) == string(attempt)
}
