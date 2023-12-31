package main

import "fmt"

type Drone struct {
	Id        int
	X         int
	Y         int
	Emergency int
	Battery   int
	Light     bool
}

func (drone *Drone) String() string {
	if drone == nil {
		return ""
	}
	return fmt.Sprintf("{id:%d, x:%d, y:%d, emergency:%d, battery:%d, light:%t}",
		drone.Id, drone.X, drone.Y, drone.Emergency, drone.Battery, drone.Light)
}

func NewDrone(id, x, y, emergency, battery int, light bool) *Drone {
	return &Drone{
		Id:        id,
		X:         x,
		Y:         y,
		Emergency: emergency,
		Battery:   battery,
		Light:     light,
	}
}
