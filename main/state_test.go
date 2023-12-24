package main

import "testing"

func TestStateString(t *testing.T) {
	text1 := NewState().String()
	if "{lastPlayer:0, nbCreatures:0, creatures:map[], player:, foe:}" != text1 {
		t.Errorf("Unexpected text1: %s", text1)
	}
	text2 := NewState().
		SetLastPlayer(2).
		SetCreatures(2, map[int]*Creature{
			1: NewCreature(1, 1, 1),
			2: NewCreature(2, 2, 2),
		}).
		SetPlayer(NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3, 3, 3, 3, false),
			4: NewDrone(4, 4, 4, 4, 4, false),
		})).
		SetFoe(NewPlayerState(10, map[int]*Drone{
			5: NewDrone(5, 5, 5, 5, 5, false),
		})).
		String()
	if "{lastPlayer:2, nbCreatures:2, creatures:map[1:{id:1, color:1, type:1, coords:, scannedBy:[]} 2:{id:2, color:2, type:2, coords:, scannedBy:[]}], player:{score:5, drones:map[3:{id:3, x:3, y:3, emergency:3, battery:3, light:false} 4:{id:4, x:4, y:4, emergency:4, battery:4, light:false}]}, foe:{score:10, drones:map[5:{id:5, x:5, y:5, emergency:5, battery:5, light:false}]}}" != text2 {
		t.Errorf("Unexpected text2: %s", text2)
	}
}

func TestStateSetters(t *testing.T) {
	state1 := NewState()
	state2 := state1.
		SetCreatures(2, map[int]*Creature{
			1: NewCreature(1, 1, 1),
			2: NewCreature(2, 2, 2),
		}).
		SetLastPlayer(2)
	state3 := state2.
		SetPlayer(NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3, 3, 3, 3, false),
			4: NewDrone(4, 4, 4, 4, 4, false),
		})).
		SetFoe(NewPlayerState(10, map[int]*Drone{
			5: NewDrone(5, 5, 5, 5, 5, false),
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
	if state1.LastPlayer != 0 {
		t.Errorf("State 1 should have last player 0 but has %d", state1.LastPlayer)
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
	if state2.LastPlayer != 2 {
		t.Errorf("State 2 should have last player 2 but has %d", state2.LastPlayer)
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
	if state3.LastPlayer != 2 {
		t.Errorf("State 3 should have last player 2 but has %d", state3.LastPlayer)
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

func TestPlayerStateFirstDrone(t *testing.T) {
	drone1 := NewDrone(1, 1, 1, 1, 1, false)
	drone2 := NewDrone(2, 2, 2, 2, 2, false)
	player1 := NewPlayerState(5, map[int]*Drone{})
	if nil != player1.GetFirstDrone() {
		t.Errorf("Player 1 should NOT have first drone but has %v", player1.GetFirstDrone())
	}
	player2 := NewPlayerState(5, map[int]*Drone{
		1: drone1,
	})
	if drone1 != player2.GetFirstDrone() {
		t.Errorf("Player 2 should have first drone to be drone1 but has %v", player2.GetFirstDrone())
	}
	player3 := NewPlayerState(5, map[int]*Drone{
		2: drone2,
		3: NewDrone(3, 3, 3, 3, 3, false),
	})
	if drone2 != player3.GetFirstDrone() {
		t.Errorf("Player 3 should have first drone to be drone2 but has %v", player3.GetFirstDrone())
	}
}
