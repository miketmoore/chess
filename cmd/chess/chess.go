package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/go-chess/board"
	"github.com/miketmoore/go-chess/pieces"
	"golang.org/x/image/colornames"
)

var boardColorSchemes = map[string]map[string]color.RGBA{
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

	// Make board
	boardThemeName := "sandcastle"
	blackFill := boardColorSchemes[boardThemeName]["black"]
	whiteFill := boardColorSchemes[boardThemeName]["white"]
	board := board.Build(50, blackFill, whiteFill)

	// Make pieces
	chessPieces := pieces.Build()

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		win.Clear(colornames.Aliceblue)

		// Draw board
		for i := 0; i < len(board); i++ {
			square := board[i]
			square.Draw(win)
		}

		// Draw pieces in starting positions
		drawPawns(win, chessPieces["white"]["pawn"], 25, 75)
		drawPawns(win, chessPieces["black"]["pawn"], 25, 410)

		drawRook(win, chessPieces["white"]["rook"], 25, 25)
		drawRook(win, chessPieces["white"]["rook"], 375, 25)

		drawRook(win, chessPieces["black"]["rook"], 25, 370)
		drawRook(win, chessPieces["black"]["rook"], 375, 370)

		win.Update()
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
