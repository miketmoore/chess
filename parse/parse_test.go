package parse_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess/model"

	"github.com/miketmoore/chess/parse"
)

const input = "004140515002811811077002214840111002111880212107100261585031310740118002400251383031602171076107800230027128213861075128710721073"

func TestParse(t *testing.T) {
	color, boardState, err := parse.Parse(input)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Unexpected error caught")
	}
	if color != model.PlayerWhite {
		t.Fatal("Unexpected current player color")
	}
	if boardState == nil {
		t.Fatal("Board state is nil")
	}

	initial := model.InitialOnBoardState()

	for coord, playerPiece := range initial {
		got, ok := boardState[coord]
		if !ok {
			t.Fatal("Board state is invalid")
		}
		if playerPiece.Color != got.Color {
			t.Fatal("Color is unexpected")
		}
		if playerPiece.Piece != got.Piece {
			t.Fatal("Piece is unexpected")
		}
	}
}
