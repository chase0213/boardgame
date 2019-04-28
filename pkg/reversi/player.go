package reversi

const Empty = 0

type Player struct {
	ID   int     `json:"id"`
	Next *Player `json:"next"`
}

var PlayerMap map[int]*Player

func NewPlayers(n int) []*Player {
	players := make([]*Player, n)
	for i := 0; i < n; i++ {
		players[i] = &Player{
			ID:   i + 1,
			Next: nil,
		}
	}

	PlayerMap = map[int]*Player{}
	for i := 0; i < n; i++ {
		players[i].Next = players[(i+1)%n]
		PlayerMap[players[i].ID] = players[i]
	}

	return players
}
