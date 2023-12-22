package main

import (
	"testing"
	"time"
)

func TestRunnerNoTurn(t *testing.T) {
	// prepare
	stdin := NewStdinInterceptor()
	stdin.Intercept()
	runner := NewChallengeRunner()
	quit := make(chan bool)
	go func() {
		quit <- true
	}()
	finalState := runner.Run(quit)
	stdin.Close()
	// assert
	if finalState == nil {
		t.Fatal("Final state should NOT be nil")
	}
}

func TestRunner(t *testing.T) {
	stdin := NewStdinInterceptor()
	stdin.Intercept()
	runner := NewChallengeRunner()
	quit := make(chan bool)
	go func() {
		// turn 1
		stdin.Write("12")     // nb creature
		stdin.Write("2  0 0") // creature 0  (creature id, color, type)
		stdin.Write("3  1 0") // creature 1  (creature id, color, type)
		stdin.Write("4  0 1") // creature 2  (creature id, color, type)
		stdin.Write("5  1 1") // creature 3  (creature id, color, type)
		stdin.Write("6  0 2") // creature 4  (creature id, color, type)
		stdin.Write("7  1 2") // creature 5  (creature id, color, type)
		stdin.Write("8  2 0") // creature 6  (creature id, color, type)
		stdin.Write("9  3 0") // creature 7  (creature id, color, type)
		stdin.Write("10 2 1") // creature 8  (creature id, color, type)
		stdin.Write("11 3 1") // creature 9  (creature id, color, type)
		stdin.Write("12 2 2") // creature 10 (creature id, color, type)
		stdin.Write("13 3 2") // creature 11 (creature id, color, type)
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
		stdin.Write("12")                     // nb visibles creatures
		stdin.Write("2  3467 3925 -141 141 ") // visible creature 0  (id, x, y, vx, vy)
		stdin.Write("3  6532 3925 141  141 ") // visible creature 1  (id, x, y, vx, vy)
		stdin.Write("4  7366 6418 0    -200") // visible creature 2  (id, x, y, vx, vy)
		stdin.Write("5  2633 6418 0    -200") // visible creature 3  (id, x, y, vx, vy)
		stdin.Write("6  1683 8830 -200 0   ") // visible creature 4  (id, x, y, vx, vy)
		stdin.Write("7  8316 8830 200  0   ") // visible creature 5  (id, x, y, vx, vy)
		stdin.Write("8  1767 3560 -200 0   ") // visible creature 6  (id, x, y, vx, vy)
		stdin.Write("9  8232 3560 200  0   ") // visible creature 7  (id, x, y, vx, vy)
		stdin.Write("10 3646 6069 141  -141") // visible creature 8  (id, x, y, vx, vy)
		stdin.Write("11 6353 6069 -141 -141") // visible creature 9  (id, x, y, vx, vy)
		stdin.Write("12 3190 8860 141  141 ") // visible creature 10 (id, x, y, vx, vy)
		stdin.Write("13 6809 8860 -141 141 ") // visible creature 11 (id, x, y, vx, vy)
		stdin.Write("12")                     // radar blip count
		stdin.Write("0 2  BR")                // blip 0  (drone id, creature id, radar)
		stdin.Write("0 3  BR")                // blip 1  (drone id, creature id, radar)
		stdin.Write("0 4  BR")                // blip 2  (drone id, creature id, radar)
		stdin.Write("0 5  BL")                // blip 3  (drone id, creature id, radar)
		stdin.Write("0 6  BL")                // blip 4  (drone id, creature id, radar)
		stdin.Write("0 7  BR")                // blip 5  (drone id, creature id, radar)
		stdin.Write("0 8  BL")                // blip 6  (drone id, creature id, radar)
		stdin.Write("0 9  BR")                // blip 7  (drone id, creature id, radar)
		stdin.Write("0 10 BR")                // blip 8  (drone id, creature id, radar)
		stdin.Write("0 11 BR")                // blip 9  (drone id, creature id, radar)
		stdin.Write("0 12 BL")                // blip 10 (drone id, creature id, radar)
		stdin.Write("0 13 BR")                // blip 11 (drone id, creature id, radar)
		time.Sleep(1000 * time.Millisecond)
		// turn 2
		stdin.Write("0") // my score
		stdin.Write("0") // foe score
		stdin.Write("0") // my scan count
		// no creature
		stdin.Write("0") // foe scan count
		// no creature
		stdin.Write("1")               // my drone count
		stdin.Write("0 3333 800 0 25") // drone 0
		stdin.Write("1")               // foe drone count
		stdin.Write("1 6666 800 0 25") // drone 1
		stdin.Write("0")               // drone scan count
		// no scans
		stdin.Write("12")                     // nb visibles creatures
		stdin.Write("2  3326 4066 -141 141 ") // visible creature 0  (id, x, y, vx, vy)
		stdin.Write("3  6673 4066 141  141 ") // visible creature 1  (id, x, y, vx, vy)
		stdin.Write("4  7366 6218 0    -200") // visible creature 2  (id, x, y, vx, vy)
		stdin.Write("5  2633 6218 0    -200") // visible creature 3  (id, x, y, vx, vy)
		stdin.Write("6  1483 8830 -200 0   ") // visible creature 4  (id, x, y, vx, vy)
		stdin.Write("7  8516 8830 200  0   ") // visible creature 5  (id, x, y, vx, vy)
		stdin.Write("8  1567 3560 -200 0   ") // visible creature 6  (id, x, y, vx, vy)
		stdin.Write("9  8432 3560 200  0   ") // visible creature 7  (id, x, y, vx, vy)
		stdin.Write("10 3787 5928 141  -141") // visible creature 8  (id, x, y, vx, vy)
		stdin.Write("11 6212 5928 -141 -141") // visible creature 9  (id, x, y, vx, vy)
		stdin.Write("12 3331 9001 141  141 ") // visible creature 10 (id, x, y, vx, vy)
		stdin.Write("13 6668 9001 -141 141 ") // visible creature 11 (id, x, y, vx, vy)
		stdin.Write("12")                     // radar blip count
		stdin.Write("0 2  BL")                // blip 0  (drone id, creature id, radar)
		stdin.Write("0 3  BR")                // blip 1  (drone id, creature id, radar)
		stdin.Write("0 4  BR")                // blip 2  (drone id, creature id, radar)
		stdin.Write("0 5  BL")                // blip 3  (drone id, creature id, radar)
		stdin.Write("0 6  BL")                // blip 4  (drone id, creature id, radar)
		stdin.Write("0 7  BR")                // blip 5  (drone id, creature id, radar)
		stdin.Write("0 8  BL")                // blip 6  (drone id, creature id, radar)
		stdin.Write("0 9  BR")                // blip 7  (drone id, creature id, radar)
		stdin.Write("0 10 BR")                // blip 8  (drone id, creature id, radar)
		stdin.Write("0 11 BR")                // blip 9  (drone id, creature id, radar)
		stdin.Write("0 12 BL")                // blip 10 (drone id, creature id, radar)
		stdin.Write("0 13 BR")                // blip 11 (drone id, creature id, radar)
		time.Sleep(100 * time.Millisecond)
		// stop
		quit <- true
	}()
	finalState := runner.Run(quit)
	stdin.Close()
	// assert
	if finalState == nil {
		t.Fatal("Final state should NOT be nil")
	}
}
