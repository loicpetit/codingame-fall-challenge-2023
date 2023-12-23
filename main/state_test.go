package main

import "testing"

func TestStateString(t *testing.T) {
	text1 := NewState().String()
	if "{nbCreatures:0, creatures:map[], player:, foe:}" != text1 {
		t.Errorf("Unexpected text1: %s", text1)
	}
	text2 := NewState().
		SetCreatures(2, map[int]*Creature{
			1: NewCreature(1, 1, 1),
			2: NewCreature(2, 2, 2),
		}).
		SetPlayer(NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3, 3, 3, 3),
			4: NewDrone(4, 4, 4, 4, 4),
		})).
		SetFoe(NewPlayerState(10, map[int]*Drone{
			5: NewDrone(5, 5, 5, 5, 5),
		})).
		String()
	if "{nbCreatures:2, creatures:map[1:{id:1, color:1, type:1, coords:, scannedBy:[]} 2:{id:2, color:2, type:2, coords:, scannedBy:[]}], player:{score:5, drones:map[3:{id:3, x:3, y:3, emergency:3, battery:3} 4:{id:4, x:4, y:4, emergency:4, battery:4}]}, foe:{score:10, drones:map[5:{id:5, x:5, y:5, emergency:5, battery:5}]}}" != text2 {
		t.Errorf("Unexpected text2: %s", text2)
	}
}

func TestStateSetters(t *testing.T) {
	t.Fatal("TODO")
}
