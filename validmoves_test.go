package chess_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess"
)

func TestIsDestinationValid(t *testing.T) {
	tests := []struct {
		whitesMove bool
		isOccupied bool
		occupant   chess.OnBoardData
		expected   bool
	}{
		{true, true, chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Pawn}, true},
		{true, true, chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Pawn}, false},
		{true, false, chess.OnBoardData{}, true},

		{false, true, chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Pawn}, true},
		{false, true, chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Pawn}, false},
		{false, false, chess.OnBoardData{}, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := chess.IsDestinationValid(
				test.whitesMove,
				test.isOccupied,
				test.occupant,
			)

			if got != test.expected {
				t.Fatal("failed")
			}
		})
	}

}
