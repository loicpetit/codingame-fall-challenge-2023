package main

import "testing"

func TestDroneString(t *testing.T) {
	text1 := NewDrone(1, 2, 3, 4, 5).String()
	if "{id:1, x:2, y:3, emergency:4, battery:5}" != text1 {
		t.Errorf("Unexpected text1: %s", text1)
	}
}
