package main

import (
	"fmt"

	"github.com/chase0213/boardgame/pkg/reversi"
)

func Train(N int) {
	players := reversi.NewPlayers(2)
	a := reversi.NewRLAgent(players[0], 0.1, 0.9)
	b := reversi.NewRLAgent(players[1], 0.1, 0.9)
	a.LoadValueFunction("./models/reversi_trained.json")
	b.LoadValueFunction("./models/reversi_trained.json")
	for i := 0; i < N; i++ {
		fmt.Printf("\n===== Trial  No.%d ====\n", i+1)
		game := reversi.NewGame(a, b, 1)
		game.Start()

		a.Update()
		b.Update()
		a.Reset()
		b.Reset()
	}

	path := "./models/reversi_trained.json"
	a.MergeValueFunction(b)
	a.SaveValueFunction(path)
}

func Play() {
	players := reversi.NewPlayers(2)
	a := reversi.NewHumanAgent(players[1])
	b := reversi.NewRLAgent(players[0], 0.0, 0.9)
	b.LoadValueFunction("./models/reversi_trained.json")

	game := reversi.NewGame(b, a, 0)
	game.Start()

	b.Update()
	b.Reset()

	path := "./models/reversi_vs-human.json"
	b.SaveValueFunction(path)
}

func main() {
	// N := 10000
	// Train(N)
	Play()
}
