package main

import (
	"testing"
)

func TestCoordsHelperSquaredDistance(t *testing.T) {
	coords := NewCoordsHelper()
	dataSet := []struct {
		testName                                                    string
		originX, originY, targetX, targetY, expectedSquaredDistance int
	}{
		{"Zero", 0, 0, 0, 0, 0},
		{"FromZero", 0, 0, 30, 40, 2500},
		{"FromZeroNegative", 0, 0, -30, -40, 2500},
		{"Positive", 1000, 2000, 1300, 2400, 250000},
		{"Negative", 1000, 2000, 700, 1600, 250000},
		{"Mix", 1000, 2000, 700, 2400, 250000},
	}
	for _, data := range dataSet {
		t.Run(data.testName, func(t *testing.T) {
			squaredDistance := coords.GetSquaredDistance(data.originX, data.originY, data.targetX, data.targetY)
			if data.expectedSquaredDistance != squaredDistance {
				t.Errorf("Expected squared distance is %d but get %d", data.expectedSquaredDistance, squaredDistance)
			}
		})
	}
}

func TestCoordsHelperGetNextCoords(t *testing.T) {
	coords := NewCoordsHelper()
	dataSet := []struct {
		testName                                                 string
		originX, originY, targetX, targetY, expectedX, expectedY int
	}{
		// Use Pythagorean triple to find integer results for distances
		// 600 is multiple of 5
		// so from triple 3,4,5, we get 360,480,600 (multiple of 120)
		// to get coords too far but in same proportion, multiply 3,4 by something greater than 120 (here 125)
		// like that next coords still are +- 360,480
		{"Zero", 0, 0, 0, 0, 0, 0},
		{"Max", 0, 0, 360, 480, 360, 480},
		{"MaxPlusOne", 0, 0, 361, 481, 360, 479},
		{"Negative", 1000, 2000, 700, 1600, 700, 1600},
		{"NegativeTooFar", 1000, 2000, 625, 1500, 640, 1520},
		{"Mix", 1000, 2000, 700, 2400, 700, 2400},
		{"MixTooFar", 1000, 2000, 625, 2500, 640, 2480},
	}
	for _, data := range dataSet {
		t.Run(data.testName, func(t *testing.T) {
			x, y := coords.GetNextCoords(data.originX, data.originY, data.targetX, data.targetY, 600)
			if data.expectedX != x {
				t.Errorf("Expected x is %d but get %d", data.expectedX, x)
			}
			if data.expectedY != y {
				t.Errorf("Expected y is %d but get %d", data.expectedY, y)
			}
		})
	}
}
