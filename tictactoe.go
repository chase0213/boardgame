package main

import (
	"github.com/chase0213/boardgame/pkg/tictactoe"
)

func main() {
	game := tictactoe.NewGame()
	a := tictactoe.NewHumanAgent(tictactoe.BlackCell)
	b := tictactoe.NewHumanAgent(tictactoe.WhiteCell)
	game.Start(a, b)
}
