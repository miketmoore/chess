package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	chessui "github.com/miketmoore/chess"
	"github.com/miketmoore/chess-api/coordsmapper"
	"github.com/miketmoore/chess-api/gamemodel"
	"github.com/miketmoore/chess-api/gamestate"
	"github.com/miketmoore/chess-api/logic"
	"github.com/miketmoore/chess-api/model"
	"github.com/miketmoore/chess-api/parse"
	"github.com/miketmoore/chess-api/save"
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

	gameLoadSuccess := false
	var currentPlayer model.PlayerColor
	var boardState model.BoardState

	if gameFilePath != "" {
		fmt.Println("Loading game from file...")

		b, err := ioutil.ReadFile(gameFilePath)
		if err != nil {
			fmt.Println("Error reading file ", err)
			os.Exit(1)
		}

		fmt.Println("Parsing game")
		currentPlayer, boardState, err = parse.Parse(string(b))
		if err != nil {
			fmt.Println("Failed to parse game ", err)
			os.Exit(1)
		}

		gameLoadSuccess = true
		fmt.Println("Finished parsing game!")
	}

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
	currentGame := gamemodel.New()

	if gameLoadSuccess == true {
		currentGame.CurrentState = gamestate.Draw
		fmt.Println("Loading game into memory...")
		currentGame.BoardState = boardState
		if currentPlayer == model.PlayerWhite {
			fmt.Println("Current player is white")
			currentGame.WhiteToMove = true
		} else {
			fmt.Println("Current player is black")
			currentGame.WhiteToMove = false
		}
	}

	fmt.Println(currentGame.BoardState)

	pendingSaveConfirm := false
	doSave := false

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(0)
		}

		if currentGame.CurrentState != gamestate.SaveGame && win.Pressed(pixelgl.KeyS) && win.Pressed(pixelgl.KeyLeftSuper) {
			fmt.Println("Save game? Y/N")
			pendingSaveConfirm = true
			currentGame.CurrentState = gamestate.SaveGame
		}

		switch currentGame.CurrentState {
		case gamestate.SaveGame:
			if pendingSaveConfirm == true {
				if win.JustPressed(pixelgl.KeyY) {
					fmt.Println("Yes")
					doSave = true
					pendingSaveConfirm = false
				} else if win.JustPressed(pixelgl.KeyN) {
					fmt.Println("No")
					doSave = false
					pendingSaveConfirm = false
				}
			} else if doSave == true {
				fmt.Println("Saving game...")
				err := save.Save(&currentGame)
				if err != nil {
					fmt.Println("Error saving game")
					os.Exit(1)
				}
				currentGame.CurrentState = gamestate.Draw
			} else {
				fmt.Println("Not saving game...")
				currentGame.CurrentState = gamestate.Draw
			}
		/*
			Draw the title screen
		*/
		case gamestate.Title:
			if currentGame.Draw {
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

				currentGame.Draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				currentGame.CurrentState = gamestate.Draw
				win.Clear(colornames.Black)
				currentGame.Draw = true
			}
		/*
			Draw the current state of the pieces on the board
		*/
		case gamestate.Draw:
			if currentGame.Draw {
				pieceDrawer.Draw(currentGame.BoardState, squares)
				currentGame.Draw = false
				currentGame.CurrentState = gamestate.SelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case gamestate.SelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				square := chessui.FindSquareByVec(squares, win.MousePosition())
				if square != nil {
					coord, ok := coordsmapper.GetCoordByXY(
						squareOriginByCoords,
						square.OriginX,
						square.OriginY,
					)
					if ok {
						occupant, isOccupied := currentGame.BoardState[coord]
						if occupant.Color == currentGame.CurrentPlayerColor() && isOccupied {
							currentGame.ValidDestinations = logic.GetValidMoves(
								currentGame.CurrentPlayerColor(),
								occupant.Piece,
								currentGame.BoardState,
								coord,
							)
							if len(currentGame.ValidDestinations) > 0 {
								currentGame.PieceToMove = occupant
								currentGame.MoveStartCoord = coord
								currentGame.CurrentState = gamestate.DrawValidMoves
								currentGame.Draw = true
							}
						}

					}

				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case gamestate.DrawValidMoves:
			if currentGame.Draw {
				pieceDrawer.Draw(currentGame.BoardState, squares)
				chessui.HighlightSquares(win, squares, currentGame.ValidDestinations, colornames.Greenyellow)
				currentGame.Draw = false
				currentGame.CurrentState = gamestate.SelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case gamestate.SelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := chessui.FindSquareByVec(squares, mpos)
				if square != nil {
					coord, ok := coordsmapper.GetCoordByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						occupant, isOccupied := currentGame.BoardState[coord]
						_, isValid := currentGame.ValidDestinations[coord]
						if isValid && logic.IsDestinationValid(currentGame.WhiteToMove, isOccupied, occupant) {
							currentGame.Move(coord)
						} else {
							currentGame.CurrentState = gamestate.SelectPiece
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
