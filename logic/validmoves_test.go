package logic_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess"
)

func TestIsDestinationValid(t *testing.T) {
	tests := []struct {
		whiteToMove bool
		isOccupied  bool
		occupant    chess.PlayerPiece
		expected    bool
	}{
		{true, true, chess.PlayerPiece{Color: chess.PlayerBlack, Piece: chess.PiecePawn}, true},
		{true, true, chess.PlayerPiece{Color: chess.PlayerWhite, Piece: chess.PiecePawn}, false},
		{true, false, chess.PlayerPiece{}, true},

		{false, true, chess.PlayerPiece{Color: chess.PlayerWhite, Piece: chess.PiecePawn}, true},
		{false, true, chess.PlayerPiece{Color: chess.PlayerBlack, Piece: chess.PiecePawn}, false},
		{false, false, chess.PlayerPiece{}, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := chess.IsDestinationValid(
				test.whiteToMove,
				test.isOccupied,
				test.occupant,
			)

			if got != test.expected {
				t.Fatal("failed")
			}
		})
	}

}
