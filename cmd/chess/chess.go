package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	ui "github.com/miketmoore/chess"
	chessapi "github.com/miketmoore/chess-api"
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
	themeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	squares, squareOriginByCoords := ui.NewBoardView(
		float64(boardOriginX),
		150,
		squareSize,
		ui.Themes[themeName]["black"],
		ui.Themes[themeName]["white"],
	)

	// Make pieces
	pieceDrawer, err := ui.NewPieceDrawer(win)
	exitOnError(err)

	// The current game data is stored here
	game := chessapi.NewGame()

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
				c := textHelper.Display.Bounds().Center()
				heightThird := screenH / 5
				c.Y = c.Y - float64(heightThird)
				textHelper.Display.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(c)))

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
				pieceDrawer.Draw(game.CurrentBoardState, squares)
				draw = false
				uiState.CurrentView = viewSelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case viewSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				square := ui.FindSquareByVec(squares, win.MousePosition())
				if square != nil {
					coord, ok := ui.GetFileRankByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						ok := game.PlyStart(coord)
						if ok {
							uiState.CurrentView = viewDrawValidMoves
							draw = true
						}
					}

				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case viewDrawValidMoves:
			if draw {
				pieceDrawer.Draw(game.CurrentBoardState, squares)
				ui.HighlightSquares(win, squares, game.ValidDestinations, colornames.Greenyellow)
				draw = false
				uiState.CurrentView = viewSelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case viewSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := ui.FindSquareByVec(squares, mpos)
				if square != nil {
					coord, ok := ui.GetFileRankByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						err, ok := game.PlyEnd(coord)
						exitOnError(err)
						if ok {
							draw = true
							uiState.CurrentView = viewDraw
						} else {
							uiState.CurrentView = viewSelectPiece
						}
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

type TextHelper struct {
	Display *text.Text
	Body    *text.Text
}

func initTextHelper() TextHelper {
	// Prepare display text
	displayFace, err := fonts.LoadTTF(displayFontPath, 80)
	exitOnError(err)

	displayAtlas := text.NewAtlas(displayFace, text.ASCII)
	displayOrig := pixel.V(screenW/2, screenH/2)

	// Prepare body text
	bodyFace, err := fonts.LoadTTF(bodyFontPath, 12)
	exitOnError(err)

	// Build body text
	bodyAtlas := text.NewAtlas(bodyFace, text.ASCII)
	bodyOrig := pixel.V(screenW/2, screenH/2)

	return TextHelper{
		Display: text.New(displayOrig, displayAtlas),
		Body:    text.New(bodyOrig, bodyAtlas),
	}
}
