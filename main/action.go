package main

import "fmt"

type Action struct {
	Type  string
	X     int
	Y     int
	Light bool
}

func (action *Action) String() string {
	if action == nil {
		return ""
	}
	return fmt.Sprintf("{type:%s, x:%d, y:%d, light:%t}", action.Type, action.X, action.Y, action.Light)
}

func NewMoveAction(x, y int, light bool) *Action {
	return &Action{
		Type:  "MOVE",
		X:     x,
		Y:     y,
		Light: light,
	}
}

func NewWaitAction(light bool) *Action {
	return &Action{
		Type:  "WAIT",
		X:     0,
		Y:     0,
		Light: light,
	}
}
