package main

import (
	"math"
)

// sqrt((3905-3333)² + (3504-500)²) = sqrt(327184 + 9024016) = sqrt(9351200) = 3057.9731849707250171651101275729

// ratio/600 = 5.0966219749512083619418502126216

// nvx = (3905-3333)/5.0966 = 112.23168386767649020915904720794
// nvy = (3504-500)/5.0966 = 589.4125495428324765529961150571

// ntx = 3333 + 112 = 3445
// nty = 500 + 589 = 1089

// nd = sqrt((3445-3333)² + (1089-500)²) = sqrt(12544 + 346921) = sqrt(359465) = 599.55400090400530993269530318661

type CoordsHelper struct{}

// Get distance without applying the square root function
// ideal to just compare distance without needing the concrete distance value
func (CoordsHelper) GetSquaredDistance(x1, y1, x2, y2 int) int {
	deltaX := x2 - x1
	deltaY := y2 - y1
	return (deltaX * deltaX) + (deltaY * deltaY)
}

func (coords CoordsHelper) GetNextCoords(originX, originY, targetX, targetY, maxDistance int) (x, y int) {
	maxSquaredDistance := maxDistance * maxDistance
	squaredDistance := coords.GetSquaredDistance(originX, originY, targetX, targetY)
	if squaredDistance <= maxSquaredDistance {
		// the target is at range, just return target coords
		x = targetX
		y = targetY
		return
	}
	// for the ratio to be correct it needs to be done on distances (not squared distances)
	// sqrt(distance) / sqrt(max) = sqrt(distance / max)
	// so just apply sqrt on the ratio of the squared distances
	ratio := math.Sqrt(float64(maxSquaredDistance) / float64(squaredDistance))
	vx := float64(targetX-originX) * ratio
	vy := float64(targetY-originY) * ratio
	x = originX + int(math.Floor(vx))
	y = originY + int(math.Floor(vy))
	return
}

func NewCoordsHelper() CoordsHelper {
	return CoordsHelper{}
}
