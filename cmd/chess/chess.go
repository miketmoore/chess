package main

import (
	"flag"
	"fmt"
	"image/color"
	_ "image/png"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	ui "github.com/miketmoore/chess"
	api "github.com/miketmoore/chess-api"
	"github.com/miketmoore/chess/fonts"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/image/colornames"
	"golang.org/x/text/language"
)

const screenW = 600
const screenH = 600
const squareSize float64 = 50
const displayFontPath = "assets/kenney_fontpackage/Fonts/Kenney Future Narrow.ttf"
const bodyFontPath = "assets/kenney_fontpackage/Fonts/Kenney Pixel Square.ttf"
const translationFile = "i18n/en.toml"
const lang = "en-US"

type view int

const (
	viewTitle view = iota
	viewDraw
	viewSelectPiece
	viewDrawValidMoves
	viewSelectDestination
)

type UIState struct {
	CurrentView view
}

func run() {

	var gameFilePath string

	flag.StringVar(&gameFilePath, "game", "", "file path of game to load")
	flag.Parse()

	// i18n
	bundle := &i18n.Bundle{DefaultLanguage: language.English}

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(translationFile)

	localizer := i18n.NewLocalizer(bundle, "en")

	i18nTitle := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "Title",
		},
	})
	i18nPressAnyKey := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "PressAnyKey",
		},
	})

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  i18nTitle,
		Bounds: pixel.R(0, 0, screenW, screenH),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	exitOnError(err)

	textHelper := initTextHelper()

	// Title
	fmt.Fprintln(textHelper.Display, i18nTitle)

	// Sub-title
	pressAnyKeyStr := i18nPressAnyKey
	fmt.Fprintln(textHelper.Body, pressAnyKeyStr)

	// Make board
	boardW := squareSize * 8
	var boardOriginX float64 = (screenW - boardW) / 2
	var boardOriginY float64 = 150
	blackSquareColor := color.RGBA{184, 139, 74, 255}
	whiteSquareColor := color.RGBA{227, 193, 111, 255}
	board := ui.NewBoard(
		boardOriginX,
		boardOriginY,
		squareSize,
		blackSquareColor,
		whiteSquareColor,
	)

	// Make pieces
	pieces, err := ui.NewPieceRenderer(win)
	exitOnError(err)

	// The current game data is stored here
	game := api.NewGame()

	uiState := UIState{
		CurrentView: viewTitle,
	}

	draw := true

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			exit()
		}

		switch uiState.CurrentView {

		/*
			Draw the title screen
		*/
		case viewTitle:
			if draw {
				win.Clear(colornames.Black)

				// Draw title text
				center := textHelper.Display.Bounds().Center()
				heightThird := screenH / 5
				center.Y = center.Y - float64(heightThird)
				textHelper.Display.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(center)))

				// Draw secondary text
				textHelper.Body.Color = colornames.White
				textHelper.Body.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(textHelper.Body.Bounds().Center())))

				draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				uiState.CurrentView = viewDraw
				win.Clear(colornames.Black)
				draw = true
			}
		/*
			Draw the current state of the pieces on the board
		*/
		case viewDraw:
			if draw {
				pieces.Draw(game.Board, board.Squares)
				draw = false
				uiState.CurrentView = viewSelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case viewSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				coord, coordOK := board.GetCoord(win.MousePosition())
				if coordOK {
					plyStartOK := game.PlyStart(coord)
					if plyStartOK {
						uiState.CurrentView = viewDrawValidMoves
						draw = true
					}
				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case viewDrawValidMoves:
			if draw {
				pieces.Draw(game.Board, board.Squares)
				ui.HighlightSquares(win, board.Squares, game.ValidDestinations, colornames.Greenyellow)
				draw = false
				uiState.CurrentView = viewSelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case viewSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				coord, coordOK := board.GetCoord(win.MousePosition())
				if coordOK {
					err, plyEndOK, capture, capturedPiece := game.PlyEnd(coord)
					exitOnError(err)
					if plyEndOK {
						if capture {
							fmt.Printf("Captured %s %s!\n", capturedPiece.Color, capturedPiece.Piece)
						}
						draw = true
						uiState.CurrentView = viewDraw
					} else {
						uiState.CurrentView = viewSelectPiece
					}
				}
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func exit() {
	os.Exit(0)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type textHelper struct {
	Display *text.Text
	Body    *text.Text
}

func initTextHelper() textHelper {
	displayFace, err := fonts.LoadTTF(displayFontPath, 80)
	exitOnError(err)

	bodyFace, err := fonts.LoadTTF(bodyFontPath, 12)
	exitOnError(err)

	return textHelper{
		Display: text.New(
			pixel.V(screenW/2, screenH/2),
			text.NewAtlas(displayFace, text.ASCII),
		),
		Body: text.New(
			pixel.V(screenW/2, screenH/2),
			text.NewAtlas(bodyFace, text.ASCII),
		),
	}
}
