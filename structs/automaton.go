package structs

type automaton struct {
	cells          []Cell
	states         map[string]State
	transitions    map[string]State
	leftNeighbors  int
	rightNeighbors int
	currentCell    int
}

func NewAutomaton(cellAmount int, states map[string]State, transitions map[string]State, leftNeighbors int, rightNeighbors int) *automaton {
	return &automaton{
		make([]Cell, cellAmount),
		states,
		transitions,
		leftNeighbors,
		rightNeighbors,
		0,
	}
}

func (automaton *automaton) CurrentState() (cells []Cell) {
	cells = make([]Cell, len(automaton.cells))
	copy(cells, automaton.cells)
	return
}

func (automaton *automaton) NextIteration() {
	cellNeighborhood := automaton.obtainCellNeighborhood(automaton.currentCell)
	automaton.updateCell(automaton.currentCell, cellNeighborhood)

	automaton.currentCell += 1

	if automaton.currentCell == len(automaton.cells) {
		automaton.currentCell = 0
	}
}

func (automaton *automaton) updateCell(currentCell int, state []int) {
	rule := automaton.obtainRule(state)
	newState := automaton.transitions[rule]
	automaton.cells[currentCell].state = newState
}

func (automaton *automaton) obtainRule(state []int) (rule string) {
	rule = ""
	for _, value := range state {
		rule += automaton.cells[value].state.value + ";"
	}

	return
}

func (automaton *automaton) obtainCellNeighborhood(cellPosition int) (answer []int) {
	rightNeighbors := getRightCells(automaton.rightNeighbors, cellPosition, len(automaton.cells))
	leftNeighbors := getLeftCells(automaton.leftNeighbors, cellPosition, len(automaton.cells))

	answer = append(leftNeighbors, cellPosition)
	answer = append(answer, rightNeighbors...)

	return
}

func getRightCells(steps int, center int, max int) (rightNeighbors []int) {
	current := center

	for i := 0; i < steps; i++ {
		current += 1

		if current == max {
			current = 0
		}

		rightNeighbors[i] = current
	}

	return
}

func getLeftCells(steps int, center int, max int) (leftNeighbors []int) {
	current := center

	for i := 0; i < steps; i++ {
		current -= 1

		if current == 0 {
			current = max - 1
		}

		leftNeighbors[max-1-i] = current
	}

	return
}
