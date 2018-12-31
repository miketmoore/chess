package logic_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess/logic"
	"github.com/miketmoore/chess/model"
)

func TestIsDestinationValid(t *testing.T) {
	tests := []struct {
		whiteToMove bool
		isOccupied  bool
		occupant    model.PlayerPiece
		expected    bool
	}{
		{true, true, model.PlayerPiece{Color: model.PlayerBlack, Piece: model.PiecePawn}, true},
		{true, true, model.PlayerPiece{Color: model.PlayerWhite, Piece: model.PiecePawn}, false},
		{true, false, model.PlayerPiece{}, true},

		{false, true, model.PlayerPiece{Color: model.PlayerWhite, Piece: model.PiecePawn}, true},
		{false, true, model.PlayerPiece{Color: model.PlayerBlack, Piece: model.PiecePawn}, false},
		{false, false, model.PlayerPiece{}, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := logic.IsDestinationValid(
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
