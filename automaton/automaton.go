package automaton

import (
	"fmt"
)

var ruleSeparator = ";"

type Automaton struct {
	cells          []cell
	states         map[string]state
	transitions    map[string]state
	leftNeighbors  int
	rightNeighbors int
}

func NewAutomaton(cellAmount int, states []string, transitions map[string]string, leftNeighbors int, rightNeighbors int) *Automaton {
	return &Automaton{
		make([]cell, cellAmount),
		createStates(states),
		createTransitions(transitions),
		leftNeighbors,
		rightNeighbors,
	}
}

func (automaton *Automaton) GetCurrentState() []string {
	cells := make([]string, len(automaton.cells))
	for key, currentCell := range automaton.cells {
		cells[key] = currentCell.state.value
	}

	return cells
}

func (automaton *Automaton) CalculateNextGeneration() {
	newCells := make([]cell, len(automaton.cells))
	for i := 0; i < len(automaton.cells); i++ {
		cellNeighbors := automaton.obtainCellNeighbors(i)
		newCells[i] = cell{state: automaton.newStateFromCellNeighbors(cellNeighbors)}
	}

	automaton.cells = newCells
}

func (automaton *Automaton) SetEachCellNewState(newCellStates []string) error {
	if len(newCellStates) != len(automaton.cells) {
		return fmt.Errorf("not equal amount of new cell states for the quantity of cells, expected %d given %d", len(automaton.cells), len(newCellStates))
	}

	newCells := make([]cell, len(automaton.cells))

	for key := range automaton.cells {
		newCellState, exists := automaton.states[newCellStates[key]]
		if !exists {
			return fmt.Errorf("passed state %s at position %d that doesn't exist in the current automaton", newCellStates[key], key)
		}

		newCells[key].state = newCellState
	}

	automaton.cells = newCells

	return nil
}

// obtains the neighbor cells according to a circular ending
func (automaton *Automaton) obtainCellNeighbors(cellPosition int) []cell {
	rightNeighbors := getRightCellsPositions(automaton.rightNeighbors, cellPosition, len(automaton.cells))
	leftNeighbors := getLeftCellsPositions(automaton.leftNeighbors, cellPosition, len(automaton.cells))

	neighborsPosition := append(leftNeighbors, cellPosition)
	neighborsPosition = append(neighborsPosition, rightNeighbors...)

	cellNeighbors := make([]cell, len(neighborsPosition))

	for key, pos := range neighborsPosition {
		cellNeighbors[key] = automaton.cells[pos]
	}

	return cellNeighbors
}

// each rule that defines a transition is a set of states separated by the ruleSeparator
func (automaton *Automaton) newStateFromCellNeighbors(cellNeighbors []cell) state {
	rule := ""
	for _, cell := range cellNeighbors {
		rule += cell.state.value + ruleSeparator
	}

	return automaton.transitions[rule]
}
