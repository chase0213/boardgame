package tictactoe

import "fmt"

type TicTacToe struct {
	State State `json:"state"`
}

func NewGame() *TicTacToe {
	return &TicTacToe{
		State: InitialState(),
	}
}

func (g *TicTacToe) Start(a, b Agent) {
	nextAgent := a
	for !g.IsEnd() {
		g.Print()
		board := g.State.Decode()
		action := nextAgent.NextAction(g.State)
		fmt.Printf("action = %+v\n", action)
		if board.IsValidPlay(action) {
			board.Update(action)
			g.State = board.Encode()
		} else {
			fmt.Printf("invalid play!\n")
			continue
		}

		// update agent
		if nextAgent == a {
			nextAgent = b
		} else {
			nextAgent = a
		}
	}
	g.Print()
}

func (g *TicTacToe) IsEnd() bool {
	winner := g.Winner()
	if winner != -1 {
		return true
	}
	return false
}

func (g *TicTacToe) Winner() Player {
	board := g.State.Decode()

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

func (g *TicTacToe) Print() {
	board := g.State.Decode()
	fmt.Printf("\nPrinting board with state %d\n", g.State)
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
