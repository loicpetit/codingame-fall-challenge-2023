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

func (player *PlayerState) GetFirstDrone() *Drone {
	if player == nil {
		return nil
	}
	for _, drone := range player.Drones {
		return drone
	}
	return nil
}

func NewPlayerState(score int, drones map[int]*Drone) *PlayerState {
	return &PlayerState{
		Score:  score,
		Drones: drones,
	}
}

type State struct {
	LastPlayer  int // player 1 = player, player 2 = foe, player 3 = game update
	NbCreatures int
	Creatures   map[int]*Creature
	Player      *PlayerState
	Foe         *PlayerState
	Round       int
}

func (state *State) String() string {
	if state == nil {
		return ""
	}
	return fmt.Sprintf("{lastPlayer:%d, nbCreatures:%d, creatures:%v, player:%v, foe:%v, round:%d}",
		state.LastPlayer, state.NbCreatures, state.Creatures, state.Player, state.Foe, state.Round)
}

func (state *State) SetLastPlayer(player int) *State {
	if state == nil {
		return nil
	}
	return &State{
		LastPlayer:  player,
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      state.Player,
		Foe:         state.Foe,
		Round:       state.Round,
	}
}

func (state *State) SetCreatures(nbCreatures int, creatures map[int]*Creature) *State {
	if state == nil {
		return nil
	}
	return &State{
		LastPlayer:  state.LastPlayer,
		NbCreatures: nbCreatures,
		Creatures:   creatures,
		Player:      state.Player,
		Foe:         state.Foe,
		Round:       state.Round,
	}
}

func (state *State) SetPlayer(player *PlayerState) *State {
	if state == nil {
		return nil
	}
	return &State{
		LastPlayer:  state.LastPlayer,
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      player,
		Foe:         state.Foe,
		Round:       state.Round,
	}
}

func (state *State) SetFoe(foe *PlayerState) *State {
	if state == nil {
		return nil
	}
	return &State{
		LastPlayer:  state.LastPlayer,
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      state.Player,
		Foe:         foe,
		Round:       state.Round,
	}
}

func (state *State) IncreaseRound() *State {
	if state == nil {
		return nil
	}
	return &State{
		LastPlayer:  state.LastPlayer,
		NbCreatures: state.NbCreatures,
		Creatures:   state.Creatures,
		Player:      state.Player,
		Foe:         state.Foe,
		Round:       state.Round + 1,
	}
}

func NewState() *State {
	return &State{
		Round: 0,
	}
}
