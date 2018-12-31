package view_test

import (
	"testing"

	"github.com/miketmoore/chess/view"
	"golang.org/x/image/colornames"
)

func TestNew(t *testing.T) {
	var squares view.BoardMap
	squares, _ = view.NewBoardView(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
