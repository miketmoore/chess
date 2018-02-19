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
	pieces := pieces.Build()
	for !win.Closed() {
		win.Update()

		// Draw board
		for _, square := range board {
			square.Draw(win)
		}

		// Draw piece
		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		pieces["white"]["king"].Draw(win, mat)
	}
}

func main() {
	pixelgl.Run(run)
}
