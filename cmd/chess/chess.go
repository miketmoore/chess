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

const screenW = 1280
const screenH = 720

// TODO make my own chess grid position tool
// Transform algebraic notation to x,y coordinates
// https://en.wikipedia.org/wiki/Algebraic_notation_(chess)
// Rows a-h
// Columns 1-8

const squareSize float64 = 50

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
		Bounds: pixel.R(0, 0, screenW, screenH),
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
	board := board.Build(squareSize, blackFill, whiteFill)

	// Make pieces
	chessPieces := pieces.Build()

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		win.Clear(colornames.Aliceblue)

		// Draw board
		for _, square := range board {
			// fmt.Printf(">%v\n", k)
			square.Draw(win)
		}

		// Draw pieces in starting positions
		mat := pixel.IM
		mat = mat.Moved(pixel.Vec{223, 23})
		chessPieces["black"]["king"].Draw(win, mat)

		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{173, 23})
		chessPieces["black"]["queen"].Draw(win, mat)

		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{128, 23})
		chessPieces["black"]["bishop"].Draw(win, mat)

		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{278, 23})
		chessPieces["black"]["bishop"].Draw(win, mat)

		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{323, 23})
		chessPieces["black"]["knight"].Draw(win, mat)

		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{73, 23})
		chessPieces["black"]["knight"].Draw(win, mat)

		// mat = pixel.IM
		// mat = mat.Moved(pixel.Vec{25, 23})
		// chessPieces["black"]["rook"].Draw(win, mat)

		rook := chessPieces["black"]["rook"]
		var rookX = center(rook.Frame().W(), 25)
		var rookY = center(rook.Frame().H(), 25)
		mat = pixel.IM
		mat = mat.Moved(pixel.Vec{rookX, rookY})
		chessPieces["black"]["rook"].Draw(win, mat)

		// TODO figure out width of shape
		// fmt.Printf("%v\n", chessPieces["black"]["pawn"].Frame().W())
		pawn := chessPieces["black"]["pawn"]
		// var pawnXDiff float64 = (squareSize - pawn.Frame().W()) / 2
		var pawnX float64 = center(pawn.Frame().W(), 25)
		var pawnY float64 = center(pawn.Frame().H(), 75)
		for i := 0; i < 8; i++ {
			mat = pixel.IM
			mat = mat.Moved(pixel.Vec{pawnX, pawnY})
			chessPieces["black"]["pawn"].Draw(win, mat)
			pawnX += 50
		}

		win.Update()
	}
}

func center(a, b float64) float64 {
	margin := (squareSize - a) / 2
	return b + margin
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
