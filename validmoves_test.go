package chess_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess"
)

func TestGetNextFile(t *testing.T) {
	tests := []struct {
		in   chess.File
		out  chess.File
		pass bool
	}{
		{chess.FileNone, chess.FileNone, false},
		{chess.FileA, chess.FileB, true},
		{chess.FileB, chess.FileC, true},
		{chess.FileC, chess.FileD, true},
		{chess.FileD, chess.FileE, true},
		{chess.FileE, chess.FileF, true},
		{chess.FileF, chess.FileG, true},
		{chess.FileG, chess.FileH, true},
		{chess.FileH, chess.FileNone, false},
	}
	for _, test := range tests {
		t.Run(string(test.in), func(t *testing.T) {
			next, ok := chess.GetNextFile(test.in)
			assertOk(t, ok == test.pass)
			assertOk(t, next == test.out)
		})

	}

}

func TestGetPreviousFile(t *testing.T) {
	tests := []struct {
		in   chess.File
		out  chess.File
		pass bool
	}{
		// {"", "", false},
		{chess.FileH, chess.FileG, true},
		{chess.FileG, chess.FileF, true},
		{chess.FileF, chess.FileE, true},
		{chess.FileE, chess.FileD, true},
		{chess.FileD, chess.FileC, true},
		{chess.FileC, chess.FileB, true},
		{chess.FileB, chess.FileA, true},
		{chess.FileA, chess.FileNone, false},
	}
	for _, test := range tests {
		t.Run(string(test.in), func(t *testing.T) {
			next, _ := chess.GetPreviousFile(test.in)
			assertOk(t, next == test.out)
		})

	}

}

func TestTranslateXYToRankFile(t *testing.T) {
	tests := []struct {
		x, y  int
		coord chess.Coord
	}{
		{0, 0, chess.NewCoord(chess.FileA, chess.Rank1)},
		{0, 1, chess.NewCoord(chess.FileA, chess.Rank2)},
		{0, 2, chess.NewCoord(chess.FileA, chess.Rank3)},
		{0, 3, chess.NewCoord(chess.FileA, chess.Rank4)},
		{0, 4, chess.NewCoord(chess.FileA, chess.Rank5)},
		{0, 5, chess.NewCoord(chess.FileA, chess.Rank6)},
		{0, 6, chess.NewCoord(chess.FileA, chess.Rank7)},
		{0, 7, chess.NewCoord(chess.FileA, chess.Rank8)},

		{7, 0, chess.NewCoord(chess.FileH, chess.Rank1)},
		{7, 1, chess.NewCoord(chess.FileH, chess.Rank2)},
		{7, 2, chess.NewCoord(chess.FileH, chess.Rank3)},
		{7, 3, chess.NewCoord(chess.FileH, chess.Rank4)},
		{7, 4, chess.NewCoord(chess.FileH, chess.Rank5)},
		{7, 5, chess.NewCoord(chess.FileH, chess.Rank6)},
		{7, 6, chess.NewCoord(chess.FileH, chess.Rank7)},
		{7, 7, chess.NewCoord(chess.FileH, chess.Rank8)},
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

func TestTranslateRankFileToXY(t *testing.T) {
	tests := []struct {
		x, y  int
		coord chess.Coord
	}{
		{0, 0, chess.NewCoord(chess.FileA, chess.Rank1)},
		{0, 1, chess.NewCoord(chess.FileA, chess.Rank2)},
		{0, 2, chess.NewCoord(chess.FileA, chess.Rank3)},
		{0, 3, chess.NewCoord(chess.FileA, chess.Rank4)},
		{0, 4, chess.NewCoord(chess.FileA, chess.Rank5)},
		{0, 5, chess.NewCoord(chess.FileA, chess.Rank6)},
		{0, 6, chess.NewCoord(chess.FileA, chess.Rank7)},
		{0, 7, chess.NewCoord(chess.FileA, chess.Rank8)},

		{7, 0, chess.NewCoord(chess.FileH, chess.Rank1)},
		{7, 1, chess.NewCoord(chess.FileH, chess.Rank2)},
		{7, 2, chess.NewCoord(chess.FileH, chess.Rank3)},
		{7, 3, chess.NewCoord(chess.FileH, chess.Rank4)},
		{7, 4, chess.NewCoord(chess.FileH, chess.Rank5)},
		{7, 5, chess.NewCoord(chess.FileH, chess.Rank6)},
		{7, 6, chess.NewCoord(chess.FileH, chess.Rank7)},
		{7, 7, chess.NewCoord(chess.FileH, chess.Rank8)},
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

func TestGetCoordsBySlope(t *testing.T) {

	// valid = append(valid, getCoordsBySlope(currCoord, 1, 1)...)
	// // valid = append(valid, getCoordsBySlope(currCoord, 1, -1)...)
	// // valid = append(valid, getCoordsBySlope(currCoord, -1, 1)...)
	// // valid = append(valid, getCoordsBySlope(currCoord, -1, -1)...)

	tests := []struct {
		file    chess.File
		rank    chess.Rank
		xChange int
		yChange int
		coords  []chess.Coord
	}{
		{chess.FileA, chess.Rank1, 1, 1, []chess.Coord{
			chess.NewCoord(chess.FileB, chess.Rank2),
			chess.NewCoord(chess.FileC, chess.Rank3),
			chess.NewCoord(chess.FileD, chess.Rank4),
			chess.NewCoord(chess.FileE, chess.Rank5),
			chess.NewCoord(chess.FileF, chess.Rank6),
			chess.NewCoord(chess.FileG, chess.Rank7),
			chess.NewCoord(chess.FileH, chess.Rank8),
		}},
		// {chess.FileH, chess.Rank1, 1, -1, []chess.Coord{
		// 	chess.NewCoord(chess.FileG, chess.Rank2),
		// 	chess.NewCoord(chess.FileF, chess.Rank3),
		// 	chess.NewCoord(chess.FileE, chess.Rank4),
		// 	chess.NewCoord(chess.FileD, chess.Rank5),
		// 	chess.NewCoord(chess.FileC, chess.Rank6),
		// 	chess.NewCoord(chess.FileB, chess.Rank7),
		// 	chess.NewCoord(chess.FileA, chess.Rank8),
		// }},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("file:%d rank:%d slope:%d,%d", test.file, test.rank, test.xChange, test.yChange), func(t *testing.T) {
			coords := chess.GetCoordsBySlope(chess.NewCoord(test.file, test.rank), test.xChange, test.yChange)
			if len(coords) != len(test.coords) {
				t.Fatal(fmt.Sprintf("length is unexpected got: %d expected: %d %v", len(coords), len(test.coords), coords))
			}
			matches := 0
			for _, expectedCoord := range test.coords {
				ok := chess.FindInSliceCoord(coords, expectedCoord)
				if ok {
					matches++
				}
				// if coord.Rank == coords[i].Rank && coord.File == coords[i].File {
				// 	fmt.Printf("%v %v\n", coord, coords[i])
				// 	matches++
				// }
			}
			if matches != len(test.coords) {
				t.Fatal(fmt.Sprintf("total matches is unexpected got: %d expected: %d %v", matches, len(test.coords), coords))
			}
		})
	}
}

func assertOk(t *testing.T, b bool) {
	if b == false {
		t.Fatal("not ok")
	}
}
