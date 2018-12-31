package model

// Piece represents the piece type
type Piece uint8

const (
	PiecePawn Piece = iota
	PieceRook
	PieceKnight
	PieceBishop
	PieceQueen
	PieceKing
)
