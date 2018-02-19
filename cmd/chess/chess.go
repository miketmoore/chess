package main

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/go-chess/board"
	"github.com/miketmoore/go-chess/pieces"
	"golang.org/x/image/colornames"
)

func run() {
	// Chess board is 8x8
	// top left is white

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  "Chess",
		Bounds: pixel.R(0, 0, 400, 400),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Darkgray)

	board := board.Build(50, colornames.Black, colornames.White)
	chessPieces := pieces.Build()

	for !win.Closed() {
		win.Update()

		// Draw board
		for _, square := range board {
			square.Draw(win)
		}

		// Draw pieces in starting positions
		drawPawns(win, chessPieces["white"]["pawn"], 25, 75)
		drawPawns(win, chessPieces["black"]["pawn"], 25, 410)

		drawRook(win, chessPieces["white"]["rook"], 25, 25)
		drawRook(win, chessPieces["white"]["rook"], 375, 25)

		drawRook(win, chessPieces["black"]["rook"], 25, 370)
		drawRook(win, chessPieces["black"]["rook"], 375, 370)
	}
}

func main() {
	pixelgl.Run(run)
}

func drawPawns(win *pixelgl.Window, piece *pixel.Sprite, x, y float64) {
	for i := 0; i < 8; i++ {
		mat := pixel.IM
		mat = mat.Moved(pixel.Vec{x, y})
		piece.Draw(win, mat)
		x += 50
	}
}

func drawRook(win *pixelgl.Window, piece *pixel.Sprite, x, y float64) {
	mat := pixel.IM
	mat = mat.Moved(pixel.Vec{x, y})
	piece.Draw(win, mat)
}
