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
	"github.com/miketmoore/chess/board"
	"github.com/miketmoore/chess/pieces"
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

// BoardState contains algebraic notation coordinates as keys and pieces.Piece values
// Only pieces currently on the board are set
type BoardState map[string]*pixel.Sprite

type gameState string

const (
	stateTitle gameState = "title"
	stateDraw  gameState = "draw"
	stateInput gameState = "input"
)

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
	fmt.Fprintln(displayTxt, titleStr)

	// Sub-title
	pressAnyKeyStr := T("title_pressAnyKey")
	fmt.Fprintln(bodyTxt, pressAnyKeyStr)

	// Make board
	boardThemeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	fmt.Printf("board origin x: %d\n", boardOriginX)
	squares := board.New(
		float64(boardOriginX),
		150,
		squareSize,
		board.Themes[boardThemeName]["black"],
		board.Themes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := pieces.New()

	state := stateTitle

	draw := true

	boardState := BoardState{
		"a8": drawer.Black.Rook,
		"b8": drawer.Black.Knight,
		"c8": drawer.Black.Bishop,
		"d8": drawer.Black.Queen,
		"e8": drawer.Black.King,
		"f8": drawer.Black.Bishop,
		"g8": drawer.Black.Knight,
		"h8": drawer.Black.Rook,

		"a1": drawer.White.Rook,
		"b1": drawer.White.Knight,
		"c1": drawer.White.Bishop,
		"d1": drawer.White.Queen,
		"e1": drawer.White.King,
		"f1": drawer.White.Bishop,
		"g1": drawer.White.Knight,
		"h1": drawer.White.Rook,
	}

	for _, name := range board.ColNames {
		boardState[fmt.Sprintf("%s7", name)] = drawer.Black.Pawn
	}

	for _, name := range board.ColNames {
		boardState[fmt.Sprintf("%s2", name)] = drawer.White.Pawn
	}

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		switch state {
		case stateTitle:
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

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				state = stateDraw
				win.Clear(colornames.Black)
				draw = true
			}
		case stateDraw:
			if draw {
				// Draw board
				for _, square := range squares {
					square.Shape.Draw(win)
				}

				// Draw pieces in starting positions
				for coord, piece := range boardState {
					placePiece(win, squares, piece, coord)
				}

				draw = false
				state = stateInput
			}
		case stateInput:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				fmt.Printf("mouse pos: %v\n", mpos)

				// What square was clicked?
				square := findSquareByVec(squares, mpos)
				if square != nil {
					fmt.Printf("Hit!\n")

					// Is the square occupied?
					// TODO
				}
			}

		}

		win.Update()
	}
}

func findSquareByVec(squares board.Map, vec pixel.Vec) *board.Square {
	for _, square := range squares {
		if vec.X > square.OriginX && vec.X < (square.OriginX+50) && vec.Y > square.OriginY && vec.Y < (square.OriginY+50) {
			return &square
		}
	}
	return nil
}

func placePiece(win *pixelgl.Window, squares board.Map, piece *pixel.Sprite, coord string) {
	square := squares[coord]
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
