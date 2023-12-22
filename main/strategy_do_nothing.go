package main

import "time"

type StrategyDoNothing struct{}

// FindAction implements Strategy.
func (StrategyDoNothing) FindAction(state *State, player int, maxTime time.Time) *Action {
	return nil
}

func NewStrategyDoNothing() Strategy[State, Action] {
	return StrategyDoNothing{}
}
