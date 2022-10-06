package automaton

type automaton struct {
	cells          []cell
	states         map[string]state
	transitions    map[string]state
	leftNeighbors  int
	rightNeighbors int
}

func NewAutomaton(cellAmount int, states []string, transitions map[string]string, leftNeighbors int, rightNeighbors int) *automaton {
	return &automaton{
		make([]cell, cellAmount),
		createStates(states),
		createTransitions(transitions),
		leftNeighbors,
		rightNeighbors,
	}
}

func (automaton *automaton) CurrentState() (cells []string) {
	cells = make([]string, len(automaton.cells))
	for key, value := range automaton.cells {
		cells[key] = value.state.value
	}

	return
}

func (automaton *automaton) NextGeneration() {
	newCells := make([]cell, len(automaton.cells))
	for i := 0; i < len(automaton.cells); i++ {
		cellNeighborhood := automaton.obtainCellNeighborhood(i)
		newCells[i] = cell{state: automaton.newStateFromRule(cellNeighborhood)}
	}

	automaton.cells = newCells
}

func (automaton *automaton) SetCells(states []string) {
	for key := range automaton.cells {
		automaton.cells[key].state = automaton.states[states[key]]
	}
}

func (automaton *automaton) newStateFromRule(state []int) state {
	rule := automaton.obtainRule(state)
	return automaton.transitions[rule]
}

func (automaton *automaton) obtainRule(state []int) (rule string) {
	rule = ""
	for _, value := range state {
		rule += automaton.cells[value].state.value + ";"
	}

	return
}

//obtains the neighbor cells indexes according to a circular ending
func (automaton *automaton) obtainCellNeighborhood(cellPosition int) (answer []int) {
	rightNeighbors := getRightCells(automaton.rightNeighbors, cellPosition, len(automaton.cells))
	leftNeighbors := getLeftCells(automaton.leftNeighbors, cellPosition, len(automaton.cells))

	answer = append(leftNeighbors, cellPosition)
	answer = append(answer, rightNeighbors...)

	return
}

func getRightCells(steps int, center int, max int) (rightNeighbors []int) {
	rightNeighbors = make([]int, steps)
	current := center

	for i := 0; i < steps; i++ {
		current += 1

		if current >= max {
			current = 0
		}

		rightNeighbors[i] = current
	}

	return
}

func getLeftCells(steps int, center int, max int) (leftNeighbors []int) {
	leftNeighbors = make([]int, steps)
	current := center

	for i := 0; i < steps; i++ {
		current -= 1

		if current < 0 {
			current = max - 1
		}

		leftNeighbors[len(leftNeighbors)-1-i] = current
	}

	return
}

func createStates(states []string) (properStates map[string]state) {
	properStates = make(map[string]state, len(states))
	for _, value := range states {
		properStates[value] = state{value: value}
	}

	return
}

func createTransitions(transitions map[string]string) (properTransitions map[string]state) {
	properTransitions = make(map[string]state, len(transitions))
	for key, value := range transitions {
		properTransitions[key] = state{value: value}
	}

	return
}
