package chess

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	chessapi "github.com/miketmoore/chess-api"
)

const totalSquares = 64
const totalRows = 8
const totalCols = 8

// BoardMap is the type used to describe the map of board squares/shapes
type BoardMap map[chessapi.Coord]Square

// Square represents one square of the board
type Square struct {
	Shape   *imdraw.IMDraw
	OriginX float64
	OriginY float64
}

type Board struct {
	Squares              BoardMap
	SquareOriginByCoords map[chessapi.Coord][]float64
}

func (board *Board) GetCoord(vec pixel.Vec) (chessapi.Coord, bool) {
	for _, square := range board.Squares {
		if vec.X > square.OriginX &&
			vec.X < (square.OriginX+50) &&
			vec.Y > square.OriginY &&
			vec.Y < (square.OriginY+50) {

			coord, ok := board.getFileRankByXY(square)
			if ok {
				return coord, true
			}
		}
	}
	return chessapi.Coord{}, false
}

// getFileRankByXY a coordinate for a set of rank (y) and file (x) coordinates
func (board *Board) getFileRankByXY(square Square) (chessapi.Coord, bool) {
	for coord, xy := range board.SquareOriginByCoords {
		if xy[0] == square.OriginX && xy[1] == square.OriginY {
			return coord, true
		}
	}
	return chessapi.Coord{}, false
}

// BoardColNames is the list of column names in algebraic notation
// var BoardColNames = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

// NewBoard returns an array of *imdraw.IMDraw instances, each representing one square
// on a chess boardview. The size argument defines the width and height of each square.
// The blackFill and whiteFill arguments define what colors are used for the "black"
// and "white" squares.
func NewBoard(
	originX, originY, size float64,
	blackFill, whiteFill color.RGBA,
) Board {
	var squareW = size
	var squareH = size
	var r, c float64
	var colorFlag = true
	var xInc = originX
	var yInc = originY
	i := 0
	squares := BoardMap{}

	squareOriginByCoords := map[chessapi.Coord][]float64{}

	for r = 0; r < totalRows; r++ {
		for c = 0; c < totalCols; c++ {

			shape := imdraw.New(nil)
			if colorFlag {
				// dark
				shape.Color = blackFill
				colorFlag = false
			} else {
				// light
				shape.Color = whiteFill
				colorFlag = true
			}
			shape.Push(pixel.V(xInc, yInc))
			shape.Push(pixel.V(squareW+xInc, squareH+yInc))
			shape.Rectangle(0)

			coord := chessapi.Coord{chessapi.Rank(int(r) + 1), chessapi.File(c + 1)}

			squares[coord] = Square{
				Shape:   shape,
				OriginX: xInc,
				OriginY: yInc,
			}

			squareOriginByCoords[coord] = []float64{xInc, yInc}

			xInc += size
			i++
		}
		colorFlag = !colorFlag
		xInc = originX
		yInc += size
	}

	return Board{
		Squares:              squares,
		SquareOriginByCoords: squareOriginByCoords,
	}
}

// HighlightSquares adds a visual marker to the list of board squares
func HighlightSquares(win *pixelgl.Window, squares BoardMap, coords chessapi.ValidDestinations, color color.RGBA) {
	for coord := range coords {

		square := squares[coord]
		x := square.OriginX + 13
		y := square.OriginY + 13

		shape := imdraw.New(nil)
		shape.Color = color
		shape.Push(pixel.V(x, y))
		shape.Push(pixel.V(x+25, y+25))
		shape.Rectangle(0)
		shape.Draw(win)
	}
}

// FindSquareByVec finds a square from the board map by it's (x,y) coordinate
func FindSquareByVec(squares BoardMap, vec pixel.Vec) *Square {
	for _, square := range squares {
		if vec.X > square.OriginX &&
			vec.X < (square.OriginX+50) &&
			vec.Y > square.OriginY &&
			vec.Y < (square.OriginY+50) {
			return &square
		}
	}
	return nil
}
