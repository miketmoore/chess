package chess_test

import (
	"testing"

	"github.com/miketmoore/chess"
	"golang.org/x/image/colornames"
)

func TestNew(t *testing.T) {
	var squares chess.Map
	squares, _ = chess.NewBoard(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
