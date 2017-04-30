package game_test

import (
	"testing"
	"github.com/zicongmei/angularTest/backEnd/game"
)

func TestGame1(t *testing.T){
	g := game.NewGame()
	assertMove(t,g,game.WHITE, 0, 0)
	assertMove(t,g,game.BLACK, 1, 0)
	assertMove(t,g,game.WHITE, 2, 0)
	assertMove(t,g,game.BLACK, 0, 1)
	assertMove(t,g,game.WHITE, 1, 1)
	assertMove(t,g,game.BLACK, 2, 2)
	assertMove(t,g,game.WHITE, 2, 1)
	assertMove(t,g,game.BLACK, 1, 2)
	assertMove(t,g,game.WHITE, 0, 2)
}

func assertMove(t *testing.T, g *game.GameInfo, color, row, col int){
	_, err := g.Move(color, row, col)
	if err != nil {
		t.Fatal(err)
	}
}
