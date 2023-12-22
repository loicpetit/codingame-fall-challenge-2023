package main

import (
	"fmt"
)

type Input struct{}

func NewInput() Input {
	return Input{}
}

type Reader struct {
	isDebug bool
}

// Read implements InputReader.
func (reader Reader) Read() chan Input {
	inputs := make(chan Input)
	go func() {
		first := true
		for {
			if first {
				var creatureCount int
				fmt.Scan(&creatureCount)
				reader.debug("nb creature", creatureCount)

				for i := 0; i < creatureCount; i++ {
					var creatureId, color, _type int
					fmt.Scan(&creatureId, &color, &_type)
					reader.debug("creature", creatureId, color, _type)
				}
				first = false
			}
			var myScore int
			fmt.Scan(&myScore)
			reader.debug("my score", myScore)

			var foeScore int
			fmt.Scan(&foeScore)
			reader.debug("foe score", foeScore)

			var myScanCount int
			fmt.Scan(&myScanCount)
			reader.debug("my scan count", myScanCount)

			for i := 0; i < myScanCount; i++ {
				var creatureId int
				fmt.Scan(&creatureId)
				reader.debug("creature id", creatureId)
			}
			var foeScanCount int
			fmt.Scan(&foeScanCount)
			reader.debug("foe scan count", foeScanCount)

			for i := 0; i < foeScanCount; i++ {
				var creatureId int
				fmt.Scan(&creatureId)
				reader.debug("creature id", creatureId)
			}
			var myDroneCount int
			fmt.Scan(&myDroneCount)
			reader.debug("my drone count", myDroneCount)

			for i := 0; i < myDroneCount; i++ {
				var droneId, droneX, droneY, emergency, battery int
				fmt.Scan(&droneId, &droneX, &droneY, &emergency, &battery)
				reader.debug("drone id", droneId, "x", droneX, "y", droneY, "emergency", emergency, "battery", battery)
			}
			var foeDroneCount int
			fmt.Scan(&foeDroneCount)
			reader.debug("foe drone count", myDroneCount)

			for i := 0; i < foeDroneCount; i++ {
				var droneId, droneX, droneY, emergency, battery int
				fmt.Scan(&droneId, &droneX, &droneY, &emergency, &battery)
				reader.debug("drone id", droneId, "x", droneX, "y", droneY, "emergency", emergency, "battery", battery)
			}
			var droneScanCount int
			fmt.Scan(&droneScanCount)
			reader.debug("drone scan count", droneScanCount)

			for i := 0; i < droneScanCount; i++ {
				var droneId, creatureId int
				fmt.Scan(&droneId, &creatureId)
				reader.debug("drone id", droneId, "creature id", creatureId)
			}
			var visibleCreatureCount int
			fmt.Scan(&visibleCreatureCount)
			reader.debug("visible creature count", visibleCreatureCount)

			for i := 0; i < visibleCreatureCount; i++ {
				var creatureId, creatureX, creatureY, creatureVx, creatureVy int
				fmt.Scan(&creatureId, &creatureX, &creatureY, &creatureVx, &creatureVy)
				reader.debug("creature id", creatureId, "x", creatureX, "y", creatureY, "vx", creatureVx, "vy", creatureVy)
			}
			var radarBlipCount int
			fmt.Scan(&radarBlipCount)
			reader.debug("radar blip count", radarBlipCount)

			for i := 0; i < radarBlipCount; i++ {
				var droneId, creatureId int
				var radar string
				fmt.Scan(&droneId, &creatureId, &radar)
				reader.debug("drone id", droneId, "creature id", creatureId, "radar", radar)
			}

			inputs <- NewInput()
		}
	}()
	return inputs
}

// UpdateState implements InputReader.
func (Reader) UpdateState(state *State, input Input) *State {
	return state
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
