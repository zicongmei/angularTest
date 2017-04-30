package game

import "errors"

const (
	BoardSize = 3
)

const(
	EMPTY = iota
	WHITE
	BLACK
	ONGOING
	WIN_WHITE
	WIN_BLACK
	DRAW
	GAMEERROR
)

type gameInfo struct {
	board [BoardSize][BoardSize]int
	status int
	whiteNext bool
}

func newGame() *gameInfo {
	var initGame gameInfo
	initGame.status = ONGOING
	return &initGame
}

// return if there is winner
func (g *gameInfo) move(color, row, col int) (int, error) {
	if err := g.checkError(color, row, col); err != nil {
		return GAMEERROR, err
	}
	g.board[row][col] = color
	if winner, finished := g.checkFinished(); finished {
		return winner
	} else {
		g.whiteNext = !g.whiteNext
		return EMPTY
	}
}

func (g *gameInfo) checkError(color, row, col int) (error){
	if color == WHITE && !g.whiteNext{
		return errors.New("Next is black")
	}
	if color == BLACK && g.whiteNext{
		return errors.New("Next is white")
	}
	if color != WHITE && color != BLACK {
		return errors.New("Inpt is neither black or white")
	}
	if row >= BoardSize || row < 0 || col >= BoardSize || col < 0 {
		return errors.New("out of range")
	}
	if g.board[row][col] != EMPTY {
		return errors.New("place occupied")
	}
	return nil
}

// return if finished
// return winner and if finished
func (g *gameInfo) checkFinished() (int, bool) {
	if winner, wins := g.checkWin(); wins {
		return winner + WIN_WHITE - WHITE, wins
	}
	if g.checkEmpty() {
		return EMPTY, false
	} else {
		return DRAW, true
	}
}

// check if there is empty space
func (g *gameInfo) checkEmpty() bool{
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++{
			if g.board[i][j] == EMPTY {
				return true
			}
		}
	}
	return false
}

// check if anyone wins
// return winner and if anyone wins
func (g *gameInfo) checkWin() (int, bool) {
	var target [BoardSize]int
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++{
			target[j] = g.board[i][j]
		}
		if allSame(target) {
			return target[0], true
		}
	}
	for j := 0; j < BoardSize; j++{
		for i := 0; i < BoardSize; i++ {
			target[i] = g.board[i][j]
		}
		if allSame(target) {
			return target[0], true
		}
	}
	target[0] = g.board[0][0]
	target[1] = g.board[1][1]
	target[2] = g.board[2][2]
	if allSame(target) {
		return target[0], true
	}
	target[0] = g.board[0][2]
	target[1] = g.board[1][1]
	target[2] = g.board[2][0]
	if allSame(target) {
		return target[0], true
	}
	return EMPTY, false
}

// check if an array are of same color
func allSame(target []int) bool {
	if target[0] == EMPTY {
		return false
	}
	for v := range target {
		if v != target[0] {
			return false
		}
	}
	return true
}
