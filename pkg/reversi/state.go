package reversi

import "strconv"

type State string

// N is the size of the board
const N = 8

func (state *State) Decode() *Board {
	board := make(Board, N)
	for i := 0; i < N; i++ {
		board[i] = make([]*Player, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			s := string([]rune(*state)[i*N+j])
			if s == "_" {
				board[j][i] = nil
			} else {
				x, _ := strconv.Atoi(s)
				board[j][i] = PlayerMap[x]
			}
		}
	}
	return &board
}
