package main

import "testing"

func TestCreatureString(t *testing.T) {
	text1 := NewCreature(1, 2, 3).String()
	if "{id:1, color:2, type:3, coords:, scannedBy:[]}" != text1 {
		t.Errorf("Unexpected text1: %s", text1)
	}
	text2 := NewCreature(1, 2, 3).
		SetCoords(NewCreatureCoords(4, 5, 7, 7)).
		SetScannedBy("player").
		SetScannedBy("foe").
		String()
	if "{id:1, color:2, type:3, coords:{x:4, y:5, vx:7, vy:7}, scannedBy:[player foe]}" != text2 {
		t.Errorf("Unexpected text2: %s", text2)
	}
}

func TestCreatureScans(t *testing.T) {
	creature1 := NewCreature(1, 2, 3)
	creature2 := creature1.SetScannedBy("player").SetScannedBy("foe")
	if creature1.IsScannedBy("player") {
		t.Error("creature1 should NOT be scanned by player")
	}
	if creature1.IsScannedBy("foe") {
		t.Error("creature1 should NOT be scanned by foe")
	}
	if !creature2.IsScannedBy("player") {
		t.Error("creature2 should be scanned by player")
	}
	if !creature2.IsScannedBy("foe") {
		t.Error("creature2 should be scanned by foe")
	}
	if creature2.IsScannedBy("chuck") {
		t.Error("creature2 should NOT be scanned by chuck")
	}
}

func TestCreatureSetters(t *testing.T) {
	creature1 := NewCreature(1, 2, 3)
	creature2 := creature1.SetScannedBy("player").SetScannedBy("foe")
	creature3 := creature1.SetCoords(NewCreatureCoords(2, 2, -1, -1))
	if creature1 == creature2 {
		t.Error("Creature1 should NOT be same as creature 2")
	}
	if creature1 == creature3 {
		t.Error("Creature1 should NOT be same as creature 3")
	}
	if creature2 == creature3 {
		t.Error("Creature2 should NOT be same as creature 3")
	}
	if len(creature1.scannedBy) != 0 {
		t.Errorf("Creature1 scannedBy should be empty but contains values: %v", creature1.scannedBy)
	}
	if len(creature2.scannedBy) != 2 {
		t.Errorf("Creature2 scannedBy should have 2 values but contains: %v", creature2.scannedBy)
	}
	if len(creature3.scannedBy) != 0 {
		t.Errorf("Creature3 scannedBy should be empty but contains values: %v", creature3.scannedBy)
	}
	if creature1.Coords != nil {
		t.Errorf("Creature1 coords should be nil but is: %v", creature1.Coords)
	}
	if creature2.Coords != nil {
		t.Errorf("Creature2 coords should be nil but is: %v", creature2.Coords)
	}
	if creature3.Coords == nil {
		t.Errorf("Creature3 coords should NOT be nil")
	}
}
