package main

import (
	"fmt"

	"github.com/chase0213/boardgame/pkg/tictactoe"
)

func Train(N int) {
	a := tictactoe.NewRLAgent(tictactoe.BlackCell, 0.1, 0.9)
	b := tictactoe.NewRLAgent(tictactoe.WhiteCell, 0.1, 0.9)
	// a.LoadValueFunction("./models/tictactoe_100000.json")
	// b.LoadValueFunction("./models/tictactoe_100000.json")
	for i := 0; i < N; i++ {
		fmt.Printf("\n===== Trial  No.%d ====\n", i+1)
		game := tictactoe.NewGame()
		game.Start(a, b)

		a.Update()
		b.Update()
		a.Reset()
		b.Reset()
	}

	path := fmt.Sprintf("./models/tictactoe_%d.json", N)
	a.MergeValueFunction(b)
	a.SaveValueFunction(path)
}

func Play() {
	a := tictactoe.NewHumanAgent(tictactoe.WhiteCell)
	b := tictactoe.NewRLAgent(tictactoe.BlackCell, 0.0, 0.9)
	b.LoadValueFunction("./models/tictactoe_10000.json")

	game := tictactoe.NewGame()
	game.Start(b, a)

	b.Update()
	b.Reset()
}

func main() {
	// N := 10000
	// Train(N)
	Play()
}
