package board_test

import (
	"testing"

	"github.com/faiface/pixel/imdraw"
	"github.com/miketmoore/go-chess/board"
	"golang.org/x/image/colornames"
)

func TestBuild(t *testing.T) {
	var squares [64]*imdraw.IMDraw
	squares = board.Build(50, colornames.Black, colornames.White)

	for _, square := range squares {
		if square == nil {
			t.Fatal("square is nil")
		}
	}
}
