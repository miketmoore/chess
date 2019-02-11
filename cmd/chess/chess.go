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
	chessui "github.com/miketmoore/chess"
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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Prepare display text
	displayFace, err := fonts.LoadTTF(displayFontPath, 80)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	displayAtlas := text.NewAtlas(displayFace, text.ASCII)
	displayOrig := pixel.V(screenW/2, screenH/2)
	displayTxt := text.New(displayOrig, displayAtlas)

	// Prepare body text
	bodyFace, err := fonts.LoadTTF(bodyFontPath, 12)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Build body text
	bodyAtlas := text.NewAtlas(bodyFace, text.ASCII)
	bodyOrig := pixel.V(screenW/2, screenH/2)
	bodyTxt := text.New(bodyOrig, bodyAtlas)

	// Title
	fmt.Fprintln(displayTxt, i18nTitle)

	// Sub-title
	pressAnyKeyStr := i18nPressAnyKey
	fmt.Fprintln(bodyTxt, pressAnyKeyStr)

	// Make board
	themeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	squares, squareOriginByCoords := chessui.NewBoardView(
		float64(boardOriginX),
		150,
		squareSize,
		chessui.Themes[themeName]["black"],
		chessui.Themes[themeName]["white"],
	)

	// Make pieces
	pieceDrawer, err := chessui.NewPieceDrawer(win)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// The current game data is stored here
	currentGame := chessapi.NewGame()

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

	uiState := UIState{
		CurrentView: viewTitle,
	}

	fmt.Println(currentGame.BoardState)

	doDraw := true

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(0)
		}

		switch uiState.CurrentView {

		/*
			Draw the title screen
		*/
		case viewTitle:
			if doDraw {
				fmt.Println("drawing")
				win.Clear(colornames.Black)

				// Draw title text
				c := displayTxt.Bounds().Center()
				heightThird := screenH / 5
				c.Y = c.Y - float64(heightThird)
				displayTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(c)))

				// Draw secondary text
				bodyTxt.Color = colornames.White
				bodyTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(bodyTxt.Bounds().Center())))

				doDraw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				uiState.CurrentView = viewDraw
				win.Clear(colornames.Black)
				doDraw = true
			}
		/*
			Draw the current state of the pieces on the board
		*/
		case viewDraw:
			if doDraw {
				pieceDrawer.Draw(currentGame.BoardState, squares)
				doDraw = false
				uiState.CurrentView = viewSelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case viewSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				fmt.Println("StateSelectPiece MoustButtonLeft")
				square := chessui.FindSquareByVec(squares, win.MousePosition())
				if square != nil {
					coord, ok := chessapi.GetCoordByXY(
						squareOriginByCoords,
						square.OriginX,
						square.OriginY,
					)
					if ok {
						occupant, isOccupied := currentGame.BoardState[coord]
						if occupant.Color == currentGame.CurrentPlayerColor() && isOccupied {
							currentGame.ValidDestinations = chessapi.GetValidMoves(
								currentGame.CurrentPlayerColor(),
								occupant.Piece,
								currentGame.BoardState,
								coord,
							)
							if len(currentGame.ValidDestinations) > 0 {
								currentGame.PieceToMove = occupant
								currentGame.MoveStartCoord = coord
								uiState.CurrentView = viewDrawValidMoves
								doDraw = true
							}
						}

					}

				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case viewDrawValidMoves:
			if doDraw {
				pieceDrawer.Draw(currentGame.BoardState, squares)
				chessui.HighlightSquares(win, squares, currentGame.ValidDestinations, colornames.Greenyellow)
				doDraw = false
				uiState.CurrentView = viewSelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case viewSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				fmt.Println("StateSelectDestination MouseButtonLeft")
				mpos := win.MousePosition()
				square := chessui.FindSquareByVec(squares, mpos)
				if square != nil {
					coord, ok := chessapi.GetCoordByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						fmt.Println("StateSelectDestination ok!")
						occupant, isOccupied := currentGame.BoardState[coord]
						_, isValid := currentGame.ValidDestinations[coord]
						if isValid && chessapi.IsDestinationValid(currentGame.WhiteToMove, isOccupied, occupant) {
							fmt.Println("StateSelectDestination isValid, calling Move()...")
							currentGame.Move(coord)
							fmt.Println("StateSelectDestination after Move()")
							doDraw = true
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
