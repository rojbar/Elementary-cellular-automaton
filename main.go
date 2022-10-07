package main

import (
	"syscall/js"

	"github.com/rojbar/acu/automaton"
)

var (
	auto       *automaton.Automaton
	generation int
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("newAutomata", js.FuncOf(newAutomata))
	js.Global().Set("getCurrentState", js.FuncOf(getCurrentState))
	js.Global().Set("getNextGeneration", js.FuncOf(getNextGeneration))
	js.Global().Set("setInitialState", js.FuncOf(setInitialState))
	js.Global().Set("getCurrentGeneration", js.FuncOf(getCurrentGeneration))
	<-c
}

//args:
// 0 -> cell amount
// 1 -> left neighbors
// 2 -> right neighbors
// 3 -> states string[]
// 4 -> keys string[]
// 5 -> values string[]
func newAutomata(this js.Value, args []js.Value) interface{} {
	generation = 0
	cellAmount := args[0].Int()
	leftNeighbors := args[1].Int()
	rightNeighbors := args[2].Int()

	states := make([]string, args[3].Length())
	states = initSliceString(args[3], args[3].Length())

	keys := make([]string, args[4].Length())
	keys = initSliceString(args[4], args[4].Length())

	values := make([]string, args[5].Length())
	values = initSliceString(args[5], args[5].Length())

	transitions := make(map[string]string, args[4].Length())

	for i := 0; i < args[4].Length(); i++ {
		transitions[keys[i]] = values[i]
	}

	auto = automaton.NewAutomaton(cellAmount, states, transitions, leftNeighbors, rightNeighbors)

	return true
}

func getCurrentState(this js.Value, args []js.Value) interface{} {
	state := auto.CurrentState()
	aux := make([]interface{}, len(state))
	for key, value := range state {
		aux[key] = value
	}

	return aux
}

func getCurrentGeneration(this js.Value, args []js.Value) interface{} {
	return generation
}

func getNextGeneration(this js.Value, args []js.Value) interface{} {
	auto.NextGeneration()
	generation += 1

	return true
}

// 0 -> []string
func setInitialState(this js.Value, args []js.Value) interface{} {
	values := make([]string, args[0].Length())
	values = initSliceString(args[0], args[0].Length())
	auto.SetCells(values)

	return true
}

func initSliceString(args js.Value, length int) (v []string) {
	v = make([]string, length)
	for i := 0; i < length; i++ {
		v[i] = args.Index(i).String()
	}

	return
}
