package structs

type Cell struct {
	state State
}

func (cell *Cell) GetState() State {
	return cell.state
}

func (cell *Cell) SetState(state State) {
	cell.state = state
}
