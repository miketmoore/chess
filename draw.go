package chess

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	chessapi "github.com/miketmoore/chess-api"
)

type pieceDrawer struct {
	win               *pixelgl.Window
	pieceSpriteSheets map[chessapi.PlayerColor]PieceSpriteSet
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
		pieceSpriteSheets: map[chessapi.PlayerColor]PieceSpriteSet{
			chessapi.PlayerBlack: PieceSpriteSet{
				King:   newSprite(pic, 0, 0, 40, 40),
				Queen:  newSprite(pic, 40, 0, 90, 40),
				Bishop: newSprite(pic, 90, 0, 140, 40),
				Knight: newSprite(pic, 130, 0, 180, 40),
				Rook:   newSprite(pic, 185, 0, 220, 40),
				Pawn:   newSprite(pic, 230, 0, 270, 40),
			},
			chessapi.PlayerWhite: PieceSpriteSet{
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
func (drawer pieceDrawer) Draw(boardState chessapi.BoardState, squares BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(drawer.win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range boardState {

		set := drawer.pieceSpriteSheets[livePieceData.Color]

		var piece *pixel.Sprite
		switch livePieceData.Piece {
		case chessapi.PieceBishop:
			piece = set.Bishop
		case chessapi.PieceKing:
			piece = set.King
		case chessapi.PieceKnight:
			piece = set.Knight
		case chessapi.PiecePawn:
			piece = set.Pawn
		case chessapi.PieceQueen:
			piece = set.Queen
		case chessapi.PieceRook:
			piece = set.Rook
		}

		square := squares[coord]
		x := square.OriginX + 25
		y := square.OriginY + 25
		piece.Draw(drawer.win, pixel.IM.Moved(pixel.V(x, y)))
	}
}
