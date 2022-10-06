package main

import (
	"fmt"

	"github.com/rojbar/acu/structs"
)

func main() {
	states := make([]string, 2)
	states[0] = "⬜"
	states[1] = "⬛"

	transitions := make(map[string]string)
	transitions["⬛;⬛;⬛;"] = "⬜"
	transitions["⬛;⬛;⬜;"] = "⬜"
	transitions["⬛;⬜;⬛;"] = "⬜"
	transitions["⬛;⬜;⬜;"] = "⬜"
	transitions["⬜;⬛;⬛;"] = "⬛"
	transitions["⬜;⬛;⬜;"] = "⬛"
	transitions["⬜;⬜;⬛;"] = "⬜"
	transitions["⬜;⬜;⬜;"] = "⬛"

	auto := structs.NewAutomaton(50, states, transitions, 1, 1)
	auto.SetCells(initial(50))

	fmt.Println(auto.CurrentState())
	for i := 0; i < 50; i++ {
		auto.NextGeneration()
		fmt.Println(auto.CurrentState())
	}

}

func initial(size int) []string {
	a := make([]string, size)

	med := size / 2
	for i := 0; i < size; i++ {

		if i == int(med) {
			a[i] = "⬛"
		} else {
			a[i] = "⬜"
		}

	}

	return a
}
