package gordle

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

	for i, letter := range s.word {
		// appending to a nil slice will return a slice, this is safe
		s.positions[letter] = append(s.positions[letter], i)
	}

	feedback := make([]status, len(s.word))

	// scan the attempts and check if they are in the solution
	for i, letter := range attempt {
		// keep track of already seen characters
		correctness := s.checkLetterAtPosition(letter, i)
		if correctness == correctPosition {
			// remove found letter from positions
			s.markLetterAsSeen(letter, i)
			feedback[i] = correctPosition
		}
	}

	for i, letter := range attempt {
		if feedback[i] == correctPosition {
			continue
		}

		correctness := s.checkLetterAtPosition(letter, i)

		if correctness == wrongPosition {
			// remove the left-most occurrence
			s.positions[letter] = s.positions[letter][1:]
		}

		feedback[i] = correctness
	}

	return feedback
}

// markLetterAsSeen removes one occurrence of the letter from the positions map.
func (s *Solution) markLetterAsSeen(letter rune, positionInWord int) {
	positions := s.positions[letter]

	if len(positions) == 0 {
		s.positions[letter] = nil
	}

	for i, pos := range positions {
		if pos == positionInWord {
			// remove the seen letter from the list
			s.positions[letter] = append(positions[:i], positions[i+1:]...)
			// we found it
			return
		}
	}
}

// checkLetterAtPosition returns the correctness of a letter
// at the specified index in the solution.
func (s *Solution) checkLetterAtPosition(letter rune, index int) status {
	positions, ok := s.positions[letter]
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
