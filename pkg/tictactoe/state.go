package tictactoe

type State int

// N is the size of the board
const N = 3

func InitialState() State {
	return State(0)
}

func (state *State) Decode() *Board {
	board := make(Board, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}

	s := int(*state)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[j][i] = s % N
			s = int(s / N)
		}
	}
	return &board
}
