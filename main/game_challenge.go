package main

import (
	"fmt"
)

const MAP_SIZE = 10000

type ChallengeGame struct{}

// GetAvailableActions implements Game.
func (ChallengeGame) GetAvailableActions(state *State, player int) []*Action {
	// player 1 = player
	// player 2 = foe
	// player 3 = game update
	if player == 3 {
		return []*Action{NewGameUpdateAction()}
	}
	return nil
}

// GetLastPlayer implements Game.
func (ChallengeGame) GetLastPlayer(state *State) int {
	// player 1 = player
	// player 2 = foe
	// player 3 = game update
	if state == nil {
		panic("State is required")
	}
	return state.LastPlayer
}

// GetNextPlayer implements Game.
func (ChallengeGame) GetNextPlayer(state *State) int {
	// player 1 = player
	// player 2 = foe
	// player 3 = game update
	if state == nil {
		panic("State is required")
	}
	if state.LastPlayer == 2 {
		return 3
	}
	if state.LastPlayer == 1 {
		return 2
	}
	return 1
}

// Play implements Game.
func (game ChallengeGame) Play(state *State, action *Action) *State {
	if state == nil || action == nil {
		return state
	}
	if action.Type == "GAME_UPDATE" {
		return game.updateGameState(state)
	}
	if action.Type == "WAIT" {
		return game.wait(state, action)
	} else if action.Type == "MOVE" {
		return game.move(state, action)
	} else {
		panic(fmt.Sprintf("Unhandled action type: %s", action.Type))
	}
}

func (ChallengeGame) updateGameState(state *State) *State {
	return state
}

func (game ChallengeGame) wait(state *State, action *Action) *State {
	if action.Player == 1 {
		drone := state.Player.GetFirstDrone()
		newY := drone.Y + 300
		if newY >= MAP_SIZE {
			newY = MAP_SIZE - 1
		}
		newDrone := game.moveDrone(drone, drone.X, newY, action.Light)
		return state.SetPlayer(
			NewPlayerState(
				state.Player.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else if action.Player == 2 {
		drone := state.Foe.GetFirstDrone()
		newY := drone.Y + 300
		if newY >= MAP_SIZE {
			newY = MAP_SIZE - 1
		}
		newDrone := game.moveDrone(drone, drone.X, newY, action.Light)
		return state.SetFoe(
			NewPlayerState(
				state.Foe.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else {
		panic(fmt.Sprintf("Unhandled player %d", action.Player))
	}
}

func (ChallengeGame) move(state *State, action *Action) *State {
	return state
}

func (ChallengeGame) moveDrone(drone *Drone, x int, y int, shouldLight bool) *Drone {
	if drone == nil {
		return nil
	}
	battery := drone.Battery
	light := false
	if shouldLight && battery >= 5 {
		battery = battery - 5
		light = true
	} else if battery < 30 {
		battery = battery + 1
	}
	return NewDrone(
		drone.Id,
		x,
		y,
		drone.Emergency,
		battery,
		light,
	)
}

// Start implements Game.
func (ChallengeGame) Start() *State {
	return NewState()
}

// Winner implements Game.
func (ChallengeGame) Winner(state *State) int {
	return 0
}

func NewChallengeGame() Game[State, Action] {
	return ChallengeGame{}
}
