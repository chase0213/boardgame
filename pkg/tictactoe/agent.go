package tictactoe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"time"
)

type States []State

func (ss States) Less(i, j int) bool {
	return ss[i] < ss[j]
}

func (ss States) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss States) Len() int {
	return len(ss)
}

type Agent interface {
	Evaluate(State, []*Action) []float64
	NextAction(State, []*Action) *Action
	AppendHistory(State)

	LoadValueFunction(string) error
	SaveValueFunction(string) error
}

func NewHumanAgent(id Player) *HumanAgent {
	return &HumanAgent{
		ID: id,

		StateHistory: States{},
	}
}

type HumanAgent struct {
	ID Player `json:"id"`

	StateHistory States `json:"state_history"`
}

func (a *HumanAgent) Evaluate(state State, possibleActions []*Action) []float64 {
	return []float64{0.0}
}

func (a *HumanAgent) NextAction(state State, possibleActions []*Action) *Action {
	fmt.Printf("Waiting for input of player %d...\n", a.ID)

	for true {
		fmt.Printf("Enter the number of actions below:\n")
		for i, action := range possibleActions {
			fmt.Printf("%d. (x, y) = (%d, %d)\n", i, action.X, action.Y)
		}

		var c int
		fmt.Printf("No.: ")
		fmt.Scan(&c)

		if c >= 0 && c < len(possibleActions) {
			action := possibleActions[c]
			action.Player = a.ID
			return action
		}

		fmt.Printf("invalid number!\n")
	}
	return nil
}

func (a *HumanAgent) AppendHistory(state State) {
	a.StateHistory = append(a.StateHistory, state)
}

func (a *HumanAgent) SaveValueFunction(path string) error {
	return nil
}

func (a *HumanAgent) LoadValueFunction(path string) error {
	return nil
}

type RLAgent struct {
	ID Player

	Epsilon      float64           `json:"epsilon"`
	Gamma        float64           `json:"gamma"`
	StateHistory States            `json:"state_history"`
	V            map[State]float64 `json:"v"`
}

func NewRLAgent(id Player, epsilon, gamma float64) *RLAgent {
	// set seed
	rand.Seed(time.Now().UnixNano())

	return &RLAgent{
		ID: id,

		Epsilon:      epsilon,
		Gamma:        gamma,
		StateHistory: States{},
		V:            map[State]float64{},
	}
}

func (a *RLAgent) Evaluate(state State, possibleActions []*Action) []float64 {
	values := make([]float64, len(possibleActions))
	for i, action := range possibleActions {
		action.Player = a.ID
		board := state.Decode()
		board.Update(action)
		newState := board.Encode()
		value := a.V[newState]
		values[i] = value
		fmt.Printf("V(%d, %d) = %f\n", action.X, action.Y, value)
	}

	return values
}

func (a *RLAgent) NextAction(state State, possibleActions []*Action) *Action {
	dice := rand.Float64()

	// choose at random
	if dice < a.Epsilon {
		fmt.Printf("Choosing by epsilon part.\n")
		x := rand.Int() % len(possibleActions)
		action := possibleActions[x]
		action.Player = a.ID
		return action
	}

	fmt.Printf("Choosing by greedy part.\n")
	values := a.Evaluate(state, possibleActions)
	x := 0
	for i, value := range values {
		if values[x] < value {
			x = i
		}
	}
	action := possibleActions[x]
	action.Player = a.ID
	return action
}

func (a *RLAgent) AppendHistory(state State) {
	a.StateHistory = append(a.StateHistory, state)
}

func (a *RLAgent) Update() {
	var prevV float64
	sort.Sort(sort.Reverse(a.StateHistory))

	for _, state := range a.StateHistory {
		board := state.Decode()

		// in case the board is a terminal state
		if board.IsEnd() {
			a.V[state] = board.Reward(a.ID)
			prevV = a.V[state]
			fmt.Printf("a.V[%d] = %f\n", state, a.V[state])
			continue
		}

		a.V[state] = (a.V[state] + a.Gamma*prevV) / (1.0 + a.Gamma)
	}
}

func (a *RLAgent) Reset() {
	a.StateHistory = []State{}
}

func (a *RLAgent) MergeValueFunction(b *RLAgent) {
	for key, value := range b.V {
		a.V[key] = value
	}
}

func (a *RLAgent) SaveValueFunction(path string) error {
	var data []byte
	data, err := json.Marshal(a.V)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func (a *RLAgent) LoadValueFunction(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &a.V); err != nil {
		return err
	}

	return nil
}
