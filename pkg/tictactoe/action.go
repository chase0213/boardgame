package tictactoe

type Action struct {
	Player Player `json:"player"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func (a *Action) Equivalent(b *Action) bool {
	return a.X == b.X && a.Y == b.Y
}
