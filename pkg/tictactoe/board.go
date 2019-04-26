package tictactoe

import "fmt"

type Board [][]int

func (b *Board) At(x, y int) int {
	return (*b)[y][x]
}

func (b *Board) IsValidPlay(a *Action) bool {
	possibleActions := b.PossibleActions()
	for _, action := range possibleActions {
		if action.Equivalent(a) {
			return true
		}
	}
	return false
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

func (board *Board) Update(action *Action) {
	(*board)[action.Y][action.X] = int(action.Player)
}

func (board *Board) Winner() Player {
	isFilled := true

	// check column
	for i := 0; i < N; i++ {
		winner := board.At(i, 0)
		c := 0
		for j := 0; j < N; j++ {
			cell := board.At(i, j)

			// check if cell is empty
			if cell == EmptyCell {
				isFilled = false
				continue
			}

			if winner == cell {
				c++
			}
		}
		if c == N {
			return Player(winner)
		}
	}

	// check row
	for i := 0; i < N; i++ {
		winner := board.At(0, i)
		c := 0
		for j := 0; j < N; j++ {
			cell := board.At(j, i)
			// check if cell is empty
			if cell == EmptyCell {
				isFilled = false
				continue
			}

			if winner == cell {
				c++
			}
		}
		if c == N {
			return Player(winner)
		}
	}

	// diag: left top to right bottom
	winner := board.At(0, 0)
	c := 0
	for i := 0; i < N; i++ {
		cell := board.At(i, i)
		// check if cell is empty
		if cell == EmptyCell {
			isFilled = false
			continue
		}

		if winner == cell {
			c++
		}
	}
	if c == N {
		return Player(winner)
	}

	// diag: right top to left bottom
	winner = board.At(N-1, 0)
	c = 0
	for i := 0; i < N; i++ {
		cell := board.At(N-i-1, i)
		// check if cell is empty
		if cell == EmptyCell {
			isFilled = false
			continue
		}

		if winner == cell {
			c++
		}
	}
	if c == N {
		return Player(winner)
	}

	if isFilled {
		return EmptyCell
	}

	return -1
}

func (board *Board) IsEnd() bool {
	winner := board.Winner()
	if winner != -1 {
		return true
	}
	return false
}

func (board *Board) PossibleActions() []*Action {
	actions := make([]*Action, 0, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board.At(i, j) == EmptyCell {
				actions = append(actions, &Action{
					X:      i,
					Y:      j,
					Player: EmptyCell,
				})
			}
		}
	}
	return actions
}

func (board *Board) Reward(player Player) float64 {
	if board.Winner() == player {
		return 1.0
	} else if board.Winner() != EmptyCell {
		return -1.0
	} else {
		return 0.0
	}
}

func (board *Board) Print() {
	state := board.Encode()
	fmt.Printf("\nPrinting board with state %d\n", state)
	fmt.Printf("+-----+\n")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == 0 {
				fmt.Printf("|")
			}
			cell := board.At(i, j)
			if cell == BlackCell {
				fmt.Printf("o")
			} else if cell == WhiteCell {
				fmt.Printf("x")
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf("|")
		}
		fmt.Printf("\n+-+-+-+\n")
	}
}
