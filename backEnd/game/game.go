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

type GameInfo struct {
	board [BoardSize][BoardSize]int
	status int
	whiteNext bool
}

func NewGame() *GameInfo {
	var initGame GameInfo
	initGame.status = ONGOING
	initGame.whiteNext = true
	return &initGame
}

// return if there is winner
func (g *GameInfo) Move(color, row, col int) (int, error) {
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

func (g *GameInfo) checkError(color, row, col int) (error){
	if g.status != ONGOING {
		return errors.New("game finished")
	}
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
func (g *GameInfo) checkFinished() (int, bool) {
	if winner, wins := g.checkWin(); wins {
		g.status = winner + WIN_WHITE - WHITE
		return g.status, wins
	}
	if g.checkEmpty() {
		return ONGOING, false
	} else {
		g.status = DRAW
		return DRAW, true
	}
}

// check if there is empty space
func (g *GameInfo) checkEmpty() bool{
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
func (g *GameInfo) checkWin() (int, bool) {
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
