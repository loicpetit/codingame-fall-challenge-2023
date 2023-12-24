package main

import "testing"

func TestActionString(t *testing.T) {
	text1 := NewWaitAction(false).String()
	if text1 != "{type:WAIT, x:0, y:0, light:false}" {
		t.Errorf("Text 1 is unexpected: %s", text1)
	}
	text2 := NewMoveAction(20, 50, true).String()
	if text2 != "{type:MOVE, x:20, y:50, light:true}" {
		t.Errorf("Text 2 is unexpected: %s", text2)
	}
}
