package reversi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type StateWithReward struct {
	State  State   `json:"state"`
	Action *Action `json:"action"`
	Reward float64 `json:"reward"`
}

type StateHistory []StateWithReward

func (ss StateHistory) Less(i, j int) bool {
	return ss[i].State < ss[j].State
}

func (ss StateHistory) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss StateHistory) Len() int {
	return len(ss)
}

type Agent interface {
	Player() *Player
	Evaluate(State, []*Action) []float64
	NextAction(State, []*Action) *Action
	AppendHistory(State, *Action, float64)

	LoadValueFunction(string) error
	SaveValueFunction(string) error
}

func NewHumanAgent(p *Player) *HumanAgent {
	return &HumanAgent{
		player: p,

		StateHistory: StateHistory{},
	}
}

type HumanAgent struct {
	player *Player `json:"player"`

	StateHistory StateHistory `json:"state_history"`
}

func (a *HumanAgent) Player() *Player {
	return a.player
}

func (a *HumanAgent) Evaluate(state State, possibleActions []*Action) []float64 {
	return []float64{0.0}
}

func (a *HumanAgent) NextAction(state State, possibleActions []*Action) *Action {
	log.Printf("Waiting for input of player %d...\n", a.player.ID)

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
			return action
		}

		fmt.Printf("invalid number!\n")
	}
	return nil
}

func (a *HumanAgent) AppendHistory(state State, action *Action, reward float64) {
	a.StateHistory = append(a.StateHistory, StateWithReward{
		State:  state,
		Action: action,
		Reward: reward,
	})
}

func (a *HumanAgent) SaveValueFunction(path string) error {
	return nil
}

func (a *HumanAgent) LoadValueFunction(path string) error {
	return nil
}

type RLAgent struct {
	player *Player

	Epsilon      float64           `json:"epsilon"`
	Gamma        float64           `json:"gamma"`
	StateHistory StateHistory      `json:"state_history"`
	V            map[State]float64 `json:"v"`
}

func NewRLAgent(p *Player, epsilon, gamma float64) *RLAgent {
	// set seed
	rand.Seed(time.Now().UnixNano())

	return &RLAgent{
		player: p,

		Epsilon:      epsilon,
		Gamma:        gamma,
		StateHistory: StateHistory{},
		V:            map[State]float64{},
	}
}

func (a *RLAgent) Player() *Player {
	return a.player
}

func (a *RLAgent) Evaluate(state State, possibleActions []*Action) []float64 {
	values := make([]float64, len(possibleActions))
	for i, action := range possibleActions {
		key := State(fmt.Sprintf("%s--%d_%d", state, action.X, action.Y))
		board := state.Decode()
		board.Update(action)
		value := a.V[key]
		values[i] = value
		log.Printf("V(%d, %d) = %f\n", action.X, action.Y, value)
	}

	return values
}

func (a *RLAgent) NextAction(state State, possibleActions []*Action) *Action {
	dice := rand.Float64()

	// choose at random
	if dice < a.Epsilon {
		log.Printf("Choosing by epsilon part.\n")
		x := rand.Int() % len(possibleActions)
		action := possibleActions[x]
		return action
	}

	log.Printf("Choosing by greedy part.\n")
	values := a.Evaluate(state, possibleActions)
	x := 0
	for i, value := range values {
		if values[x] < value {
			x = i
		}
	}
	action := possibleActions[x]
	return action
}

func (a *RLAgent) AppendHistory(state State, action *Action, reward float64) {
	a.StateHistory = append(a.StateHistory, StateWithReward{
		State:  state,
		Action: action,
		Reward: reward,
	})
}

func (a *RLAgent) Update() {
	prevV := 0.0

	n := len(a.StateHistory)
	for i := n - 1; i >= 0; i-- {
		history := a.StateHistory[i]
		state := history.State
		action := history.Action
		reward := history.Reward

		key := State(fmt.Sprintf("%s--%d_%d", state, action.X, action.Y))

		a.V[key] = a.V[key] + (reward + a.Gamma*prevV - a.V[key])
		prevV = a.V[key]
	}
}

func (a *RLAgent) Reset() {
	a.StateHistory = StateHistory{}
}

func (a *RLAgent) MergeValueFunction(b *RLAgent) {
	for key, value := range b.V {
		if v, ok := a.V[key]; ok {
			a.V[key] = (v + value) / 2
		} else {
			a.V[key] = value
		}
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
