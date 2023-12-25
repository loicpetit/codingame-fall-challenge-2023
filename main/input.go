package main

import (
	"fmt"
)

type Input struct {
	NbCreatures int
	Creatures   map[int]*Creature
	Player      *PlayerState
	Foe         *PlayerState
}

func (input Input) String() string {
	return fmt.Sprintf("{nbCreatures:%d, creatures:%v, player:%v, foe:%v}",
		input.NbCreatures, input.Creatures, input.Player, input.Foe)
}

func NewInput(nbCreatures int, creatures map[int]*Creature, player *PlayerState, foe *PlayerState) Input {
	return Input{
		NbCreatures: nbCreatures,
		Creatures:   creatures,
		Player:      player,
		Foe:         foe,
	}
}

type Reader struct {
	isDebug bool
}

// Read implements InputReader.
func (reader Reader) Read() chan Input {
	inputs := make(chan Input)
	go func() {
		first := true
		creatureCount := 0
		creatures := make(map[int]*Creature)
		for {
			creatures = reader.resetCreatures(creatures)
			// first round get init all creatures
			if first {
				fmt.Scan(&creatureCount)
				reader.debug("nb creature", creatureCount)
				for i := 0; i < creatureCount; i++ {
					var id, color, _type int
					fmt.Scan(&id, &color, &_type)
					reader.debug("creature", id, color, _type)
					creatures[id] = NewCreature(id, color, _type)
				}
				first = false
			}
			// player score
			var playerScore int
			fmt.Scan(&playerScore)
			reader.debug("player score", playerScore)
			// foe score
			var foeScore int
			fmt.Scan(&foeScore)
			reader.debug("foe score", foeScore)
			// player scans
			var playerScanCount int
			fmt.Scan(&playerScanCount)
			reader.debug("player scan count", playerScanCount)
			for i := 0; i < playerScanCount; i++ {
				var id int
				fmt.Scan(&id)
				reader.debug("id", id)
				creatures[id] = creatures[id].SetScannedBy("player")
			}
			// foe scans
			var foeScanCount int
			fmt.Scan(&foeScanCount)
			reader.debug("foe scan count", foeScanCount)
			for i := 0; i < foeScanCount; i++ {
				var id int
				fmt.Scan(&id)
				reader.debug("id", id)
				creatures[id] = creatures[id].SetScannedBy("foe")
			}
			// player drones
			var playerDroneCount int
			fmt.Scan(&playerDroneCount)
			reader.debug("player drone count", playerDroneCount)
			playerDrones := make(map[int]*Drone)
			for i := 0; i < playerDroneCount; i++ {
				var id, x, y, emergency, battery int
				fmt.Scan(&id, &x, &y, &emergency, &battery)
				reader.debug("drone id", id, "x", x, "y", y, "emergency", emergency, "battery", battery)
				playerDrones[id] = NewDrone(id, x, y, emergency, battery, false)
			}
			// foe drones
			var foeDroneCount int
			fmt.Scan(&foeDroneCount)
			reader.debug("foe drone count", foeDroneCount)
			foeDrones := make(map[int]*Drone)
			for i := 0; i < foeDroneCount; i++ {
				var id, x, y, emergency, battery int
				fmt.Scan(&id, &x, &y, &emergency, &battery)
				reader.debug("drone id", id, "x", x, "y", y, "emergency", emergency, "battery", battery)
				foeDrones[id] = NewDrone(id, x, y, emergency, battery, false)
			}
			// drone scans
			var droneScanCount int
			fmt.Scan(&droneScanCount)
			reader.debug("drone scan count", droneScanCount)
			for i := 0; i < droneScanCount; i++ {
				var droneId, creatureId int
				fmt.Scan(&droneId, &creatureId)
				reader.debug("drone id", droneId, "creature id", creatureId)
			}
			// visible creatures
			var visibleCreatureCount int
			fmt.Scan(&visibleCreatureCount)
			reader.debug("visible creature count", visibleCreatureCount)
			for i := 0; i < visibleCreatureCount; i++ {
				var id, x, y, vx, vy int
				fmt.Scan(&id, &x, &y, &vx, &vy)
				reader.debug("creature id", id, "x", x, "y", y, "vx", vx, "vy", vy)
				creatures[id] = creatures[id].SetCoords(NewCreatureCoords(x, y, vx, vy))
			}
			// radar blips
			var radarBlipCount int
			fmt.Scan(&radarBlipCount)
			reader.debug("radar blip count", radarBlipCount)
			for i := 0; i < radarBlipCount; i++ {
				var droneId, creatureId int
				var radar string
				fmt.Scan(&droneId, &creatureId, &radar)
				reader.debug("drone id", droneId, "creature id", creatureId, "radar", radar)
			}
			// send inputs
			inputs <- NewInput(
				creatureCount,
				creatures,
				NewPlayerState(playerScore, playerDrones),
				NewPlayerState(foeScore, foeDrones),
			)
		}
	}()
	return inputs
}

// UpdateState implements InputReader.
func (Reader) resetCreatures(creatures map[int]*Creature) map[int]*Creature {
	newCreatures := make(map[int]*Creature)
	for id, creature := range creatures {
		newCreatures[id] = creature.SetCoords(nil)
	}
	return newCreatures
}

// UpdateState implements InputReader.
func (reader Reader) UpdateState(state *State, input Input) *State {
	reader.debug("Update state, input:", input)
	return state.
		SetLastPlayer(2).
		SetCreatures(input.NbCreatures, input.Creatures).
		SetPlayer(input.Player).
		SetFoe(input.Foe)
}

// ValidateAction implements InputReader.
func (Reader) ValidateAction(action *Action, input Input) {
	// no validation
}

func (reader Reader) scan(a ...any) {
	_, err := fmt.Scan(a...)
	if err != nil {
		WriteDebug("input: error scan,", err)
	}
}

func (reader Reader) debug(a ...any) {
	if reader.isDebug {
		WriteDebug(a...)
	}
}

func NewReader() InputReader[Input, State, Action] {
	return Reader{
		isDebug: false,
	}
}
