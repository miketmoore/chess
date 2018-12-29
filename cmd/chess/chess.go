package main

import (
	"fmt"
	_ "image/png"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/miketmoore/chess"
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

	// Game data
	model := chess.Model{
		BoardState:   chess.InitialOnBoardState(),
		Draw:         true,
		WhitesMove:   true,
		CurrentState: chess.StateTitle,
	}

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  i18nTitle,
		Bounds: pixel.R(0, 0, screenW, screenH),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Prepare display text
	displayFace, err := chess.LoadTTF(displayFontPath, 80)
	if err != nil {
		panic(err)
	}

	displayAtlas := text.NewAtlas(displayFace, text.ASCII)
	displayOrig := pixel.V(screenW/2, screenH/2)
	displayTxt := text.New(displayOrig, displayAtlas)

	// Prepare body text
	bodyFace, err := chess.LoadTTF(bodyFontPath, 12)
	if err != nil {
		panic(err)
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
	boardThemeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	squares, squareOriginByCoords := chess.NewBoardView(
		float64(boardOriginX),
		150,
		squareSize,
		chess.BoardThemes[boardThemeName]["black"],
		chess.BoardThemes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := chess.NewSpriteByColor()

	validDestinations := chess.ValidMoves{}

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(0)
		}

		switch model.CurrentState {
		/*
			Draw the title screen
		*/
		case chess.StateTitle:
			if model.Draw {
				win.Clear(colornames.Black)

				// Draw title text
				c := displayTxt.Bounds().Center()
				heightThird := screenH / 5
				c.Y = c.Y - float64(heightThird)
				displayTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(c)))

				// Draw secondary text
				bodyTxt.Color = colornames.White
				bodyTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(bodyTxt.Bounds().Center())))

				model.Draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				model.CurrentState = chess.StateDraw
				win.Clear(colornames.Black)
				model.Draw = true
			}
		/*
			Draw the current state of the pieces on the board
		*/
		case chess.StateDraw:
			if model.Draw {
				draw(win, model.BoardState, drawer, squares)
				model.Draw = false
				model.CurrentState = chess.StateSelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case chess.StateSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				square := chess.FindSquareByVec(squares, win.MousePosition())
				if square != nil {
					coord, ok := chess.GetCoordByXY(
						squareOriginByCoords,
						square.OriginX,
						square.OriginY,
					)
					if ok {
						occupant, isOccupied := model.BoardState[coord]
						if occupant.Color == model.CurrentPlayerColor() && isOccupied {
							validDestinations = chess.GetValidMoves(
								model.CurrentPlayerColor(),
								occupant.Piece,
								model.BoardState,
								coord,
							)
							if len(validDestinations) > 0 {
								model.PieceToMove = occupant
								model.MoveStartCoord = coord
								model.CurrentState = chess.DrawValidMoves
								model.Draw = true
							}
						}

					}

				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case chess.DrawValidMoves:
			if model.Draw {
				draw(win, model.BoardState, drawer, squares)
				chess.HighlightSquares(win, squares, validDestinations, colornames.Greenyellow)
				model.Draw = false
				model.CurrentState = chess.StateSelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case chess.StateSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := chess.FindSquareByVec(squares, mpos)
				if square != nil {
					coord, ok := chess.GetCoordByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						occupant, isOccupied := model.BoardState[coord]
						_, isValid := validDestinations[coord]
						if isValid && chess.IsDestinationValid(model.WhitesMove, isOccupied, occupant) {
							move(&model, coord)
						} else {
							model.CurrentState = chess.StateSelectPiece
						}
					}
				}
			}
		}

		win.Update()
	}
}

func move(model *chess.Model, destCoord chess.Coord) {
	model.CurrentState = chess.StateDraw
	model.Draw = true
	model.MoveDestinationCoord = destCoord

	model.BoardState[destCoord] = model.PieceToMove
	delete(model.BoardState, model.MoveStartCoord)

	model.WhitesMove = !model.WhitesMove
}

func main() {
	pixelgl.Run(run)
}

func draw(win *pixelgl.Window, boardState chess.BoardState, drawer chess.Drawer, squares chess.BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range boardState {
		var set chess.PieceSpriteSet
		if livePieceData.Color == chess.PlayerBlack {
			set = drawer.Black
		} else {
			set = drawer.White
		}

		var piece *pixel.Sprite
		switch livePieceData.Piece {
		case chess.Bishop:
			piece = set.Bishop
		case chess.King:
			piece = set.King
		case chess.Knight:
			piece = set.Knight
		case chess.Pawn:
			piece = set.Pawn
		case chess.Queen:
			piece = set.Queen
		case chess.Rook:
			piece = set.Rook
		}

		chess.DrawPiece(win, squares, piece, coord)
	}
}
