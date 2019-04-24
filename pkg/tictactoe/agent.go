package tictactoe

import "fmt"

type Agent interface {
	Evaluate(State, Action) float64
	NextAction(State) Action
}

func NewHumanAgent(id Player) *HumanAgent {
	return &HumanAgent{
		ID: id,
	}
}

type HumanAgent struct {
	ID Player
}

func (a *HumanAgent) Evaluate(state State, action Action) float64 {
	return 0.0
}

func (a *HumanAgent) NextAction(state State) Action {
	fmt.Printf("Waiting for input of player %d...\n", a.ID)
	var x int
	var y int
	fmt.Printf("x: ")
	fmt.Scan(&x)
	fmt.Printf("y: ")
	fmt.Scan(&y)
	return Action{
		Player: a.ID,
		X:      x,
		Y:      y,
	}
}
