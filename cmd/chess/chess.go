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

					placePiece(win, squares, piece, coord)
				}

				model.Draw = false
				model.CurrentState = chess.StateSelectPiece
			}
		case chess.StateSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()

				// What square was clicked?
				square := findSquareByVec(squares, mpos)
				if square != nil {
					squareName := getSquareAlgebraicNotationByOriginCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						// fmt.Printf("moveStartCoord: %s\n", squareName)
						// Is there a piece on this square?
						occupant, isOccupied := model.BoardState[squareName]
						if isOccupied {
							// TODO
							// Is this a valid piece to move?
							valid := false
							if model.WhitesMove && occupant.Color == chess.PlayerWhite {
								// Are there valid moves for the piece?
								if occupant.Piece == chess.Pawn {
									// TODO has the pawn moved yet?
									// pawn can move one or two spaces ahead on first move
									// pawn can move one space ahead on moves after first
									// pawn can capture a piece by moving diagonal ahead, if it puts it behind an enemy piece
									validDestinations = chess.CanPawnMove(model, squareName)
									if len(validDestinations) > 0 {
										valid = true
									}
								}
							} else if !model.WhitesMove && occupant.Color == chess.PlayerBlack {
								// Are there valid moves for the piece?
								valid = true
							}
							if valid {
								model.PieceToMove = occupant
								// fmt.Printf("pieceToMove: %v\n", model.PieceToMove)
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

				// What square was clicked?
				square := findSquareByVec(squares, mpos)
				if square != nil {
					squareName := getSquareAlgebraicNotationByOriginCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						// fmt.Printf("moveDestinationCoord: %s\n", squareName)
						_, isOccupied := model.BoardState[squareName]
						if !isOccupied {
							// is the destination in the valid destinations list?
							if inSliceString(validDestinations, squareName) {
								// fmt.Println("No occupant at destination")
								model.MoveDestinationCoord = squareName
								model.CurrentState = chess.StateDraw
								model.Draw = true

								model.History = append(model.History, chess.HistoryEntry{
									WhitesMove: model.WhitesMove,
									Piece:      model.PieceToMove.Piece,
									FromCoord:  model.MoveStartCoord,
									ToCoord:    model.MoveDestinationCoord,
								})
								model.BoardState[squareName] = model.PieceToMove
								delete(model.BoardState, model.MoveStartCoord)
								model.WhitesMove = !model.WhitesMove
								fmt.Println(model.History[len(model.History)-1])
							}

						} else {
							fmt.Println("Destination is occupied :(")
						}
					}

				}
			}
		}

		win.Update()
	}
}

func getSquareAlgebraicNotationByOriginCoords(squareOriginByCoords map[string][]float64, x, y float64) string {
	for squareName, originCoords := range squareOriginByCoords {
		if originCoords[0] == x && originCoords[1] == y {
			return squareName
		}
	}
	return ""
}

func findSquareByVec(squares chess.BoardMap, vec pixel.Vec) *chess.Square {
	for _, square := range squares {
		if vec.X > square.OriginX && vec.X < (square.OriginX+50) && vec.Y > square.OriginY && vec.Y < (square.OriginY+50) {
			return &square
		}
	}
	return nil
}

func placePiece(win *pixelgl.Window, squares chess.BoardMap, piece *pixel.Sprite, coord string) {
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

func inSliceString(slice []string, needle string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == needle {
			return true
		}
	}
	return false
}
