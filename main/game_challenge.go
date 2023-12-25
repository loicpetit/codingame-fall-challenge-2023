package main

import (
	"fmt"
)

const MAP_SIZE = 10000
const MAX_DRONE_DISTANCE = 600

type ChallengeGame struct {
	coords CoordsHelper
}

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
		return game.updateGameState(state, action.Player)
	}
	if action.Type == "WAIT" {
		return game.wait(state, action.Player, action.Light)
	} else if action.Type == "MOVE" {
		return game.move(state, action.Player, action.X, action.Y, action.Light)
	} else {
		panic(fmt.Sprintf("Unhandled action type: %s", action.Type))
	}
}

func (ChallengeGame) updateGameState(state *State, player int) *State {
	if player != 3 {
		panic("Only player 3 can update game state")
	}
	return state
}

func (game ChallengeGame) wait(state *State, player int, shouldLight bool) *State {
	if player == 1 {
		newDrone := game.waitDrone(state.Player.GetFirstDrone(), shouldLight)
		return state.SetPlayer(
			NewPlayerState(
				state.Player.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else if player == 2 {
		newDrone := game.waitDrone(state.Foe.GetFirstDrone(), shouldLight)
		return state.SetFoe(
			NewPlayerState(
				state.Foe.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else {
		panic(fmt.Sprintf("Unhandled player %d", player))
	}
}

func (game ChallengeGame) move(state *State, player int, x int, y int, shouldLight bool) *State {
	if player == 1 {
		newDrone := game.moveDrone(state.Player.GetFirstDrone(), x, y, shouldLight)
		return state.SetPlayer(
			NewPlayerState(
				state.Player.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else if player == 2 {
		newDrone := game.moveDrone(state.Foe.GetFirstDrone(), x, y, shouldLight)
		return state.SetFoe(
			NewPlayerState(
				state.Foe.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else {
		panic(fmt.Sprintf("Unhandled player %d", player))
	}
}

func (game ChallengeGame) waitDrone(drone *Drone, shouldLight bool) *Drone {
	return game.moveDrone(drone, drone.X, drone.Y+300, shouldLight)
}

func (game ChallengeGame) moveDrone(drone *Drone, x int, y int, shouldLight bool) *Drone {
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
	x, y = game.coords.GetNextCoords(drone.X, drone.Y, x, y, MAX_DRONE_DISTANCE)
	return NewDrone(
		drone.Id,
		game.fixCoord(x),
		game.fixCoord(y),
		drone.Emergency,
		battery,
		light,
	)
}

// Fix coords, if < 0 set 0, if >= max set max - 1
func (ChallengeGame) fixCoord(coord int) int {
	if coord < 0 {
		return 0
	} else if coord >= MAP_SIZE {
		return MAP_SIZE - 1
	}
	return coord
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
	return ChallengeGame{
		coords: NewCoordsHelper(),
	}
}
