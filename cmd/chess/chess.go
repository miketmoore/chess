package main

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/miketmoore/go-chess/board"
	"github.com/miketmoore/go-chess/pieces"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const screenW = 1280
const screenH = 720

// TODO make my own chess grid position tool
// Transform algebraic notation to x,y coordinates
// https://en.wikipedia.org/wiki/Algebraic_notation_(chess)
// Rows a-h
// Columns 1-8

const squareSize float64 = 50

const displayFontPath = "assets/kenney_fontpackage/Fonts/Kenney Future Narrow.ttf"
const bodyFontPath = "assets/kenney_fontpackage/Fonts/Kenney Pixel Square.ttf"

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

	// win.SetSmooth(true)

	// Prepare display text
	displayFace, err := loadTTF(displayFontPath, 80)
	if err != nil {
		panic(err)
	}

	displayAtlas := text.NewAtlas(displayFace, text.ASCII)
	displayOrig := pixel.V(screenW/2, screenH/2)
	displayTxt := text.New(displayOrig, displayAtlas)

	// Prepare body text
	bodyFace, err := loadTTF(bodyFontPath, 12)
	if err != nil {
		panic(err)
	}

	// Build body text
	bodyAtlas := text.NewAtlas(bodyFace, text.ASCII)
	bodyOrig := pixel.V(screenW/2, screenH/2)
	bodyTxt := text.New(bodyOrig, bodyAtlas)

	// Title
	titleStr := "Chess"
	// displayTxt.Dot.X = displayTxt.BoundsOf(titleStr).W() / 2
	fmt.Fprintln(displayTxt, titleStr)

	pressAnyKeyStr := "Press any key"
	// bodyTxt.Dot.X = bodyTxt.BoundsOf(pressAnyKeyStr).W() / 2
	fmt.Fprintln(bodyTxt, pressAnyKeyStr)

	// Make board
	boardThemeName := "sandcastle"
	board := board.Build(
		squareSize,
		board.Themes[boardThemeName]["black"],
		board.Themes[boardThemeName]["white"],
	)

	// Make pieces
	chessPieces := pieces.Build()

	state := "title"

	draw := true

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		switch state {
		case "title":
			if draw {
				fmt.Printf("Drawing title state...\n")
				win.Clear(colornames.Black)

				// Draw title text
				winCenter := win.Bounds().Center()
				displayTxtCenter := displayTxt.Bounds().Center()
				vec := winCenter.Sub(displayTxtCenter)
				displayTxt.Color = colornames.White
				displayTxt.Draw(win, pixel.IM.Moved(vec))

				// Draw secondary text
				bodyTxt.Color = colornames.White
				vec = win.Bounds().Center().Sub(bodyTxt.Bounds().Center())
				bodyTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(bodyTxt.Bounds().Center())))

				draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) {
				fmt.Printf("title enter!\n")
			}
		case "start":
			// Draw board
			for _, square := range board {
				// fmt.Printf(">%v\n", k)
				square.Draw(win)
			}

			// Draw pieces in starting positions
			mat := pixel.IM
			mat = mat.Moved(pixel.V(223, 23))
			chessPieces["black"]["king"].Draw(win, mat)

			mat = pixel.IM
			mat = mat.Moved(pixel.V(173, 23))
			chessPieces["black"]["queen"].Draw(win, mat)

			mat = pixel.IM
			mat = mat.Moved(pixel.V(128, 23))
			chessPieces["black"]["bishop"].Draw(win, mat)

			mat = pixel.IM
			mat = mat.Moved(pixel.V(278, 23))
			chessPieces["black"]["bishop"].Draw(win, mat)

			mat = pixel.IM
			mat = mat.Moved(pixel.V(323, 23))
			chessPieces["black"]["knight"].Draw(win, mat)

			mat = pixel.IM
			mat = mat.Moved(pixel.V(73, 23))
			chessPieces["black"]["knight"].Draw(win, mat)
			// mat = pixel.IM
			// mat = mat.Moved(pixel.Vec{25, 23})
			// chessPieces["black"]["rook"].Draw(win, mat)

			rook := chessPieces["black"]["rook"]
			var rookX = center(rook.Frame().W(), 25)
			var rookY = center(rook.Frame().H(), 25)
			mat = pixel.IM
			mat = mat.Moved(pixel.Vec{X: rookX, Y: rookY})
			chessPieces["black"]["rook"].Draw(win, mat)

			// TODO figure out width of shape
			// fmt.Printf("%v\n", chessPieces["black"]["pawn"].Frame().W())
			pawn := chessPieces["black"]["pawn"]
			// var pawnXDiff float64 = (squareSize - pawn.Frame().W()) / 2
			var pawnX = center(pawn.Frame().W(), 25)
			var pawnY = center(pawn.Frame().H(), 75)
			for i := 0; i < 8; i++ {
				mat = pixel.IM
				mat = mat.Moved(pixel.Vec{X: pawnX, Y: pawnY})
				chessPieces["black"]["pawn"].Draw(win, mat)
				pawnX += 50
			}
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
		mat = mat.Moved(pixel.V(x, y))
		piece.Draw(win, mat)
		x += 50
	}
}

func drawRook(win *pixelgl.Window, piece *pixel.Sprite, x, y float64) {
	mat := pixel.IM
	mat = mat.Moved(pixel.V(x, y))
	piece.Draw(win, mat)
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
