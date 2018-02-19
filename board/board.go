package board

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const totalSquares = 64
const totalRows = 8
const totalCols = 8

// Build returns an array of *imdraw.IMDraw instances, each representing one square
// on a chess board. The size argument defines the width and height of each square.
// The blackFill and whiteFill arguments define what colors are used for the "black"
// and "white" squares.
func Build(size float64, blackFill, whiteFill color.RGBA) [totalSquares]*imdraw.IMDraw {
	var squareW float64 = size
	var squareH float64 = size
	var r, c float64
	var squares [totalSquares]*imdraw.IMDraw
	var colorFlag = true
	var xInc, yInc float64
	i := 0
	for r = 0; r < totalRows; r++ {
		for c = 0; c < totalCols; c++ {
			imd := imdraw.New(nil)
			if colorFlag {
				imd.Color = blackFill
				colorFlag = false
			} else {
				imd.Color = whiteFill
				colorFlag = true
			}
			imd.Push(pixel.V(xInc, yInc))
			imd.Push(pixel.V(squareW+xInc, squareH+yInc))
			imd.Rectangle(0)
			squares[i] = imd
			xInc += size
			i++
		}
		colorFlag = !colorFlag
		xInc = 0
		yInc += size
	}
	return squares
}
