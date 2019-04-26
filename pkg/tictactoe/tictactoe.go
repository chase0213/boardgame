package tictactoe

import "fmt"

type TicTacToe struct {
	State State `json:"state"`
}

func NewGame() *TicTacToe {
	return &TicTacToe{
		State: InitialState(),
	}
}

func (g *TicTacToe) Start(a, b Agent) {
	agent := a
	board := g.State.Decode()
	for !board.IsEnd() {
		board.Print()
		possibleActions := board.PossibleActions()
		action := agent.NextAction(g.State, possibleActions)
		fmt.Printf("action = %+v\n", action)
		board.Update(action)

		state := board.Encode()
		agent.AppendHistory(state)
		g.State = state

		// update agent
		if agent == a {
			agent = b
		} else {
			agent = a
		}

		// update board
		board = g.State.Decode()
	}
	board.Print()
}
