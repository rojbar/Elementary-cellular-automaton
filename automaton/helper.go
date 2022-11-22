package automaton

func createStates(states []string) map[string]state {
	properStates := make(map[string]state, len(states))
	for _, value := range states {
		properStates[value] = state{value: value}
	}

	return properStates
}

func createTransitions(transitions map[string]string) map[string]state {
	properTransitions := make(map[string]state, len(transitions))
	for key, value := range transitions {
		properTransitions[key] = state{value: value}
	}

	return properTransitions
}

func getRightCellsPositions(steps int, center int, max int) []int {
	rightNeighbors := make([]int, steps)
	current := center

	for i := 0; i < steps; i++ {
		current += 1

		if current >= max {
			current = 0
		}

		rightNeighbors[i] = current
	}

	return rightNeighbors
}

func getLeftCellsPositions(steps int, center int, max int) []int {
	leftNeighbors := make([]int, steps)
	current := center

	for i := 0; i < steps; i++ {
		current -= 1

		if current < 0 {
			current = max - 1
		}

		leftNeighbors[len(leftNeighbors)-1-i] = current
	}

	return leftNeighbors
}
