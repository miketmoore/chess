package board_test

import (
	"testing"

	"github.com/miketmoore/chess/board"
	"golang.org/x/image/colornames"
)

func TestNew(t *testing.T) {
	var squares board.Map
	squares = board.New(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
