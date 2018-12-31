package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/miketmoore/chess/model"
)

func Draw(win *pixelgl.Window, boardState model.BoardState, drawer Drawer, squares BoardMap) {
	// Draw board
	for _, square := range squares {
		square.Shape.Draw(win)
	}

	// Draw pieces in the correct position
	for coord, livePieceData := range boardState {
		var set PieceSpriteSet
		if livePieceData.Color == model.PlayerBlack {
			set = drawer.Black
		} else {
			set = drawer.White
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

		DrawPiece(win, squares, piece, coord)
	}
}
