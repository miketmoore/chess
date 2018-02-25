package boardview_test

import (
	"testing"

	"github.com/miketmoore/chess/boardview"
	"golang.org/x/image/colornames"
)

func TestNew(t *testing.T) {
	var squares boardview.Map
	squares = boardview.New(0, 0, 50, colornames.Black, colornames.White)

	if len(squares) != 64 {
		t.Fatal("not enough squares")
	}
}
