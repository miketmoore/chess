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
	"github.com/miketmoore/chess"
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
	fmt.Fprintln(displayTxt, titleStr)

	// Sub-title
	pressAnyKeyStr := T("title_pressAnyKey")
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

	currentState := chess.StateTitle

	draw := true

	whitesMove := true

	livePieces := chess.Live{
		"a8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Rook},
		"b8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Knight},
		"c8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Bishop},
		"d8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Queen},
		"e8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.King},
		"f8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Bishop},
		"g8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Knight},
		"h8": chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Rook},

		"a1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Rook},
		"b1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Knight},
		"c1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Bishop},
		"d1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Queen},
		"e1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.King},
		"f1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Bishop},
		"g1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Knight},
		"h1": chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Rook},
	}

	for _, name := range chess.BoardColNames {
		livePieces[fmt.Sprintf("%s7", name)] = chess.LiveData{Color: chess.PlayerBlack, Piece: chess.Pawn}
	}

	for _, name := range chess.BoardColNames {
		livePieces[fmt.Sprintf("%s2", name)] = chess.LiveData{Color: chess.PlayerWhite, Piece: chess.Pawn}
	}

	var pieceToMove chess.LiveData
	var moveStartCoord string
	var moveDestinationCoord string

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		switch currentState {
		case chess.StateTitle:
			if draw {
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

				draw = false
			}

			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButtonLeft) {
				currentState = chess.StateDraw
				win.Clear(colornames.Black)
				draw = true
			}
		case chess.StateDraw:
			if draw {
				// Draw board
				for _, square := range squares {
					square.Shape.Draw(win)
				}

				// Draw pieces in the correct position
				for coord, livePieceData := range livePieces {
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

				draw = false
				currentState = chess.StateSelectPiece
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
						occupant, ok := livePieces[squareName]
						if ok {
							// TODO
							// Is this a valid piece to move?
							valid := false
							if whitesMove && occupant.Color == chess.PlayerWhite {
								// Are there valid moves for the piece?
								if occupant.Piece == chess.Pawn {
									// TODO has the pawn moved yet?
									// pawn can move one or two spaces ahead on first move
									// pawn can move one space ahead on moves after first
									// pawn can capture a piece by moving diagonal ahead, if it puts it behind an enemy piece
								}
								valid = true
							} else if !whitesMove && occupant.Color == chess.PlayerBlack {
								// Are there valid moves for the piece?
								valid = true
							}
							if valid {
								pieceToMove = occupant
								fmt.Printf("pieceToMove: %v\n", pieceToMove)
								moveStartCoord = squareName
								currentState = chess.StateSelectDestination
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
						moveDestinationCoord = squareName
						currentState = chess.StateDrawMove
						// TODO add validation
						draw = true
					}

				}
			}
		case chess.StateDrawMove:
			if draw {
				fmt.Printf("Drawing move %v from %s to %s\n", pieceToMove, moveStartCoord, moveDestinationCoord)
				draw = false
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
