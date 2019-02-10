package chess

import (
	"testing"

	"golang.org/x/image/colornames"
)

func TestNew(t *testing.T) {
	var squares BoardMap
	squares, _ = NewBoardView(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
