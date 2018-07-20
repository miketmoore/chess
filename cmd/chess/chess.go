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

type gameModel struct {
	OnBoard              chess.OnBoard
	pieceToMove          chess.OnBoardData
	moveStartCoord       string
	moveDestinationCoord string
	draw                 bool
	whitesMove           bool
	currentState         chess.State
}

func initialOnBoardState() chess.OnBoard {
	return chess.OnBoard{
		"a8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Rook},
		"b8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Knight},
		"c8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Bishop},
		"d8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Queen},
		"e8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.King},
		"f8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Bishop},
		"g8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Knight},
		"h8": chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Rook},

		"a1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Rook},
		"b1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Knight},
		"c1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Bishop},
		"d1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Queen},
		"e1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.King},
		"f1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Bishop},
		"g1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Knight},
		"h1": chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Rook},
	}
}

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
	model := gameModel{
		OnBoard:      initialOnBoardState(),
		draw:         true,
		whitesMove:   true,
		currentState: chess.StateTitle,
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
		model.OnBoard[fmt.Sprintf("%s7", name)] = chess.OnBoardData{Color: chess.PlayerBlack, Piece: chess.Pawn}
	}

	for _, name := range chess.BoardColNames {
		model.OnBoard[fmt.Sprintf("%s2", name)] = chess.OnBoardData{Color: chess.PlayerWhite, Piece: chess.Pawn}
	}

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		switch model.currentState {
		case chess.StateTitle:
			if model.draw {
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

				model.draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				model.currentState = chess.StateDraw
				win.Clear(colornames.Black)
				model.draw = true
			}
		case chess.StateDraw:
			if model.draw {
				// Draw board
				for _, square := range squares {
					square.Shape.Draw(win)
				}

				// Draw pieces in the correct position
				for coord, livePieceData := range model.OnBoard {
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

				model.draw = false
				model.currentState = chess.StateSelectPiece
			}
		case chess.StateSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()

				// What square was clicked?
				square := findSquareByVec(squares, mpos)
				if square != nil {
					squareName := getSquareAlgebraicNotationByOriginCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						fmt.Printf("moveStartCoord: %s\n", squareName)
						// Is there a piece on this square?
						occupant, isOccupied := model.OnBoard[squareName]
						if isOccupied {
							// TODO
							// Is this a valid piece to move?
							valid := false
							if model.whitesMove && occupant.Color == chess.PlayerWhite {
								// Are there valid moves for the piece?
								if occupant.Piece == chess.Pawn {
									// TODO has the pawn moved yet?
									// pawn can move one or two spaces ahead on first move
									// pawn can move one space ahead on moves after first
									// pawn can capture a piece by moving diagonal ahead, if it puts it behind an enemy piece
								}
								valid = true
							} else if !model.whitesMove && occupant.Color == chess.PlayerBlack {
								// Are there valid moves for the piece?
								valid = true
							}
							if valid {
								model.pieceToMove = occupant
								fmt.Printf("pieceToMove: %v\n", model.pieceToMove)
								model.moveStartCoord = squareName
								model.currentState = chess.StateSelectDestination
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
						fmt.Printf("moveDestinationCoord: %s\n", squareName)
						_, isOccupied := model.OnBoard[squareName]
						if !isOccupied {
							fmt.Println("No occupant at destination")
							model.moveDestinationCoord = squareName
							model.currentState = chess.StateDraw
							model.draw = true

							// update model.OnBoard
							model.OnBoard[squareName] = model.pieceToMove
							delete(model.OnBoard, model.moveStartCoord)
							model.whitesMove = !model.whitesMove
						} else {
							fmt.Println("Destination is occupied :(")
						}
					}

				}
			}
		case chess.StateDrawMove:
			if model.draw {
				fmt.Printf("Drawing move %v from %s to %s\n", model.pieceToMove, model.moveStartCoord, model.moveDestinationCoord)
				model.draw = false
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
