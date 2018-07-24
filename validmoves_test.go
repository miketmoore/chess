package chess_test

import (
	"testing"

	"github.com/miketmoore/chess"
)

func TestIsDestinationValid(t *testing.T) {
	got := chess.IsDestinationValid(true, true, chess.OnBoardData{
		Color: chess.PlayerBlack, Piece: chess.Pawn,
	})

	if got == false {
		t.Fatal("failed")
	}
}
