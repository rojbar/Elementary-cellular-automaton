// to export our functions we must wrap our functions to be exported in main using syscall/js
package main

import (
	"syscall/js"

	"github.com/rojbar/acu/automaton"
)

var (
	currentAutomaton  *automaton.Automaton
	currentGeneration int
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("newAutomaton", js.FuncOf(newAutomaton))
	js.Global().Set("getCurrentState", js.FuncOf(getCurrentState))
	js.Global().Set("calculateNextGeneration", js.FuncOf(calculateNextGeneration))
	js.Global().Set("setInitialState", js.FuncOf(setInitialState))
	js.Global().Set("getCurrentGeneration", js.FuncOf(getCurrentGeneration))
	<-c
}

/*
	args:

	pos  |name          |type
	-------------------------
	0     cellAmount	 int
	1     leftNeighbors  int
	2     rightNeighbors int
	3     states         string[]
	4     keys           string[]
	5     values         string[]
*/
func newAutomaton(this js.Value, args []js.Value) interface{} {
	currentGeneration = 0
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

	currentAutomaton = automaton.NewAutomaton(cellAmount, states, transitions, leftNeighbors, rightNeighbors)

	return true
}

func getCurrentState(this js.Value, args []js.Value) interface{} {
	state := currentAutomaton.GetCurrentState()
	aux := make([]interface{}, len(state))
	for key, value := range state {
		aux[key] = value
	}

	return aux
}

func calculateNextGeneration(this js.Value, args []js.Value) interface{} {
	currentAutomaton.CalculateNextGeneration()
	currentGeneration += 1

	return true
}

/*
	args:

	pos  |name          |type
	-------------------------
	0     cellStates	 string[]
*/
func setInitialState(this js.Value, args []js.Value) interface{} {
	cellStates := make([]string, args[0].Length())
	cellStates = initSliceString(args[0], args[0].Length())
	currentAutomaton.SetEachCellNewState(cellStates)

	return true
}

func getCurrentGeneration(this js.Value, args []js.Value) interface{} {
	return currentGeneration
}

func initSliceString(args js.Value, length int) []string {
	v := make([]string, length)
	for i := 0; i < length; i++ {
		v[i] = args.Index(i).String()
	}

	return v
}
