package reversi

import (
	"fmt"
)

// Environment is a rule of the game
type Environment struct {
	Players []*Player `json:"players"`
	Board   *Board    `json:"board"`
}

func NewEnvironment(players []*Player) *Environment {
	board := NewBoard(players[0])
	return &Environment{
		Players: players,
		Board:   board,
	}
}

func (env *Environment) Winner() *Player {
	n1 := 0
	n2 := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if env.Board.At(j, i) == env.Players[0] {
				n1++
			} else if env.Board.At(j, i) == env.Players[1] {
				n2++
			}
		}
	}

	if n1 > n2 {
		return env.Players[0]
	} else if n1 < n2 {
		return env.Players[1]
	}

	return nil
}

func (env *Environment) Reward(board *Board, player *Player) float64 {
	if !board.IsEnd(env.Players) {
		return 0.0
	}

	if env.Winner() == player {
		return 1.0
	} else if env.Winner() == player.Next {
		return -1.0
	}
	return 0.0
}

func (env *Environment) Print() {
	state := env.Board.Encode()
	fmt.Printf("\nPrinting board with state %s\n", state)
	fmt.Printf("+-+-+-+-+-+-+-+-+\n")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == 0 {
				fmt.Printf("|")
			}
			cell := env.Board.At(j, i)
			if cell == env.Players[0] {
				fmt.Printf("o")
			} else if cell == env.Players[1] {
				fmt.Printf("●")
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf("|")
		}
		fmt.Printf("\n+-+-+-+-+-+-+-+-+\n")
	}
}
