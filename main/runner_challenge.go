package main

type ChallengeRunner struct{}

func NewChallengeRunner() Runner[Input, State, Action] {
	return NewRunner[Input, State, Action](
		NewChallengeGame(),
		NewReader(),
		NewStrategyDoNothing(),
		NewWriter(),
	)
}
