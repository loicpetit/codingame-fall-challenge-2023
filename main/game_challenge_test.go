package main

import (
	"testing"
)

func TestGameGetAvailableActionsForEmptyState(t *testing.T) {
	game := NewChallengeGame()
	state := NewState()
	avalaibleActions := game.GetAvailableActions(state, 0)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 0")
	}
	avalaibleActions = game.GetAvailableActions(state, 1)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 1")
	}
	avalaibleActions = game.GetAvailableActions(state, 2)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 2")
	}
	avalaibleActions = game.GetAvailableActions(state, 3)
	if len(avalaibleActions) != 1 {
		t.Errorf("1 action should be return for player 3")
	} else if avalaibleActions[0].Type != "GAME_UPDATE" {
		t.Errorf("GAME_UPDATE action should be return for player 3")
	}
	avalaibleActions = game.GetAvailableActions(state, 4)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 4")
	}
}

func TestGameGetAvailableActions(t *testing.T) {
	game := NewChallengeGame()
	state := NewState().
		SetLastPlayer(2).
		SetCreatures(2, map[int]*Creature{
			1: NewCreature(1, 1, 1).SetCoords(NewCreatureCoords(4000, 8000, 0, -100)),
			2: NewCreature(2, 2, 2).SetCoords(NewCreatureCoords(6000, 6000, 100, 100)),
		}).
		SetPlayer(NewPlayerState(5, map[int]*Drone{
			3: NewDrone(3, 3000, 500, 5, 30, false),
		})).
		SetFoe(NewPlayerState(10, map[int]*Drone{
			4: NewDrone(4, 7000, 500, 5, 30, false),
		}))
	avalaibleActions := game.GetAvailableActions(state, 0)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 0")
	}
	avalaibleActions = game.GetAvailableActions(state, 1)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 1")
	}
	avalaibleActions = game.GetAvailableActions(state, 2)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 2")
	}
	avalaibleActions = game.GetAvailableActions(state, 3)
	if len(avalaibleActions) != 1 {
		t.Errorf("1 action should be return for player 3")
	} else if avalaibleActions[0].Type != "GAME_UPDATE" {
		t.Errorf("GAME_UPDATE action should be return for player 3")
	}
	avalaibleActions = game.GetAvailableActions(state, 4)
	if avalaibleActions != nil {
		t.Errorf("No actions should be return for player 4")
	}
}

func TestGameGetLastPlayer(t *testing.T) {
	game := NewChallengeGame()
	state := NewState()
	lastPlayer := game.GetLastPlayer(state)
	if lastPlayer != 0 {
		t.Error("Last player should be 0")
	}
	state = state.SetLastPlayer(2)
	lastPlayer = game.GetLastPlayer(state)
	if lastPlayer != 2 {
		t.Error("Last player should be 2")
	}
}

func TestGameGetNextPlayer(t *testing.T) {
	game := NewChallengeGame()
	state := NewState()
	nextPlayer := game.GetNextPlayer(state)
	if nextPlayer != 1 {
		t.Error("Initial next player should be 1")
	}
	state = state.SetLastPlayer(1)
	nextPlayer = game.GetNextPlayer(state)
	if nextPlayer != 2 {
		t.Error("Next player should be 2")
	}
	state = state.SetLastPlayer(2)
	nextPlayer = game.GetNextPlayer(state)
	if nextPlayer != 3 {
		t.Error("Next player should be 3")
	}
	state = state.SetLastPlayer(3)
	nextPlayer = game.GetNextPlayer(state)
	if nextPlayer != 1 {
		t.Error("Next player should be 1")
	}
}
