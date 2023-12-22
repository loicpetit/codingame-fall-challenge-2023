package main

type ChallengeGame struct{}

// GetAvailableActions implements Game.
func (ChallengeGame) GetAvailableActions(state *State, player int) []*Action {
	return nil
}

// GetLastPlayer implements Game.
func (ChallengeGame) GetLastPlayer(state *State) int {
	return -1
}

// GetNextPlayer implements Game.
func (ChallengeGame) GetNextPlayer(state *State) int {
	return -1
}

// Play implements Game.
func (ChallengeGame) Play(state *State, action *Action) *State {
	return state
}

// Start implements Game.
func (ChallengeGame) Start() *State {
	return NewState()
}

// Winner implements Game.
func (ChallengeGame) Winner(state *State) int {
	return 0
}

func NewChallengeGame() Game[State, Action] {
	return ChallengeGame{}
}
