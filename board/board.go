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

// Themes is a collection of color themes for the board
var Themes = map[string]map[string]color.RGBA{
	"classic": map[string]color.RGBA{
		"black": color.RGBA{0, 0, 0, 255},
		"white": color.RGBA{255, 255, 255, 255},
	},
	"coral": map[string]color.RGBA{
		"black": color.RGBA{112, 162, 163, 255},
		"white": color.RGBA{177, 228, 185, 255},
	},
	"emerald": map[string]color.RGBA{
		"black": color.RGBA{111, 143, 114, 255},
		"white": color.RGBA{173, 189, 143, 255},
	},
	"sandcastle": map[string]color.RGBA{
		"black": color.RGBA{184, 139, 74, 255},
		"white": color.RGBA{227, 193, 111, 255},
	},
}

// Map is the type used to describe the map of board squares/shapes
type Map map[string]*imdraw.IMDraw

// Build returns an array of *imdraw.IMDraw instances, each representing one square
// on a chess board. The size argument defines the width and height of each square.
// The blackFill and whiteFill arguments define what colors are used for the "black"
// and "white" squares.
func Build(originX, originY, size float64, blackFill, whiteFill color.RGBA) Map {
	var squareW = size
	var squareH = size
	var r, c float64
	var colorFlag = true
	var xInc = originX
	var yInc = originY
	i := 0
	squares := Map{}

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
		xInc = originX
		yInc += size
	}
	return squares
}
