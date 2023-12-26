package main

import (
	"testing"
	"time"
)

func TestInputString(t *testing.T) {
	text := NewInput(
		2,
		map[int]*Creature{
			1: NewCreature(1, 1, 1),
			2: NewCreature(2, 2, 2),
		},
		NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3, 3, 3, 3, false),
			4: NewDrone(4, 4, 4, 4, 4, false),
		}),
		NewPlayerState(10, map[int]*Drone{
			5: NewDrone(5, 5, 5, 5, 5, false),
		}),
	).String()
	if text != "{nbCreatures:2, creatures:map[1:{id:1, color:1, type:1, coords:, scannedBy:[]} 2:{id:2, color:2, type:2, coords:, scannedBy:[]}], player:{score:5, drones:map[3:{id:3, x:3, y:3, emergency:3, battery:3, light:false} 4:{id:4, x:4, y:4, emergency:4, battery:4, light:false}]}, foe:{score:10, drones:map[5:{id:5, x:5, y:5, emergency:5, battery:5, light:false}]}}" {
		t.Errorf("Unexpected text: %s", text)
	}
}

func TestReaderRead(t *testing.T) {
	stdin := NewStdinInterceptor()
	stdin.Intercept()
	go func() {
		// turn 1
		stdin.Write("2")      // nb creature
		stdin.Write("2  0 0") // creature 0  (creature id, color, type)
		stdin.Write("3  1 0") // creature 1  (creature id, color, type)
		stdin.Write("0")      // my score
		stdin.Write("0")      // foe score
		stdin.Write("0")      // my scan count
		// no creature
		stdin.Write("0") // foe scan count
		// no creature
		stdin.Write("1")               // my drone count
		stdin.Write("0 3333 500 0 30") // drone 0
		stdin.Write("1")               // foe drone count
		stdin.Write("1 6666 500 0 30") // drone 1
		stdin.Write("0")               // drone scan count
		// no scans
		stdin.Write("2")                      // nb visibles creatures
		stdin.Write("2  3467 3925 -141 141 ") // visible creature 0  (id, x, y, vx, vy)
		stdin.Write("3  6532 3925 141  141 ") // visible creature 1  (id, x, y, vx, vy)
		stdin.Write("2")                      // radar blip count
		stdin.Write("0 2  BR")                // blip 0  (drone id, creature id, radar)
		stdin.Write("0 3  BR")                // blip 1  (drone id, creature id, radar)
		time.Sleep(1000 * time.Millisecond)
		// turn 2
		stdin.Write("0") // my score
		stdin.Write("0") // foe score
		stdin.Write("1") // my scan count
		stdin.Write("2") // creature 2 scanned by player
		stdin.Write("1") // foe scan count
		stdin.Write("3") // creature 3 scanned by foe
		// no creature
		stdin.Write("1")               // my drone count
		stdin.Write("0 3333 800 0 25") // drone 0
		stdin.Write("1")               // foe drone count
		stdin.Write("1 6666 800 0 25") // drone 1
		stdin.Write("0")               // drone scan count
		// no scans
		stdin.Write("2")                      // nb visibles creatures
		stdin.Write("2  3326 4066 -141 141 ") // visible creature 0  (id, x, y, vx, vy)
		stdin.Write("3  6673 4066 141  141 ") // visible creature 1  (id, x, y, vx, vy)
		stdin.Write("2")                      // radar blip count
		stdin.Write("0 2  BL")                // blip 0  (drone id, creature id, radar)
		stdin.Write("0 3  BR")                // blip 1  (drone id, creature id, radar)
		time.Sleep(100 * time.Millisecond)
	}()
	inputs := NewReader().Read()
	input1 := <-inputs
	input2 := <-inputs
	stdin.Close()
	// assert
	if input1.NbCreatures != 2 {
		t.Errorf("Input 1 nb creatures should be 2 but is %d", input1.NbCreatures)
	}
	if len(input1.Creatures) != 2 {
		t.Errorf("Input 1 nb creatures in map should be 2 but is %d", len(input1.Creatures))
	} else {
		if input1.Creatures[2].Id != 2 {
			t.Errorf("Input 1 creature[2] should have id 2 but is %d", input1.Creatures[2].Id)
		}
		if input1.Creatures[2].Color != 0 {
			t.Errorf("Input 1 creature[2] should have color 0 but is %d", input1.Creatures[2].Color)
		}
		if input1.Creatures[2].Type != 0 {
			t.Errorf("Input 1 creature[2] should have type 0 but is %d", input1.Creatures[2].Type)
		}
		if input1.Creatures[2].Coords.X != 3467 {
			t.Errorf("Input 1 creature[2] should have x 3467 but is %d", input1.Creatures[2].Coords.X)
		}
		if input1.Creatures[2].Coords.Y != 3925 {
			t.Errorf("Input 1 creature[2] should have y 3925 but is %d", input1.Creatures[2].Coords.Y)
		}
		if input1.Creatures[2].Coords.Vx != -141 {
			t.Errorf("Input 1 creature[2] should have vx -141 but is %d", input1.Creatures[2].Coords.Vx)
		}
		if input1.Creatures[2].Coords.Vy != 141 {
			t.Errorf("Input 1 creature[2] should have vy 141 but is %d", input1.Creatures[2].Coords.Vy)
		}
		if len(input1.Creatures[2].scannedBy) != 0 {
			t.Errorf("Input 1 creature[2] should have nb scanned by 0 but is %d", len(input1.Creatures[2].scannedBy))
		}
		if input1.Creatures[3].Id != 3 {
			t.Errorf("Input 1 creature[3] should have id 3 but is %d", input1.Creatures[3].Id)
		}
		if input1.Creatures[3].Color != 1 {
			t.Errorf("Input 1 creature[3] should have color 1 but is %d", input1.Creatures[3].Color)
		}
		if input1.Creatures[3].Type != 0 {
			t.Errorf("Input 1 creature[3] should have type 0 but is %d", input1.Creatures[3].Type)
		}
		if input1.Creatures[3].Coords.X != 6532 {
			t.Errorf("Input 1 creature[3] should have x 6532 but is %d", input1.Creatures[3].Coords.X)
		}
		if input1.Creatures[3].Coords.Y != 3925 {
			t.Errorf("Input 1 creature[3] should have y 3925 but is %d", input1.Creatures[3].Coords.Y)
		}
		if input1.Creatures[3].Coords.Vx != 141 {
			t.Errorf("Input 1 creature[3] should have vx 141 but is %d", input1.Creatures[3].Coords.Vx)
		}
		if input1.Creatures[3].Coords.Vy != 141 {
			t.Errorf("Input 1 creature[3] should have vy 141 but is %d", input1.Creatures[3].Coords.Vy)
		}
		if len(input1.Creatures[3].scannedBy) != 0 {
			t.Errorf("Input 1 creature[3] should have nb scanned by 0 but is %d", len(input1.Creatures[3].scannedBy))
		}
	}
	if input2.NbCreatures != 2 {
		t.Errorf("Input 2 nb creatures should be 2 but is %d", input2.NbCreatures)
	} else {
		if input2.Creatures[2].Id != 2 {
			t.Errorf("Input 2 creature[2] should have id 2 but is %d", input2.Creatures[2].Id)
		}
		if input2.Creatures[2].Color != 0 {
			t.Errorf("Input 2 creature[2] should have color 0 but is %d", input2.Creatures[2].Color)
		}
		if input2.Creatures[2].Type != 0 {
			t.Errorf("Input 2 creature[2] should have type 0 but is %d", input2.Creatures[2].Type)
		}
		if input2.Creatures[2].Coords.X != 3326 {
			t.Errorf("Input 2 creature[2] should have x 3326 but is %d", input2.Creatures[2].Coords.X)
		}
		if input2.Creatures[2].Coords.Y != 4066 {
			t.Errorf("Input 2 creature[2] should have y 4066 but is %d", input2.Creatures[2].Coords.Y)
		}
		if input2.Creatures[2].Coords.Vx != -141 {
			t.Errorf("Input 2 creature[2] should have vx -141 but is %d", input2.Creatures[2].Coords.Vx)
		}
		if input2.Creatures[2].Coords.Vy != 141 {
			t.Errorf("Input 2 creature[2] should have vy 141 but is %d", input2.Creatures[2].Coords.Vy)
		}
		if len(input2.Creatures[2].scannedBy) != 1 {
			t.Errorf("Input 2 creature[2] should have nb scanned by 1 but is %d", len(input2.Creatures[2].scannedBy))
		}
		if input2.Creatures[2].scannedBy[0] != "player" {
			t.Errorf("Input 2 creature[2] should be scanned by player but is %s", input2.Creatures[2].scannedBy[0])
		}
		if input2.Creatures[3].Id != 3 {
			t.Errorf("Input 2 creature[3] should have id 3 but is %d", input2.Creatures[3].Id)
		}
		if input2.Creatures[3].Color != 1 {
			t.Errorf("Input 2 creature[3] should have color 1 but is %d", input2.Creatures[3].Color)
		}
		if input2.Creatures[3].Type != 0 {
			t.Errorf("Input 2 creature[3] should have type 0 but is %d", input2.Creatures[3].Type)
		}
		if input2.Creatures[3].Coords.X != 6673 {
			t.Errorf("Input 2 creature[3] should have x 6673 but is %d", input2.Creatures[3].Coords.X)
		}
		if input2.Creatures[3].Coords.Y != 4066 {
			t.Errorf("Input 2 creature[3] should have y 4066 but is %d", input2.Creatures[3].Coords.Y)
		}
		if input2.Creatures[3].Coords.Vx != 141 {
			t.Errorf("Input 2 creature[3] should have vx 141 but is %d", input2.Creatures[3].Coords.Vx)
		}
		if input2.Creatures[3].Coords.Vy != 141 {
			t.Errorf("Input 2 creature[3] should have vy 141 but is %d", input2.Creatures[3].Coords.Vy)
		}
		if len(input2.Creatures[3].scannedBy) != 1 {
			t.Errorf("Input 2 creature[3] should have nb scanned by 1 but is %d", len(input2.Creatures[3].scannedBy))
		}
		if input2.Creatures[3].scannedBy[0] != "foe" {
			t.Errorf("Input 2 creature[3] should be scanned by foe but is %s", input2.Creatures[3].scannedBy[0])
		}
	}
}

func TestReaderUpdateState(t *testing.T) {
	baseState := NewState()
	state := NewReader().
		UpdateState(
			baseState,
			NewInput(
				1,
				map[int]*Creature{3: NewCreature(3, 0, 0)},
				NewPlayerState(0, map[int]*Drone{}),
				NewPlayerState(0, map[int]*Drone{}),
			),
		)
	if baseState == state {
		t.Error("State should be different of base state")
	}
	if baseState.LastPlayer != 0 {
		t.Error("Base state last player should be 0")
	}
	if baseState.NbCreatures != 0 {
		t.Error("Base state nb creatures should be 0")
	}
	if len(baseState.Creatures) != 0 {
		t.Error("Base state nb creatures in map should be 0")
	}
	if baseState.Player != nil {
		t.Error("Base state player should be nil")
	}
	if baseState.Foe != nil {
		t.Error("Base state foe should be nil")
	}
	if baseState.Round != 0 {
		t.Error("Base state round should be 0")
	}
	if state.LastPlayer != 3 {
		t.Error("State last player should be 3")
	}
	if state.NbCreatures != 1 {
		t.Error("State nb creatures should be 1")
	}
	if len(state.Creatures) != 1 {
		t.Error("State nb creatures in map should be 1")
	}
	if state.Player == nil {
		t.Error("State player should NOT be nil")
	}
	if state.Foe == nil {
		t.Error("State foe should NOT be nil")
	}
	if state.Round != 1 {
		t.Error("State round should be 0")
	}
}
