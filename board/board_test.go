package board_test

import (
	"testing"

	"github.com/miketmoore/go-chess/board"
	"golang.org/x/image/colornames"
)

func TestBuild(t *testing.T) {
	var squares board.Map
	squares = board.Build(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
