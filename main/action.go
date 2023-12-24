package main

import "fmt"

type Action struct {
	Player int // player 1 = player, player 2 = foe, player 3 = game update
	Type   string
	X      int
	Y      int
	Light  bool
}

func (action *Action) String() string {
	if action == nil {
		return ""
	}
	return fmt.Sprintf("{player:%d, type:%s, x:%d, y:%d, light:%t}", action.Player, action.Type, action.X, action.Y, action.Light)
}

func NewMoveAction(player, x, y int, light bool) *Action {
	if player < 1 || player > 2 {
		panic("Player must be 1 or 2")
	}
	return &Action{
		Player: player,
		Type:   "MOVE",
		X:      x,
		Y:      y,
		Light:  light,
	}
}

func NewWaitAction(player int, light bool) *Action {
	if player < 1 || player > 2 {
		panic("Player must be 1 or 2")
	}
	return &Action{
		Player: player,
		Type:   "WAIT",
		X:      0,
		Y:      0,
		Light:  light,
	}
}

func NewGameUpdateAction() *Action {
	return &Action{
		Player: 3,
		Type:   "GAME_UPDATE",
		X:      0,
		Y:      0,
		Light:  false,
	}
}
