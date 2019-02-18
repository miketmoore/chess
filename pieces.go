package chess

import (
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	chessapi "github.com/miketmoore/chess-api"
)

type pieces struct {
	win               *pixelgl.Window
	pieceSpriteSheets map[chessapi.Player]pieceSpriteSheet
}

// NewPieceRenderer builds a pieces instance that is responsible for drawing chess pieces
func NewPieceRenderer(win *pixelgl.Window) (pieces, error) {
	// Load sprite sheet graphic
	pic, err := loadPicture(spriteSheetPath)
	if err != nil {
		return pieces{}, err
	}

	return pieces{
		win: win,
		pieceSpriteSheets: map[chessapi.Player]pieceSpriteSheet{
			chessapi.Black: pieceSpriteSheet{
				King:   newSprite(pic, 0, 0, 40, 40),
				Queen:  newSprite(pic, 40, 0, 90, 40),
				Bishop: newSprite(pic, 90, 0, 140, 40),
				Knight: newSprite(pic, 130, 0, 180, 40),
				Rook:   newSprite(pic, 185, 0, 220, 40),
				Pawn:   newSprite(pic, 230, 0, 270, 40),
			},
			chessapi.White: pieceSpriteSheet{
				King:   newSprite(pic, 0, 40, 40, 85),
				Queen:  newSprite(pic, 40, 40, 90, 85),
				Bishop: newSprite(pic, 90, 40, 140, 85),
				Knight: newSprite(pic, 130, 40, 185, 85),
				Rook:   newSprite(pic, 185, 40, 220, 85),
				Pawn:   newSprite(pic, 230, 40, 270, 85),
			},
		},
	}, nil
}

// Draw renders the chess pieces in the correct position on the board
func (drawer pieces) Draw(board chessapi.Board, squares BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(drawer.win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range board {

		set := drawer.pieceSpriteSheets[livePieceData.Color]

		var piece *pixel.Sprite
		switch livePieceData.Piece {
		case chessapi.Bishop:
			piece = set.Bishop
		case chessapi.King:
			piece = set.King
		case chessapi.Knight:
			piece = set.Knight
		case chessapi.Pawn:
			piece = set.Pawn
		case chessapi.Queen:
			piece = set.Queen
		case chessapi.Rook:
			piece = set.Rook
		}

		square := squares[coord]
		x := square.OriginX + 25
		y := square.OriginY + 25
		piece.Draw(drawer.win, pixel.IM.Moved(pixel.V(x, y)))
	}
}

var spriteSheetPath = "assets/standard_chess_pieces_sprite_sheet.png"

// pieceSpriteSheet contains one sprite per type of piece
type pieceSpriteSheet struct {
	King   *pixel.Sprite
	Queen  *pixel.Sprite
	Bishop *pixel.Sprite
	Knight *pixel.Sprite
	Rook   *pixel.Sprite
	Pawn   *pixel.Sprite
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{Min: pixel.V(xa, ya), Max: pixel.V(xb, yb)})
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
