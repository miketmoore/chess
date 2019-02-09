package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/chess-api/model"
)

type pieceDrawer struct {
	win   *pixelgl.Window
	black PieceSpriteSet
	white PieceSpriteSet
}

// NewPieceDrawer builds a pieceDrawer instance that is responsible for drawing chess pieces
func NewPieceDrawer(win *pixelgl.Window) (pieceDrawer, error) {
	// Load sprite sheet graphic
	pic, err := loadPicture(spriteSheetPath)
	if err != nil {
		return pieceDrawer{}, err
	}

	return pieceDrawer{
		win: win,
		black: PieceSpriteSet{
			King:   newSprite(pic, 0, 0, 40, 40),
			Queen:  newSprite(pic, 40, 0, 90, 40),
			Bishop: newSprite(pic, 90, 0, 140, 40),
			Knight: newSprite(pic, 130, 0, 180, 40),
			Rook:   newSprite(pic, 185, 0, 220, 40),
			Pawn:   newSprite(pic, 230, 0, 270, 40),
		},
		white: PieceSpriteSet{
			King:   newSprite(pic, 0, 40, 40, 85),
			Queen:  newSprite(pic, 40, 40, 90, 85),
			Bishop: newSprite(pic, 90, 40, 140, 85),
			Knight: newSprite(pic, 130, 40, 185, 85),
			Rook:   newSprite(pic, 185, 40, 220, 85),
			Pawn:   newSprite(pic, 230, 40, 270, 85),
		},
	}, nil
}

// Draw renders the chess pieces in the correct position on the board
func (drawer pieceDrawer) Draw(boardState model.BoardState, squares BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(drawer.win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range boardState {
		var set PieceSpriteSet
		if livePieceData.Color == model.PlayerBlack {
			set = drawer.black
		} else {
			set = drawer.white
		}

		var piece *pixel.Sprite
		switch livePieceData.Piece {
		case model.PieceBishop:
			piece = set.Bishop
		case model.PieceKing:
			piece = set.King
		case model.PieceKnight:
			piece = set.Knight
		case model.PiecePawn:
			piece = set.Pawn
		case model.PieceQueen:
			piece = set.Queen
		case model.PieceRook:
			piece = set.Rook
		}

		square := squares[coord]
		x := square.OriginX + 25
		y := square.OriginY + 25
		piece.Draw(drawer.win, pixel.IM.Moved(pixel.V(x, y)))
	}
}
