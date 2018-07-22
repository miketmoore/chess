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
		expected  string
	}{
		// 1 in every direction
		{chess.Rank2, chess.FileB, chess.West, 1, "a2"},
		{chess.Rank2, chess.FileB, chess.NorthWest, 1, "a3"},
		{chess.Rank2, chess.FileA, chess.North, 1, "a3"},
		{chess.Rank2, chess.FileA, chess.NorthEast, 1, "b3"},
		{chess.Rank2, chess.FileA, chess.East, 1, "b2"},
		{chess.Rank2, chess.FileA, chess.SouthEast, 1, "b1"},
		{chess.Rank3, chess.FileA, chess.South, 1, "a2"},
		{chess.Rank3, chess.FileB, chess.SouthWest, 1, "a2"},

		{chess.Rank2, chess.FileA, chess.North, 2, "a4"},
		{chess.Rank5, chess.FileA, chess.South, 2, "a3"},
		{chess.Rank2, chess.FileA, chess.East, 2, "c2"},
		{chess.Rank2, chess.FileC, chess.West, 2, "a2"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			got, ok := chess.GetRelativeCoord(
				test.rank,
				test.file,
				test.direction,
				test.distance,
			)
			fmt.Println(got)
			assertOk(t, ok)
			assertOk(t, got == test.expected)
		})
	}
}

func assertOk(t *testing.T, b bool) {
	if b == false {
		t.Fatal("not ok")
	}
}
