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

func TestGetNextRanks(t *testing.T) {
	tests := []struct {
		inputRank     chess.Rank
		expectedRanks []chess.Rank
		matches       int
	}{
		{chess.RankNone, []chess.Rank{}, 0},
		{chess.Rank1, []chess.Rank{chess.Rank2, chess.Rank3, chess.Rank4, chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8}, 7},
		{chess.Rank2, []chess.Rank{chess.Rank3, chess.Rank4, chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8}, 6},
		{chess.Rank3, []chess.Rank{chess.Rank4, chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8}, 5},
		{chess.Rank4, []chess.Rank{chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8}, 4},
		{chess.Rank5, []chess.Rank{chess.Rank6, chess.Rank7, chess.Rank8}, 3},
		{chess.Rank6, []chess.Rank{chess.Rank7, chess.Rank8}, 2},
		{chess.Rank7, []chess.Rank{chess.Rank8}, 1},
		{chess.Rank8, []chess.Rank{}, 0},
	}
	for _, test := range tests {
		t.Run(string(test.inputRank), func(t *testing.T) {
			got := chess.GetNextRanks(test.inputRank)
			if len(got) != len(test.expectedRanks) {
				t.Fatal(fmt.Sprintf("length is wrong - got %d expected %d", len(got), len(test.expectedRanks)))
			}
			matches := 0
			for i, rank := range got {
				if rank == test.expectedRanks[i] {
					matches++
				}
			}
			if matches != test.matches {
				t.Fatal(fmt.Sprintf("result is unexpected - got %d matches but expected %d", matches, test.matches))
			}
		})

	}
}

func TestGetRelativeCoord(t *testing.T) {
	tests := []struct {
		rank      chess.Rank
		file      chess.File
		direction chess.Direction
		distance  int
		expected  chess.Coord
	}{
		// 1 in every direction
		{chess.Rank2, chess.FileB, chess.West, 1, chess.NewCoord(chess.FileA, chess.Rank2)},
		{chess.Rank2, chess.FileB, chess.NorthWest, 1, chess.NewCoord(chess.FileA, chess.Rank3)},
		{chess.Rank2, chess.FileA, chess.North, 1, chess.NewCoord(chess.FileA, chess.Rank3)},
		{chess.Rank2, chess.FileA, chess.NorthEast, 1, chess.NewCoord(chess.FileB, chess.Rank3)},
		{chess.Rank2, chess.FileA, chess.East, 1, chess.NewCoord(chess.FileB, chess.Rank2)},
		{chess.Rank2, chess.FileA, chess.SouthEast, 1, chess.NewCoord(chess.FileB, chess.Rank1)},
		{chess.Rank3, chess.FileA, chess.South, 1, chess.NewCoord(chess.FileA, chess.Rank2)},
		{chess.Rank3, chess.FileB, chess.SouthWest, 1, chess.NewCoord(chess.FileA, chess.Rank2)},

		{chess.Rank2, chess.FileA, chess.North, 2, chess.NewCoord(chess.FileA, chess.Rank4)},
		{chess.Rank5, chess.FileA, chess.South, 2, chess.NewCoord(chess.FileA, chess.Rank3)},
		{chess.Rank2, chess.FileA, chess.East, 2, chess.NewCoord(chess.FileC, chess.Rank2)},
		{chess.Rank2, chess.FileC, chess.West, 2, chess.NewCoord(chess.FileA, chess.Rank2)},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			got, ok := chess.GetRelativeCoord(
				test.rank,
				test.file,
				test.direction,
				test.distance,
			)
			assertOk(t, ok)
			assertOk(t, got == test.expected)
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
		{chess.FileH, chess.Rank1, 1, -1, []chess.Coord{
			chess.NewCoord(chess.FileG, chess.Rank2),
			chess.NewCoord(chess.FileF, chess.Rank3),
			chess.NewCoord(chess.FileE, chess.Rank4),
			chess.NewCoord(chess.FileD, chess.Rank5),
			chess.NewCoord(chess.FileC, chess.Rank6),
			chess.NewCoord(chess.FileB, chess.Rank7),
			chess.NewCoord(chess.FileA, chess.Rank8),
		}},
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
