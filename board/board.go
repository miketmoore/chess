package board

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const totalSquares = 64
const totalRows = 8
const totalCols = 8

type boardMap map[string]*imdraw.IMDraw

// Build returns an array of *imdraw.IMDraw instances, each representing one square
// on a chess board. The size argument defines the width and height of each square.
// The blackFill and whiteFill arguments define what colors are used for the "black"
// and "white" squares.
func Build(size float64, blackFill, whiteFill color.RGBA) boardMap {
	var squareW float64 = size
	var squareH float64 = size
	var r, c float64
	var colorFlag = true
	var xInc, yInc float64
	i := 0
	squares := boardMap{}

	colNames := [totalCols]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for r = 0; r < totalRows; r++ {
		for c = 0; c < totalCols; c++ {
			square := imdraw.New(nil)
			if colorFlag {
				// dark
				square.Color = blackFill
				colorFlag = false
			} else {
				// light
				square.Color = whiteFill
				colorFlag = true
			}
			square.Push(pixel.V(xInc, yInc))
			square.Push(pixel.V(squareW+xInc, squareH+yInc))
			square.Rectangle(0)
			name := colNames[int(c)] + fmt.Sprintf("%d", int(r)+1)
			squares[name] = square
			xInc += size
			i++
		}
		colorFlag = !colorFlag
		xInc = 0
		yInc += size
	}
	return squares
}
