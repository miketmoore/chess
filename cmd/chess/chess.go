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
	"github.com/nicksnyder/go-i18n/i18n"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const screenW = 600
const screenH = 600
const squareSize float64 = 50
const displayFontPath = "assets/kenney_fontpackage/Fonts/Kenney Future Narrow.ttf"
const bodyFontPath = "assets/kenney_fontpackage/Fonts/Kenney Pixel Square.ttf"
const translationFile = "i18n/chess/en-US.all.json"
const lang = "en-US"

func run() {
	// i18n
	i18n.MustLoadTranslationFile(translationFile)
	T, err := i18n.Tfunc(lang)
	if err != nil {
		panic(err)
	}

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  T("title"),
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

	pressAnyKeyStr := T("title_pressAnyKey")
	// bodyTxt.Dot.X = bodyTxt.BoundsOf(pressAnyKeyStr).W() / 2
	fmt.Fprintln(bodyTxt, pressAnyKeyStr)

	// Make board
	boardThemeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	fmt.Printf("board origin x: %d\n", boardOriginX)
	board := board.New(
		float64(boardOriginX),
		150,
		squareSize,
		board.Themes[boardThemeName]["black"],
		board.Themes[boardThemeName]["white"],
	)

	// Make pieces
	chessPieces := pieces.New()

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
				c := displayTxt.Bounds().Center()
				heightThird := screenH / 5
				c.Y = c.Y - float64(heightThird)
				displayTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(c)))

				// Draw secondary text
				bodyTxt.Color = colornames.White
				bodyTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(bodyTxt.Bounds().Center())))

				draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) {
				fmt.Printf("title enter!\n")
				state = "start"
				win.Clear(colornames.Black)
				draw = true
			}
		case "start":
			if draw {
				// Draw board
				for _, square := range board {
					// fmt.Printf(">%v\n", k)
					square.Shape.Draw(win)
				}

				// Draw pieces in starting positions
				mat := pixel.IM
				mat = mat.Moved(pixel.V(223, 23))
				// chessPieces["black"]["king"].Draw(win, mat)

				// TODO "place" piece on square by algebraic notation
				// Black King - e8
				// square := board["e8"]
				// fmt.Printf("black king square: %v\n", square)
				// chessPieces["black"]["king"].Draw(win, pixel.IM.Moved(pixel.V(square.OriginX+25, square.OriginY+25)))
				placePiece(win, board, chessPieces["black"]["rook"], "a8")
				placePiece(win, board, chessPieces["black"]["knight"], "b8")
				placePiece(win, board, chessPieces["black"]["bishop"], "c8")
				placePiece(win, board, chessPieces["black"]["queen"], "d8")
				placePiece(win, board, chessPieces["black"]["king"], "e8")
				placePiece(win, board, chessPieces["black"]["bishop"], "f8")
				placePiece(win, board, chessPieces["black"]["knight"], "g8")
				placePiece(win, board, chessPieces["black"]["rook"], "h8")

				placePiece(win, board, chessPieces["black"]["pawn"], "a7")
				placePiece(win, board, chessPieces["black"]["pawn"], "b7")
				placePiece(win, board, chessPieces["black"]["pawn"], "c7")
				placePiece(win, board, chessPieces["black"]["pawn"], "d7")
				placePiece(win, board, chessPieces["black"]["pawn"], "e7")
				placePiece(win, board, chessPieces["black"]["pawn"], "f7")
				placePiece(win, board, chessPieces["black"]["pawn"], "g7")
				placePiece(win, board, chessPieces["black"]["pawn"], "h7")

				placePiece(win, board, chessPieces["white"]["rook"], "a1")
				placePiece(win, board, chessPieces["white"]["knight"], "b1")
				placePiece(win, board, chessPieces["white"]["bishop"], "c1")
				placePiece(win, board, chessPieces["white"]["queen"], "d1")
				placePiece(win, board, chessPieces["white"]["king"], "e1")
				placePiece(win, board, chessPieces["white"]["bishop"], "f1")
				placePiece(win, board, chessPieces["white"]["knight"], "g1")
				placePiece(win, board, chessPieces["white"]["rook"], "h1")

				placePiece(win, board, chessPieces["white"]["pawn"], "a2")
				placePiece(win, board, chessPieces["white"]["pawn"], "b2")
				placePiece(win, board, chessPieces["white"]["pawn"], "c2")
				placePiece(win, board, chessPieces["white"]["pawn"], "d2")
				placePiece(win, board, chessPieces["white"]["pawn"], "e2")
				placePiece(win, board, chessPieces["white"]["pawn"], "f2")
				placePiece(win, board, chessPieces["white"]["pawn"], "g2")
				placePiece(win, board, chessPieces["white"]["pawn"], "h2")

				// mat = pixel.IM
				// mat = mat.Moved(pixel.V(173, 23))
				// chessPieces["black"]["queen"].Draw(win, mat)

				// mat = pixel.IM
				// mat = mat.Moved(pixel.V(128, 23))
				// chessPieces["black"]["bishop"].Draw(win, mat)

				// mat = pixel.IM
				// mat = mat.Moved(pixel.V(278, 23))
				// chessPieces["black"]["bishop"].Draw(win, mat)

				// mat = pixel.IM
				// mat = mat.Moved(pixel.V(323, 23))
				// chessPieces["black"]["knight"].Draw(win, mat)

				// mat = pixel.IM
				// mat = mat.Moved(pixel.V(73, 23))
				// chessPieces["black"]["knight"].Draw(win, mat)
				// mat = pixel.IM
				// mat = mat.Moved(pixel.Vec{25, 23})
				// chessPieces["black"]["rook"].Draw(win, mat)

				// rook := chessPieces["black"]["rook"]
				// var rookX = center(rook.Frame().W(), 25)
				// var rookY = center(rook.Frame().H(), 25)
				// mat = pixel.IM
				// mat = mat.Moved(pixel.Vec{X: rookX, Y: rookY})
				// chessPieces["black"]["rook"].Draw(win, mat)

				// // TODO figure out width of shape
				// // fmt.Printf("%v\n", chessPieces["black"]["pawn"].Frame().W())
				// pawn := chessPieces["black"]["pawn"]
				// // var pawnXDiff float64 = (squareSize - pawn.Frame().W()) / 2
				// var pawnX = center(pawn.Frame().W(), 25)
				// var pawnY = center(pawn.Frame().H(), 75)
				// for i := 0; i < 8; i++ {
				// 	mat = pixel.IM
				// 	mat = mat.Moved(pixel.Vec{X: pawnX, Y: pawnY})
				// 	chessPieces["black"]["pawn"].Draw(win, mat)
				// 	pawnX += 50
				// }

				draw = false
			}

		}

		win.Update()
	}
}

func placePiece(win *pixelgl.Window, board board.Map, piece *pixel.Sprite, coord string) {
	square := board[coord]
	x := square.OriginX + 25
	y := square.OriginY + 25
	piece.Draw(win, pixel.IM.Moved(pixel.V(x, y)))
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
