package reversi

import (
	"log"
	"math/rand"
	"time"
)

type Reversi struct {
	Agents []Agent      `json:"agents"`
	Env    *Environment `json:"env"`
}

func NewGame(a, b Agent, logLevel int) *Reversi {
	// set random seed
	rand.Seed(time.Now().UnixNano())

	env := NewEnvironment([]*Player{a.Player(), b.Player()})
	return &Reversi{
		Agents: []Agent{a, b},
		Env:    env,
	}
}

func (g *Reversi) Start() {
	var state State
	var action *Action
	agentIdx := 0
	agent := g.Agents[agentIdx]
	board := g.Env.Board
	players := g.Env.Players
	for !board.IsEnd(players) {
		g.Env.Print()
		possibleActions := board.PossibleActions(agent.Player())
		if len(possibleActions) > 0 {
			action = agent.NextAction(g.Env.Board.Encode(), possibleActions)
			log.Printf("action = %+v\n", action)

			// state before take action
			state = board.Encode()

			// update board state
			board.Update(action)

			agent.AppendHistory(state, action, g.Env.Reward(board, action.Player))
		} else {
			log.Printf("you have no possible actions. skipping...")
		}

		// update agent
		agentIdx++
		agent = g.Agents[agentIdx%len(g.Agents)]
	}

	// append the final board state to all the agents
	agent.AppendHistory(state, action, g.Env.Reward(board, action.Player))

	g.Env.Print()
}
