package main

import "testing"

func TestOutputWrite(t *testing.T) {
	interceptor := NewStdoutInterceptor()
	interceptor.Intercept()
	writer := NewWriter()
	writer.Write(NewWaitAction(1, true))
	writer.Write(NewMoveAction(1, 30, -20, false))
	var type1, type2 string
	var light1, light2 int
	var x, y int
	interceptor.Scan(&type1, &light1)
	interceptor.Scan(&type2, &x, &y, &light2)
	interceptor.Close()
	if type1 != "WAIT" {
		t.Errorf("Expected type1 to be WAIT but is %s", type1)
	}
	if light1 != 1 {
		t.Errorf("Expected light1 to be 1 but is %d", light1)
	}
	if type2 != "MOVE" {
		t.Errorf("Expected type2 to be MOVE but is %s", type2)
	}
	if x != 30 {
		t.Errorf("Expected x to be 30 but is %d", x)
	}
	if y != -20 {
		t.Errorf("Expected y to be -20 but is %d", y)
	}
	if light2 != 0 {
		t.Errorf("Expected light2 to be 0 but is %d", light2)
	}
}
