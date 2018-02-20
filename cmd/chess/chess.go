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
	squares := board.New(
		float64(boardOriginX),
		150,
		squareSize,
		board.Themes[boardThemeName]["black"],
		board.Themes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := pieces.New()

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

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				fmt.Printf("title enter!\n")
				state = "drawGameState"
				win.Clear(colornames.Black)
				draw = true
			}
		case "drawGameState":
			if draw {
				// Draw board
				for _, square := range squares {
					// fmt.Printf(">%v\n", k)
					square.Shape.Draw(win)
				}

				// Draw pieces in starting positions
				placePiece(win, squares, drawer.Black.Rook, "a8")
				placePiece(win, squares, drawer.Black.Knight, "b8")
				placePiece(win, squares, drawer.Black.Bishop, "c8")
				placePiece(win, squares, drawer.Black.Queen, "d8")
				placePiece(win, squares, drawer.Black.King, "e8")
				placePiece(win, squares, drawer.Black.Bishop, "f8")
				placePiece(win, squares, drawer.Black.Knight, "g8")
				placePiece(win, squares, drawer.Black.Rook, "h8")

				for _, name := range board.ColNames {
					placePiece(win, squares, drawer.Black.Pawn, fmt.Sprintf("%s7", name))
				}

				placePiece(win, squares, drawer.White.Rook, "a1")
				placePiece(win, squares, drawer.White.Knight, "b1")
				placePiece(win, squares, drawer.White.Bishop, "c1")
				placePiece(win, squares, drawer.White.Queen, "d1")
				placePiece(win, squares, drawer.White.King, "e1")
				placePiece(win, squares, drawer.White.Bishop, "f1")
				placePiece(win, squares, drawer.White.Knight, "g1")
				placePiece(win, squares, drawer.White.Rook, "h1")

				for _, name := range board.ColNames {
					placePiece(win, squares, drawer.White.Pawn, fmt.Sprintf("%s2", name))
				}

				draw = false
			}

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
