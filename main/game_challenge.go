package main

import (
	"fmt"
)

const MAP_SIZE = 10000
const MAX_DRONE_DISTANCE = 600
const SCAN_SQUARED_DISTANCE = 800 * 800
const SCAN_WITH_LIGHT_SQUARED_DISTANCE = 2000 * 2000

type ChallengeGame struct {
	coords CoordsHelper
}

// GetAvailableActions implements Game.
func (ChallengeGame) GetAvailableActions(state *State, player int) []*Action {
	// todo
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
	expectedPlayer := game.GetNextPlayer(state)
	if action.Player != expectedPlayer {
		panic(fmt.Sprintf("Expected player to play is %d but get %d", expectedPlayer, action.Player))
	}
	lastPlayerState := state.SetLastPlayer(action.Player)
	switch action.Type {
	case "GAME_UPDATE":
		return game.updateGameState(lastPlayerState)
	case "WAIT":
		return game.wait(lastPlayerState, action.Light)
	case "MOVE":
		return game.move(lastPlayerState, action.X, action.Y, action.Light)
	default:
		panic(fmt.Sprintf("Unhandled action type: %s", action.Type))
	}
}

func (game ChallengeGame) updateGameState(state *State) *State {
	if state == nil {
		return nil
	}
	if state.LastPlayer != 3 {
		panic("Only player 3 can update game state")
	}
	updatedCreatures := make(map[int]*Creature)
	for _, creature := range state.Creatures {
		updatedCreature := game.moveCreature(creature)
		for _, drone := range state.Player.Drones {
			if game.isScannable(drone, creature) {
				updatedCreature = updatedCreature.SetScannedBy("player")
			}
		}
		for _, drone := range state.Foe.Drones {
			if game.isScannable(drone, creature) {
				updatedCreature = updatedCreature.SetScannedBy("foe")
			}
		}
		updatedCreatures[updatedCreature.Id] = updatedCreature
	}
	// TODO : update player/foe scores
	return state.SetCreatures(state.NbCreatures, updatedCreatures)
}

func (game ChallengeGame) moveCreature(creature *Creature) *Creature {
	if creature == nil {
		return nil
	}
	// TODO check collisions next turn, if collided fix Vx/Vy
	// IDEA check only previously moved creature to already have new coords
	// Update Vx/Vy of both if collision
	return creature.SetCoords(
		NewCreatureCoords(
			game.fixCoord(creature.Coords.X+creature.Coords.Vx),
			game.fixCoord(creature.Coords.Y+creature.Coords.Vy),
			creature.Coords.Vx,
			creature.Coords.Vy,
		),
	)
}

func (game ChallengeGame) isScannable(drone *Drone, creature *Creature) bool {
	if drone == nil || creature == nil {
		return false
	}
	scanSquaredDistance := SCAN_SQUARED_DISTANCE
	if drone.Light {
		scanSquaredDistance = SCAN_WITH_LIGHT_SQUARED_DISTANCE
	}
	return game.coords.GetSquaredDistance(drone.X, drone.Y, creature.Coords.X, creature.Coords.Y) <= scanSquaredDistance
}

func (game ChallengeGame) wait(state *State, shouldLight bool) *State {
	if state.LastPlayer == 1 {
		newDrone := game.waitDrone(state.Player.GetFirstDrone(), shouldLight)
		return state.SetPlayer(
			NewPlayerState(
				state.Player.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else if state.LastPlayer == 2 {
		newDrone := game.waitDrone(state.Foe.GetFirstDrone(), shouldLight)
		return state.SetFoe(
			NewPlayerState(
				state.Foe.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else {
		panic(fmt.Sprintf("Unhandled player %d", state.LastPlayer))
	}
}

func (game ChallengeGame) move(state *State, x int, y int, shouldLight bool) *State {
	if state.LastPlayer == 1 {
		newDrone := game.moveDrone(state.Player.GetFirstDrone(), x, y, shouldLight)
		return state.SetPlayer(
			NewPlayerState(
				state.Player.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else if state.LastPlayer == 2 {
		newDrone := game.moveDrone(state.Foe.GetFirstDrone(), x, y, shouldLight)
		return state.SetFoe(
			NewPlayerState(
				state.Foe.Score,
				map[int]*Drone{newDrone.Id: newDrone},
			),
		)
	} else {
		panic(fmt.Sprintf("Unhandled player %d", state.LastPlayer))
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

// Get current score of the player
func (ChallengeGame) GetScore(state *State, player int) int {
	return 0
}

// Get delta between scores of the 2 players from player perspective
// Ex: player1 has 50 player 2 has 30, delta for player 1 is 20, but for player 2 is -20
func (ChallengeGame) GetDeltaScore(state *State, player int) int {
	return 0
}

// Start implements Game.
func (ChallengeGame) GetMaxScore() int {
	return 0
}

// Winner implements Game.
func (game ChallengeGame) Winner(state *State) int {
	if state.Round == 200 || game.isEveryCreatureScanned(state) {
		if state.Player.Score > state.Foe.Score {
			return 1
		}
		if state.Foe.Score > state.Player.Score {
			return 2
		}
		return 0
	}
	if state.Player.Score > state.Foe.Score {
		// todo
		// if max possible score of foe < player score
		// return 1
	}
	if state.Foe.Score > state.Player.Score {
		// todo
		// if max possible score of player < foe score
		// return 2
	}
	return 0
}

func (ChallengeGame) isEveryCreatureScanned(state *State) bool {
	if state == nil {
		return false
	}
	for _, creature := range state.Creatures {
		if len(creature.scannedBy) < 2 {
			return false
		}
	}
	return true
}

func NewChallengeGame() Game[State, Action] {
	return ChallengeGame{
		coords: NewCoordsHelper(),
	}
}
