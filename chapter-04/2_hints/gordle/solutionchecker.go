package gordle

// solutionChecker holds all the information we need to check the solution.
type solutionChecker struct {
	// the solution word
	solution []rune
	// keep track of the positions of the runes in the solution word
	positions map[rune][]int
}

// check verifies every character of the word against the solution.
func (sc *solutionChecker) check(word []rune) feedback {
	// reset the positions map
	sc.reset()

	fb := make(feedback, len(sc.solution))

	// scan the guesses and evaluate if they are in the solution
	for i, character := range word {
		// keep track of already seen characters
		correctness := sc.checkCharacterAtPosition(character, i)
		if correctness == correctPosition {
			// remove found character from positions at this index
			sc.markCharacterAsSeen(character, i)
			fb[i] = correctPosition
		}
	}

	for i, character := range word {
		if fb[i] == correctPosition {
			continue
		}

		correctness := sc.checkCharacterAtPosition(character, i)
		if correctness == wrongPosition {
			// remove the left-most occurrence
			sc.positions[character] = sc.positions[character][1:]
			fb[i] = wrongPosition
		}
	}

	// characters not found in the word have the zero value absentCharacter

	return fb
}

// reset rebuilds the initial map holding the characters and their positions.
func (sc *solutionChecker) reset() {
	sc.positions = make(map[rune][]int)
	for i, character := range sc.solution {
		// appending to a nil slice will return a slice, this is safe
		sc.positions[character] = append(sc.positions[character], i)
	}
}

// checkCharacterAtPosition returns the correctness of a character at the specified index in the solution.
func (sc *solutionChecker) checkCharacterAtPosition(character rune, index int) hint {
	positions, ok := sc.positions[character]
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

// markCharacterAsSeen removes one occurrence of the character from the positions map.
func (sc *solutionChecker) markCharacterAsSeen(character rune, positionInWord int) {
	positions := sc.positions[character]

	for i, pos := range positions {
		if pos == positionInWord {
			// remove the seen character from the list
			sc.positions[character] = append(positions[:i], positions[i+1:]...)
			// we found it
			return
		}
	}
}
