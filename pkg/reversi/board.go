package reversi

import (
	"strconv"
)

type Board [][]*Player

func NewBoard(firstPlayer *Player) *Board {
	board := make(Board, N)
	for i := 0; i < N; i++ {
		board[i] = make([]*Player, N)
	}

	ltx := N/2 - 1
	lty := N/2 - 1
	rbx := N / 2
	rby := N / 2
	board[ltx][lty] = firstPlayer
	board[rbx][lty] = firstPlayer.Next
	board[ltx][rby] = firstPlayer.Next
	board[rbx][rby] = firstPlayer

	return &board
}

func (b *Board) At(x, y int) *Player {
	return (*b)[y][x]
}

func (b *Board) Encode() State {
	state := ""
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			p := b.At(i, j)
			if p == nil {
				state += "_"
			} else {
				state += strconv.Itoa(p.ID)
			}
		}
	}
	return State(state)
}

func (b *Board) IsEnd(players []*Player) bool {
	for _, p := range players {
		if len(b.PossibleActions(p)) != 0 {
			return false
		}
	}
	return true
}

func (b *Board) PossibleActions(player *Player) []*Action {
	actions := make([]*Action, 0, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			action := Action{
				X:      j,
				Y:      i,
				Player: player,
			}
			if b.IsPossible(&action) {
				actions = append(actions, &action)
			}
		}
	}

	return actions
}

func (b *Board) IsPossible(action *Action) bool {
	if b.At(action.X, action.Y) != nil {
		return false
	}

	// upward
	if b.checkUpward(action) {
		return true
	}

	// downward
	if b.checkDownward(action) {
		return true
	}

	// rightward
	if b.checkRightward(action) {
		return true
	}

	// leftward
	if b.checkLeftward(action) {
		return true
	}

	// up right
	if b.checkUpRight(action) {
		return true
	}

	// up left
	if b.checkUpLeft(action) {
		return true
	}

	// down right
	if b.checkDownRight(action) {
		return true
	}

	// down left
	if b.checkDownLeft(action) {
		return true
	}

	return false
}

func (b *Board) Update(action *Action) {
	// upward
	if b.checkUpward(action) {
		x := action.X
		y := action.Y
		for b.At(x, y-1) == action.Player.Next {
			(*b)[y-1][x] = action.Player
			y--
		}
	}

	// downward
	if b.checkDownward(action) {
		x := action.X
		y := action.Y
		for b.At(x, y+1) == action.Player.Next {
			(*b)[y+1][x] = action.Player
			y++
		}
	}

	// rightward
	if b.checkRightward(action) {
		x := action.X
		y := action.Y
		for b.At(x+1, y) == action.Player.Next {
			(*b)[y][x+1] = action.Player
			x++
		}
	}

	// leftward
	if b.checkLeftward(action) {
		x := action.X
		y := action.Y
		for b.At(x-1, y) == action.Player.Next {
			(*b)[y][x-1] = action.Player
			x--
		}
	}

	// upright
	if b.checkUpRight(action) {
		x := action.X
		y := action.Y
		for b.At(x+1, y-1) == action.Player.Next {
			(*b)[y-1][x+1] = action.Player
			x++
			y--
		}
	}

	// upleft
	if b.checkUpLeft(action) {
		x := action.X
		y := action.Y
		for b.At(x-1, y-1) == action.Player.Next {
			(*b)[y-1][x-1] = action.Player
			x--
			y--
		}
	}

	// downright
	if b.checkDownRight(action) {
		x := action.X
		y := action.Y
		for b.At(x+1, y+1) == action.Player.Next {
			(*b)[y+1][x+1] = action.Player
			x++
			y++
		}
	}

	// downleft
	if b.checkDownLeft(action) {
		x := action.X
		y := action.Y
		for b.At(x-1, y+1) == action.Player.Next {
			(*b)[y+1][x-1] = action.Player
			x--
			y++
		}
	}

	(*b)[action.Y][action.X] = action.Player
}

func (b *Board) checkUpward(action *Action) bool {
	_x := action.X
	_y := action.Y - 1
	for _y >= 1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x, _y-1) == action.Player {
			return true
		}
		_y = _y - 1
	}
	return false
}

func (b *Board) checkDownward(action *Action) bool {
	_x := action.X
	_y := action.Y + 1
	for _y < N-1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x, _y+1) == action.Player {
			return true
		}
		_y = _y + 1
	}
	return false
}

func (b *Board) checkRightward(action *Action) bool {
	_x := action.X + 1
	_y := action.Y
	for _x < N-1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x+1, _y) == action.Player {
			return true
		}
		_x = _x + 1
	}

	return false
}

func (b *Board) checkLeftward(action *Action) bool {
	_x := action.X - 1
	_y := action.Y
	for _x >= 1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x-1, _y) == action.Player {
			return true
		}
		_x = _x - 1
	}
	return false
}

func (b *Board) checkUpRight(action *Action) bool {
	_x := action.X + 1
	_y := action.Y - 1
	for _y >= 1 && _x < N-1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x+1, _y-1) == action.Player {
			return true
		}
		_x = _x + 1
		_y = _y - 1
	}

	return false
}

func (b *Board) checkUpLeft(action *Action) bool {
	_x := action.X - 1
	_y := action.Y - 1
	for _y >= 1 && _x >= 1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x-1, _y-1) == action.Player {
			return true
		}
		_x = _x - 1
		_y = _y - 1
	}

	return false
}

func (b *Board) checkDownRight(action *Action) bool {
	_x := action.X + 1
	_y := action.Y + 1
	for _y < N-1 && _x < N-1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x+1, _y+1) == action.Player {
			return true
		}
		_x = _x + 1
		_y = _y + 1
	}

	return false
}

func (b *Board) checkDownLeft(action *Action) bool {
	_x := action.X - 1
	_y := action.Y + 1
	for _y < N-1 && _x >= 1 && (b.At(_x, _y) == action.Player.Next) {
		if b.At(_x-1, _y+1) == action.Player {
			return true
		}
		_x = _x - 1
		_y = _y + 1
	}

	return false
}
