package main

import (
	"fmt"
)

func main() {
	var creatureCount int
	fmt.Scan(&creatureCount)

	for i := 0; i < creatureCount; i++ {
		var creatureId, color, _type int
		fmt.Scan(&creatureId, &color, &_type)
	}
	for {
		var myScore int
		fmt.Scan(&myScore)

		var foeScore int
		fmt.Scan(&foeScore)

		var myScanCount int
		fmt.Scan(&myScanCount)

		for i := 0; i < myScanCount; i++ {
			var creatureId int
			fmt.Scan(&creatureId)
		}
		var foeScanCount int
		fmt.Scan(&foeScanCount)

		for i := 0; i < foeScanCount; i++ {
			var creatureId int
			fmt.Scan(&creatureId)
		}
		var myDroneCount int
		fmt.Scan(&myDroneCount)

		for i := 0; i < myDroneCount; i++ {
			var droneId, droneX, droneY, emergency, battery int
			fmt.Scan(&droneId, &droneX, &droneY, &emergency, &battery)
		}
		var foeDroneCount int
		fmt.Scan(&foeDroneCount)

		for i := 0; i < foeDroneCount; i++ {
			var droneId, droneX, droneY, emergency, battery int
			fmt.Scan(&droneId, &droneX, &droneY, &emergency, &battery)
		}
		var droneScanCount int
		fmt.Scan(&droneScanCount)

		for i := 0; i < droneScanCount; i++ {
			var droneId, creatureId int
			fmt.Scan(&droneId, &creatureId)
		}
		var visibleCreatureCount int
		fmt.Scan(&visibleCreatureCount)

		for i := 0; i < visibleCreatureCount; i++ {
			var creatureId, creatureX, creatureY, creatureVx, creatureVy int
			fmt.Scan(&creatureId, &creatureX, &creatureY, &creatureVx, &creatureVy)
		}
		var radarBlipCount int
		fmt.Scan(&radarBlipCount)

		for i := 0; i < radarBlipCount; i++ {
			var droneId, creatureId int
			var radar string
			fmt.Scan(&droneId, &creatureId, &radar)
		}
		for i := 0; i < myDroneCount; i++ {

			// fmt.Fprintln(os.Stderr, "Debug messages...")
			fmt.Println("WAIT 1") // MOVE <x> <y> <light (1|0)> | WAIT <light (1|0)>
		}
	}
}
