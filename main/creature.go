package main

import "fmt"

type CreatureCoords struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func (coords *CreatureCoords) String() string {
	if coords == nil {
		return ""
	}
	return fmt.Sprintf("{x:%d, y:%d, vx:%d, vy:%d}", coords.X, coords.Y, coords.Vx, coords.Vy)
}

func NewCreatureCoords(x, y, vx, vy int) *CreatureCoords {
	return &CreatureCoords{X: x, Y: y, Vx: vx, Vy: vy}
}

type Creature struct {
	Id        int
	Color     int
	Type      int
	Coords    *CreatureCoords
	scannedBy []string
}

func (creature *Creature) String() string {
	if creature == nil {
		return ""
	}
	return fmt.Sprintf("{id:%d, color:%d, type:%d, coords:%v, scannedBy:%v}",
		creature.Id, creature.Color, creature.Type, creature.Coords, creature.scannedBy)
}

func (creature *Creature) IsScannedBy(value string) bool {
	if creature == nil {
		return false
	}
	for _, scannedBy := range creature.scannedBy {
		if scannedBy == value {
			return true
		}
	}
	return false
}

func (creature *Creature) SetScannedBy(value string) *Creature {
	if creature == nil || creature.IsScannedBy(value) {
		return creature
	}
	return &Creature{
		Id:        creature.Id,
		Color:     creature.Color,
		Type:      creature.Type,
		Coords:    creature.Coords,
		scannedBy: append(creature.scannedBy, value),
	}
}

func (creature *Creature) SetCoords(coords *CreatureCoords) *Creature {
	if creature == nil {
		return nil
	}
	return &Creature{
		Id:        creature.Id,
		Color:     creature.Color,
		Type:      creature.Type,
		Coords:    coords,
		scannedBy: creature.scannedBy,
	}
}

func NewCreature(id int, color int, _type int) *Creature {
	return &Creature{
		Id:        id,
		Color:     color,
		Type:      _type,
		Coords:    nil,
		scannedBy: make([]string, 0),
	}
}
