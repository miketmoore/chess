package chess

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

// PlayerPiece represents one player's piece
type PlayerPiece struct {
	Color PlayerColor
	Piece Piece
}

// BoardState is a data structure used to track pieces currently on the board
type BoardState map[Coord]PlayerPiece
