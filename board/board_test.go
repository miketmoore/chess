package board_test

import (
	"testing"

	"github.com/miketmoore/go-chess/board"
	"golang.org/x/image/colornames"
)

func TestBuild(t *testing.T) {
	squares := board.Build(50, colornames.Black, colornames.White)
	if len(squares) != 64 {
		t.Fatal("Total squares is unexpected")
	}
}
