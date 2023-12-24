package main

import (
	"time"
)

const MAX_DISTANCE = 10000 * 10000 * 2
const MIN_LIGHT_DISTANCE = 800 * 800
const MAX_LIGHT_DISTANCE = 2000 * 2000

type StrategyNearest struct{}

// FindAction implements Strategy.
func (StrategyNearest) FindAction(state *State, player int, maxTime time.Time) *Action {
	// find nearest unscanned fish
	minDistance := MAX_DISTANCE
	var nearestCreature *Creature
	for _, creature := range state.Creatures {
		if creature.IsScannedBy("player") {
			continue
		}
		deltaX := (state.Player.Drones[0].X - creature.Coords.X)
		deltaY := (state.Player.Drones[0].Y - creature.Coords.Y)
		distance := deltaX*deltaX + deltaY*deltaY
		if distance < minDistance {
			minDistance = distance
			nearestCreature = creature
		}
	}
	if nearestCreature == nil {
		return nil
	}
	light := minDistance > MIN_LIGHT_DISTANCE && minDistance <= MAX_LIGHT_DISTANCE
	return NewMoveAction(nearestCreature.Coords.X, nearestCreature.Coords.Y, light)
}

func NewStrategyNearest() Strategy[State, Action] {
	return StrategyNearest{}
}
