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
	"github.com/miketmoore/zelduh"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/image/colornames"
	"golang.org/x/text/language"
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
	bundle := &i18n.Bundle{DefaultLanguage: language.English}

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(zelduh.TranslationFile)

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
		History:      []chess.HistoryEntry{},
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
	titleStr := "Chess"
	fmt.Fprintln(displayTxt, titleStr)

	// Sub-title
	pressAnyKeyStr := i18nPressAnyKey
	fmt.Fprintln(bodyTxt, pressAnyKeyStr)

	// Make board
	boardThemeName := "sandcastle"
	boardW := squareSize * 8
	boardOriginX := (screenW - int(boardW)) / 2
	fmt.Printf("board origin x: %d\n", boardOriginX)
	squares, squareOriginByCoords := chess.NewBoardView(
		float64(boardOriginX),
		150,
		squareSize,
		chess.BoardThemes[boardThemeName]["black"],
		chess.BoardThemes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := chess.NewSpriteByColor()

	for _, name := range chess.BoardColNames {
		model.BoardState[fmt.Sprintf("%s7", name)] = chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Pawn}
	}

	for _, name := range chess.BoardColNames {
		model.BoardState[fmt.Sprintf("%s2", name)] = chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Pawn}
	}

	validDestinations := []string{}

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			printHistory(model.History)
			os.Exit(0)
		}

		switch model.CurrentState {
		case chess.StateTitle:
			if model.Draw {
				fmt.Printf("Drawing title State..\n")
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
		case chess.StateDraw:
			if model.Draw {
				// Draw board
				for _, square := range squares {
					square.Shape.Draw(win)
				}

				// Draw pieces in the correct position
				for coord, livePieceData := range model.BoardState {
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

				model.Draw = false
				model.CurrentState = chess.StateSelectPiece
			}
		case chess.StateSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := chess.FindSquareByVec(squares, mpos)
				if square != nil {
					squareName := chess.GetNotationByCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						occupant, isOccupied := model.BoardState[squareName]
						if isOccupied {
							valid := false
							if occupant.Piece == chess.Pawn {
								validDestinations = chess.CanPawnMove(
									model.CurrentPlayerColor(),
									model.BoardState,
									squareName,
								)
								if len(validDestinations) > 0 {
									valid = true
								}
							} else if occupant.Piece == chess.King {
								validDestinations = chess.CanKingMove(model, squareName)
								if len(validDestinations) > 0 {
									valid = true
								}
							} else if occupant.Piece == chess.Rook {
								validDestinations = chess.CanRookMove(model, squareName)
								if len(validDestinations) > 0 {
									valid = true
								}
							} else if occupant.Piece == chess.Knight {
								fmt.Println("knight")
								validDestinations = chess.CanKnightMove(model, squareName)
								if len(validDestinations) > 0 {
									valid = true
								}
							}
							if valid {
								model.PieceToMove = occupant
								model.MoveStartCoord = squareName
								model.CurrentState = chess.StateSelectDestination
							}
						}
					}

				}
			}
		case chess.StateSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := chess.FindSquareByVec(squares, mpos)
				if square != nil {
					squareName := chess.GetNotationByCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						occupant, isOccupied := model.BoardState[squareName]
						isValid := chess.FindInSliceString(validDestinations, squareName)
						if isValid && isDestinationValid(model.WhitesMove, squareName, isOccupied, occupant) {
							move(&model, squareName)
							printHistory(model.History)
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

func isDestinationValid(whitesMove bool, squareName string, isOccupied bool, occupant chess.OnBoardData) bool {
	if isOccupied {
		if whitesMove && occupant.Color == chess.PlayerBlack {
			return true
		} else if !whitesMove && occupant.Color == chess.PlayerWhite {
			return true
		}
	} else {
		return true
	}
	return false
}

func move(model *chess.Model, destCoord string) {
	model.CurrentState = chess.StateDraw
	model.Draw = true
	model.MoveDestinationCoord = destCoord

	entry := chess.HistoryEntry{
		WhitesMove: model.WhitesMove,
		Piece:      model.PieceToMove.Piece,
		FromCoord:  model.MoveStartCoord,
		ToCoord:    model.MoveDestinationCoord,
	}

	captor, ok := model.BoardState[destCoord]
	if ok {
		// record capture
		entry.CapturedPiece = captor.Piece
	}

	model.History = append(model.History, entry)

	model.BoardState[destCoord] = model.PieceToMove
	delete(model.BoardState, model.MoveStartCoord)
	model.WhitesMove = !model.WhitesMove
}

func main() {
	pixelgl.Run(run)
}

func printHistory(history []chess.HistoryEntry) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Player Color", "From", "To", "Captured"})
	for _, entry := range history {

		var playerColor = "white"
		if !entry.WhitesMove {
			playerColor = "black"
		}

		table.Append([]string{
			playerColor,
			entry.FromCoord,
			entry.ToCoord,
			string(entry.CapturedPiece),
		})
	}
	table.Render()
}
