package tictactoe

type Action struct {
	Player Player `json:"player"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}
