package main

import "fmt"

type PlayerState struct {
	Score  int
	Drones map[int]*Drone
}

func (player *PlayerState) String() string {
	if player == nil {
		return ""
	}
	return fmt.Sprintf("{score:%d, drones:%v}", player.Score, player.Drones)
}

func NewPlayerState(score int, drones map[int]*Drone) *PlayerState {
	return &PlayerState{
		Score:  score,
		Drones: drones,
	}
}

type State struct {
	NbCreatures int
	Creatures   map[int]*Creature
	Player      *PlayerState
	Foe         *PlayerState
}

func (state *State) String() string {
	if state == nil {
		return ""
	}
	return fmt.Sprintf("{nbCreatures:%d, creatures:%v, player:%v, foe:%v}",
		state.NbCreatures, state.Creatures, state.Player, state.Foe)
}

func (state *State) SetCreatures(nbCreatures int, creatures map[int]*Creature) *State {
	if state == nil {
		return nil
	}
	return &State{
		NbCreatures: nbCreatures,
		Creatures:   creatures,
		Player:      state.Player,
		Foe:         state.Foe,
	}
}

func (state *State) SetPlayer(player *PlayerState) *State {
	if state == nil {
		return nil
	}
	return &State{
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      player,
		Foe:         state.Foe,
	}
}

func (state *State) SetFoe(foe *PlayerState) *State {
	if state == nil {
		return nil
	}
	return &State{
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      state.Player,
		Foe:         foe,
	}
}

func NewState() *State {
	return &State{}
}
