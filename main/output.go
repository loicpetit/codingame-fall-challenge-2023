package main

import "fmt"

type Writer struct{}

func (Writer) Write(action *Action) {
	if action == nil {
		panic("No action to write")
	}
	if action.Player != 1 {
		panic("Only manage player 1 actions")
	}
	light := 0
	if action.Light {
		light = 1
	}
	if action.Type == "MOVE" {
		WriteOutput(fmt.Sprintf("MOVE %d %d %d", action.X, action.Y, light))
	} else if action.Type == "WAIT" {
		WriteOutput(fmt.Sprintf("WAIT %d", light))
	} else {
		panic(fmt.Sprintf("Unknown action type: %s", action.Type))
	}
}

func NewWriter() OutputWriter[Action] {
	return Writer{}
}
