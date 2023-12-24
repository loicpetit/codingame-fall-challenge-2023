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
	state1 := NewState()
	state2 := state1.SetCreatures(2, map[int]*Creature{
		1: NewCreature(1, 1, 1),
		2: NewCreature(2, 2, 2),
	})
	state3 := state2.
		SetPlayer(NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3, 3, 3, 3),
			4: NewDrone(4, 4, 4, 4, 4),
		})).
		SetFoe(NewPlayerState(10, map[int]*Drone{
			5: NewDrone(5, 5, 5, 5, 5),
		}))
	if state1 == state2 {
		t.Error("State 1 should NOT be the same as state 2")
	}
	if state1 == state3 {
		t.Error("State 1 should NOT be the same as state 3")
	}
	if state2 == state3 {
		t.Error("State 2 should NOT be the same as state 3")
	}
	if state1.NbCreatures != 0 {
		t.Errorf("State 1 should have 0 creatures but has %d", state1.NbCreatures)
	}
	if len(state1.Creatures) != 0 {
		t.Errorf("State 1 should have 0 creatures in map but has %d", len(state1.Creatures))
	}
	if state1.Player != nil {
		t.Error("State 1 should NOT have a player")
	}
	if state1.Foe != nil {
		t.Error("State 1 should NOT have a foe")
	}
	if state2.NbCreatures != 2 {
		t.Errorf("State 2 should have 2 creatures but has %d", state2.NbCreatures)
	}
	if len(state2.Creatures) != 2 {
		t.Errorf("State 2 should have 2 creatures in map but has %d", len(state2.Creatures))
	}
	if state2.Player != nil {
		t.Error("State 2 should NOT have a player")
	}
	if state2.Foe != nil {
		t.Error("State 2 should NOT have a foe")
	}
	if state3.NbCreatures != 2 {
		t.Errorf("State 3 should have 2 creatures but has %d", state3.NbCreatures)
	}
	if len(state3.Creatures) != 2 {
		t.Errorf("State 3 should have 2 creatures in map but has %d", len(state3.Creatures))
	}
	if state3.Creatures[0] != state2.Creatures[0] {
		t.Error("State 3 creature 0 should be same as in state 2")
	}
	if state3.Creatures[1] != state2.Creatures[1] {
		t.Error("State 3 creature 1 should be same as in state 2")
	}
	if state3.Player == nil {
		t.Error("State 3 should have a player")
	}
	if state3.Foe == nil {
		t.Error("State 3 should have a foe")
	}
}
