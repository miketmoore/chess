package chess_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/chess"
)

func TestGetNextFile(t *testing.T) {
	tests := []struct {
		in   string
		out  string
		pass bool
	}{
		{"invalidinput", "", false},
		{"a", "b", true},
		{"b", "c", true},
		{"c", "d", true},
		{"d", "e", true},
		{"e", "f", true},
		{"f", "g", true},
		{"g", "h", true},
		{"h", "", false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			next, ok := chess.GetNextFile(test.in)
			assertOk(t, ok == test.pass)
			assertOk(t, next == test.out)
		})

	}

}

func TestGetPreviousFile(t *testing.T) {
	tests := []struct {
		in   string
		out  string
		pass bool
	}{
		{"", "", false},
		{"h", "g", true},
		{"g", "f", true},
		{"f", "e", true},
		{"e", "d", true},
		{"d", "c", true},
		{"c", "b", true},
		{"b", "a", true},
		{"a", "", false},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			next, _ := chess.GetPreviousFile(test.in)
			assertOk(t, next == test.out)
		})

	}

}

func TestGetNextRanks(t *testing.T) {
	tests := []struct {
		inputRank     string
		expectedRanks []string
		matches       int
	}{
		{"1", []string{"2", "3", "4", "5", "6", "7", "8"}, 7},
		{"2", []string{"3", "4", "5", "6", "7", "8"}, 6},
		{"3", []string{"4", "5", "6", "7", "8"}, 5},
		{"4", []string{"5", "6", "7", "8"}, 4},
		{"5", []string{"6", "7", "8"}, 3},
		{"6", []string{"7", "8"}, 2},
		{"7", []string{"8"}, 1},
	}
	for _, test := range tests {
		t.Run(test.inputRank, func(t *testing.T) {
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

func assertOk(t *testing.T, b bool) {
	if b == false {
		t.Fatal("not ok")
	}
}
