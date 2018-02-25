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
	"github.com/miketmoore/chess/boardview"
	"github.com/miketmoore/chess/piecesdata"
	"github.com/miketmoore/chess/piecesview"
	"github.com/miketmoore/chess/player"
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

type gameState string

const (
	stateTitle             gameState = "title"
	stateDraw              gameState = "draw"
	stateSelectPiece       gameState = "selectSpace"
	stateSelectDestination gameState = "selectDestination"
	stateDrawMove          gameState = "drawMove"
)

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
	squares, squareOriginByCoords := boardview.New(
		float64(boardOriginX),
		150,
		squareSize,
		boardview.Themes[boardThemeName]["black"],
		boardview.Themes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := piecesview.New()

	state := stateTitle

	draw := true

	whitesMove := true

	livePieces := piecesdata.Live{
		"a8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Rook},
		"b8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Knight},
		"c8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Bishop},
		"d8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Queen},
		"e8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.King},
		"f8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Bishop},
		"g8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Knight},
		"h8": piecesdata.LiveData{Color: player.Black, Type: piecesdata.Rook},

		"a1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Rook},
		"b1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Knight},
		"c1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Bishop},
		"d1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Queen},
		"e1": piecesdata.LiveData{Color: player.White, Type: piecesdata.King},
		"f1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Bishop},
		"g1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Knight},
		"h1": piecesdata.LiveData{Color: player.White, Type: piecesdata.Rook},
	}

	for _, name := range boardview.ColNames {
		livePieces[fmt.Sprintf("%s7", name)] = piecesdata.LiveData{Color: player.Black, Type: piecesdata.Pawn}
	}

	for _, name := range boardview.ColNames {
		livePieces[fmt.Sprintf("%s2", name)] = piecesdata.LiveData{Color: player.White, Type: piecesdata.Pawn}
	}

	var pieceToMove piecesdata.LiveData
	var moveStartCoord string
	var moveDestinationCoord string

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		switch state {
		case stateTitle:
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
				state = stateDraw
				win.Clear(colornames.Black)
				draw = true
			}
		case stateDraw:
			if draw {
				// Draw board
				for _, square := range squares {
					square.Shape.Draw(win)
				}

				// Draw pieces in the correct position
				for coord, livePieceData := range livePieces {
					var set piecesview.Set
					if livePieceData.Color == player.Black {
						set = drawer.Black
					} else {
						set = drawer.White
					}

					var piece *pixel.Sprite
					switch livePieceData.Type {
					case piecesdata.Bishop:
						piece = set.Bishop
					case piecesdata.King:
						piece = set.King
					case piecesdata.Knight:
						piece = set.Knight
					case piecesdata.Pawn:
						piece = set.Pawn
					case piecesdata.Queen:
						piece = set.Queen
					case piecesdata.Rook:
						piece = set.Rook
					}

					placePiece(win, squares, piece, coord)
				}

				draw = false
				state = stateSelectPiece
			}
		case stateSelectPiece:
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
							if whitesMove && occupant.Color == player.White {
								// Are there valid moves for the piece?
								if occupant.Type == piecesdata.Pawn {
									// TODO has the pawn moved yet?
									// pawn can move one or two spaces ahead on first move
									// pawn can move one space ahead on moves after first
									// pawn can capture a piece by moving diagonal ahead, if it puts it behind an enemy piece
								}
								valid = true
							} else if !whitesMove && occupant.Color == player.Black {
								// Are there valid moves for the piece?
								valid = true
							}
							if valid {
								pieceToMove = occupant
								fmt.Printf("pieceToMove: %v\n", pieceToMove)
								moveStartCoord = squareName
								state = stateSelectDestination
							}
						}
					}

				}
			}
		case stateSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()

				// What square was clicked?
				square := findSquareByVec(squares, mpos)
				if square != nil {
					squareName := getSquareAlgebraicNotationByOriginCoords(squareOriginByCoords, square.OriginX, square.OriginY)
					if squareName != "" {
						fmt.Printf("moveDestinationCoord: %s\n", squareName)
						moveDestinationCoord = squareName
						state = stateDrawMove
						// TODO add validation
						draw = true
					}

				}
			}
		case stateDrawMove:
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

func findSquareByVec(squares boardview.Map, vec pixel.Vec) *boardview.Square {
	for _, square := range squares {
		if vec.X > square.OriginX && vec.X < (square.OriginX+50) && vec.Y > square.OriginY && vec.Y < (square.OriginY+50) {
			return &square
		}
	}
	return nil
}

func placePiece(win *pixelgl.Window, squares boardview.Map, piece *pixel.Sprite, coord string) {
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
