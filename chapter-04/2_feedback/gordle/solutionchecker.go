package gordle

// solutionChecker holds all the information we need to evaluate the solution.
type solutionChecker struct {
	// the solution word
	solution []rune
	// keep track of the positions of the runes in the solution word
	positions map[rune][]int
}

// evaluate verifies every letter of the word against the solution.
func (sc *solutionChecker) evaluate(word []rune) feedback {
	// reset the positions map
	sc.reset()

	fb := make(feedback, len(sc.solution))

	// scan the attempts and evaluate if they are in the solution
	for i, letter := range word {
		correctness := sc.checkLetterAtPosition(letter, i)
		if correctness == correctPosition {
			// keep track of already seen characters
			sc.markLetterAsSeen(letter, i)
			fb[i] = correctPosition
		}
	}

	for i, letter := range word {
		if fb[i] == correctPosition {
			continue
		}

		correctness := sc.checkLetterAtPosition(letter, i)
		if correctness == wrongPosition {
			// remove the left-most occurrence
			sc.positions[letter] = sc.positions[letter][1:]
			fb[i] = wrongPosition
		}
	}

	// letters not found in the word have the zero value absentCharacter

	return fb
}

// reset rebuilds the initial map holding the letters and their positions.
func (sc *solutionChecker) reset() {
	sc.positions = make(map[rune][]int)
	for i, letter := range sc.solution {
		// appending to a nil slice will return a slice, this is safe
		sc.positions[letter] = append(sc.positions[letter], i)
	}
}

// checkLetterAtPosition returns the correctness of a letter at the specified index in the solution.
func (sc *solutionChecker) checkLetterAtPosition(letter rune, index int) status {
	positions, ok := sc.positions[letter]
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

// markLetterAsSeen removes one occurrence of the letter from the positions map.
func (sc *solutionChecker) markLetterAsSeen(letter rune, positionInWord int) {
	positions := sc.positions[letter]

	for i, pos := range positions {
		if pos == positionInWord {
			// remove the seen letter from the list
			sc.positions[letter] = append(positions[:i], positions[i+1:]...)
			// we found it
			return
		}
	}
}
