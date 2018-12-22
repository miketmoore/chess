package chess_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess"
)

func TestTranslateRankFileToXY(t *testing.T) {
	tests := []struct {
		x, y  int
		coord chess.Coord
	}{
		{0, 0, chess.Coord{File:chess.FileA,Rank:chess.Rank1}},
		{0, 1, chess.Coord{File:chess.FileA,Rank:chess.Rank2}},
		{0, 2, chess.Coord{File:chess.FileA,Rank:chess.Rank3}},
		{0, 3, chess.Coord{File:chess.FileA,Rank:chess.Rank4}},
		{0, 4, chess.Coord{File:chess.FileA,Rank:chess.Rank5}},
		{0, 5, chess.Coord{File:chess.FileA,Rank:chess.Rank6}},
		{0, 6, chess.Coord{File:chess.FileA,Rank:chess.Rank7}},
		{0, 7, chess.Coord{File:chess.FileA,Rank:chess.Rank8}},

		{7, 0, chess.Coord{File:chess.FileH,Rank:chess.Rank1}},
		{7, 1, chess.Coord{File:chess.FileH,Rank:chess.Rank2}},
		{7, 2, chess.Coord{File:chess.FileH,Rank:chess.Rank3}},
		{7, 3, chess.Coord{File:chess.FileH,Rank:chess.Rank4}},
		{7, 4, chess.Coord{File:chess.FileH,Rank:chess.Rank5}},
		{7, 5, chess.Coord{File:chess.FileH,Rank:chess.Rank6}},
		{7, 6, chess.Coord{File:chess.FileH,Rank:chess.Rank7}},
		{7, 7, chess.Coord{File:chess.FileH,Rank:chess.Rank8}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.x, test.y), func(t *testing.T) {
			x, y := chess.TranslateRankFileToXY(test.coord)
			if x != test.x {
				t.Fatal("x is unexpected")
			}
			if y != test.y {
				t.Fatal("y is unexpected")
			}
		})
	}
}

func TestTranslateXYToRankFile(t *testing.T) {
	tests := []struct {
		x, y  int
		coord chess.Coord
	}{
		{0, 0, chess.Coord{File:chess.FileA,Rank:chess.Rank1}},
		{0, 1, chess.Coord{File:chess.FileA,Rank:chess.Rank2}},
		{0, 2, chess.Coord{File:chess.FileA,Rank:chess.Rank3}},
		{0, 3, chess.Coord{File:chess.FileA,Rank:chess.Rank4}},
		{0, 4, chess.Coord{File:chess.FileA,Rank:chess.Rank5}},
		{0, 5, chess.Coord{File:chess.FileA,Rank:chess.Rank6}},
		{0, 6, chess.Coord{File:chess.FileA,Rank:chess.Rank7}},
		{0, 7, chess.Coord{File:chess.FileA,Rank:chess.Rank8}},

		{7, 0, chess.Coord{File:chess.FileH,Rank:chess.Rank1}},
		{7, 1, chess.Coord{File:chess.FileH,Rank:chess.Rank2}},
		{7, 2, chess.Coord{File:chess.FileH,Rank:chess.Rank3}},
		{7, 3, chess.Coord{File:chess.FileH,Rank:chess.Rank4}},
		{7, 4, chess.Coord{File:chess.FileH,Rank:chess.Rank5}},
		{7, 5, chess.Coord{File:chess.FileH,Rank:chess.Rank6}},
		{7, 6, chess.Coord{File:chess.FileH,Rank:chess.Rank7}},
		{7, 7, chess.Coord{File:chess.FileH,Rank:chess.Rank8}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.x, test.y), func(t *testing.T) {
			coord := chess.TranslateXYToRankFile(test.x, test.y)
			if coord != test.coord {
				t.Fatal("failed")
			}
		})
	}
}
