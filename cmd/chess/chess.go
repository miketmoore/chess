package main

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/miketmoore/chess/coordsmapper"
	"github.com/miketmoore/chess/logic"
	"github.com/miketmoore/chess/state"
	"github.com/miketmoore/chess/view"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/text/language"
)

const screenW = 600
const screenH = 600
const squareSize float64 = 50
const displayFontPath = "assets/kenney_fontpackage/Fonts/Kenney Future Narrow.ttf"
const bodyFontPath = "assets/kenney_fontpackage/Fonts/Kenney Pixel Square.ttf"
const translationFile = "i18n/en.toml"
const lang = "en-US"

// State is the type for the state enum
type State string

const (
	StateTitle             State = "title"
	StateDraw              State = "draw"
	StateSelectPiece       State = "selectSpace"
	StateSelectDestination State = "selectDestination"
	DrawValidMoves         State = "drawValidMoves"
)

// Model contains data used for the game
type Model struct {
	BoardState           state.BoardState
	PieceToMove          state.PlayerPiece
	MoveStartCoord       state.Coord
	MoveDestinationCoord state.Coord
	Draw                 bool
	WhiteToMove          bool
	CurrentState         State
}

// CurrentPlayerColor returns the current player color
func (m *Model) CurrentPlayerColor() state.PlayerColor {
	if m.WhiteToMove {
		return state.PlayerWhite
	}
	return state.PlayerBlack
}

// EnemyPlayerColor returns the enemy player color
func (m *Model) EnemyPlayerColor() state.PlayerColor {
	if m.WhiteToMove {
		return state.PlayerBlack
	}
	return state.PlayerWhite
}

// LoadTTF loads a TTF font file
func LoadTTF(path string, size float64) (font.Face, error) {
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

	/*
		The current game data is stored here
	*/
	model := Model{
		BoardState:   state.InitialOnBoardState(),
		Draw:         true,
		WhiteToMove:  true,
		CurrentState: StateTitle,
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
	displayFace, err := LoadTTF(displayFontPath, 80)
	if err != nil {
		panic(err)
	}

	displayAtlas := text.NewAtlas(displayFace, text.ASCII)
	displayOrig := pixel.V(screenW/2, screenH/2)
	displayTxt := text.New(displayOrig, displayAtlas)

	// Prepare body text
	bodyFace, err := LoadTTF(bodyFontPath, 12)
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
	squares, squareOriginByCoords := view.NewBoardView(
		float64(boardOriginX),
		150,
		squareSize,
		view.BoardThemes[boardThemeName]["black"],
		view.BoardThemes[boardThemeName]["white"],
	)

	// Make pieces
	drawer := view.NewSpriteByColor()

	validDestinations := logic.ValidMoves{}

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyQ) {
			os.Exit(0)
		}

		switch model.CurrentState {
		/*
			Draw the title screen
		*/
		case StateTitle:
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
				model.CurrentState = StateDraw
				win.Clear(colornames.Black)
				model.Draw = true
			}
		/*
			Draw the current state of the pieces on the board
		*/
		case StateDraw:
			if model.Draw {
				draw(win, model.BoardState, drawer, squares)
				model.Draw = false
				model.CurrentState = StateSelectPiece
			}
		/*
			Listen for input - the current player may select a piece to move
		*/
		case StateSelectPiece:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				square := view.FindSquareByVec(squares, win.MousePosition())
				if square != nil {
					coord, ok := coordsmapper.GetCoordByXY(
						squareOriginByCoords,
						square.OriginX,
						square.OriginY,
					)
					if ok {
						occupant, isOccupied := model.BoardState[coord]
						if occupant.Color == model.CurrentPlayerColor() && isOccupied {
							validDestinations = logic.GetValidMoves(
								model.CurrentPlayerColor(),
								occupant.Piece,
								model.BoardState,
								coord,
							)
							if len(validDestinations) > 0 {
								model.PieceToMove = occupant
								model.MoveStartCoord = coord
								model.CurrentState = DrawValidMoves
								model.Draw = true
							}
						}

					}

				}
			}
		/*
			Highlight squares that are valid moves for the piece that was just selected
		*/
		case DrawValidMoves:
			if model.Draw {
				draw(win, model.BoardState, drawer, squares)
				view.HighlightSquares(win, squares, validDestinations, colornames.Greenyellow)
				model.Draw = false
				model.CurrentState = StateSelectDestination
			}
		/*
			Listen for input - the current player may select a destination square for their selected piece
		*/
		case StateSelectDestination:
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mpos := win.MousePosition()
				square := view.FindSquareByVec(squares, mpos)
				if square != nil {
					coord, ok := coordsmapper.GetCoordByXY(squareOriginByCoords, square.OriginX, square.OriginY)
					if ok {
						occupant, isOccupied := model.BoardState[coord]
						_, isValid := validDestinations[coord]
						if isValid && logic.IsDestinationValid(model.WhiteToMove, isOccupied, occupant) {
							move(&model, coord)
						} else {
							model.CurrentState = StateSelectPiece
						}
					}
				}
			}
		}

		win.Update()
	}
}

func move(model *Model, destCoord state.Coord) {
	model.CurrentState = StateDraw
	model.Draw = true
	model.MoveDestinationCoord = destCoord

	model.BoardState[destCoord] = model.PieceToMove
	delete(model.BoardState, model.MoveStartCoord)

	model.WhiteToMove = !model.WhiteToMove
}

func main() {
	pixelgl.Run(run)
}

func draw(win *pixelgl.Window, boardState state.BoardState, drawer view.Drawer, squares view.BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range boardState {
		var set view.PieceSpriteSet
		if livePieceData.Color == state.PlayerBlack {
			set = drawer.Black
		} else {
			set = drawer.White
		}

		var piece *pixel.Sprite
		switch livePieceData.Piece {
		case state.PieceBishop:
			piece = set.Bishop
		case state.PieceKing:
			piece = set.King
		case state.PieceKnight:
			piece = set.Knight
		case state.PiecePawn:
			piece = set.Pawn
		case state.PieceQueen:
			piece = set.Queen
		case state.PieceRook:
			piece = set.Rook
		}

		view.DrawPiece(win, squares, piece, coord)
	}
}
