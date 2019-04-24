package tictactoe

type Board [][]int

func (b *Board) At(x, y int) int {
	return (*b)[y][x]
}

func (b *Board) IsValidPlay(a Action) bool {
	if a.X < 0 || a.X >= N || a.Y < 0 || a.Y >= N {
		return false
	}
	return b.At(a.X, a.Y) == EmptyCell
}

func (board *Board) Encode() State {
	basis := 1
	state := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			state += board.At(i, j) * basis
			basis *= N
		}
	}
	return State(state)
}

func (board *Board) Update(action Action) {
	(*board)[action.Y][action.X] = int(action.Player)
}
