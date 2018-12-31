package model_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess/coordsmapper"
	"github.com/miketmoore/chess/model"
)

func TestTranslateRankFileToXY(t *testing.T) {
	tests := []struct {
		x, y  int
		coord model.Coord
	}{
		{0, 0, model.Coord{File: model.FileA, Rank: model.Rank1}},
		{0, 1, model.Coord{File: model.FileA, Rank: model.Rank2}},
		{0, 2, model.Coord{File: model.FileA, Rank: model.Rank3}},
		{0, 3, model.Coord{File: model.FileA, Rank: model.Rank4}},
		{0, 4, model.Coord{File: model.FileA, Rank: model.Rank5}},
		{0, 5, model.Coord{File: model.FileA, Rank: model.Rank6}},
		{0, 6, model.Coord{File: model.FileA, Rank: model.Rank7}},
		{0, 7, model.Coord{File: model.FileA, Rank: model.Rank8}},

		{7, 0, model.Coord{File: model.FileH, Rank: model.Rank1}},
		{7, 1, model.Coord{File: model.FileH, Rank: model.Rank2}},
		{7, 2, model.Coord{File: model.FileH, Rank: model.Rank3}},
		{7, 3, model.Coord{File: model.FileH, Rank: model.Rank4}},
		{7, 4, model.Coord{File: model.FileH, Rank: model.Rank5}},
		{7, 5, model.Coord{File: model.FileH, Rank: model.Rank6}},
		{7, 6, model.Coord{File: model.FileH, Rank: model.Rank7}},
		{7, 7, model.Coord{File: model.FileH, Rank: model.Rank8}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.x, test.y), func(t *testing.T) {
			x, y := coordsmapper.TranslateRankFileToXY(test.coord)
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
		coord model.Coord
	}{
		{0, 0, model.Coord{File: model.FileA, Rank: model.Rank1}},
		{0, 1, model.Coord{File: model.FileA, Rank: model.Rank2}},
		{0, 2, model.Coord{File: model.FileA, Rank: model.Rank3}},
		{0, 3, model.Coord{File: model.FileA, Rank: model.Rank4}},
		{0, 4, model.Coord{File: model.FileA, Rank: model.Rank5}},
		{0, 5, model.Coord{File: model.FileA, Rank: model.Rank6}},
		{0, 6, model.Coord{File: model.FileA, Rank: model.Rank7}},
		{0, 7, model.Coord{File: model.FileA, Rank: model.Rank8}},

		{7, 0, model.Coord{File: model.FileH, Rank: model.Rank1}},
		{7, 1, model.Coord{File: model.FileH, Rank: model.Rank2}},
		{7, 2, model.Coord{File: model.FileH, Rank: model.Rank3}},
		{7, 3, model.Coord{File: model.FileH, Rank: model.Rank4}},
		{7, 4, model.Coord{File: model.FileH, Rank: model.Rank5}},
		{7, 5, model.Coord{File: model.FileH, Rank: model.Rank6}},
		{7, 6, model.Coord{File: model.FileH, Rank: model.Rank7}},
		{7, 7, model.Coord{File: model.FileH, Rank: model.Rank8}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.x, test.y), func(t *testing.T) {
			coord := coordsmapper.TranslateXYToRankFile(test.x, test.y)
			if coord != test.coord {
				t.Fatal("failed")
			}
		})
	}
}
