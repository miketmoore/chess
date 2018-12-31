package model

// PlayerColor is a type for player colors
type PlayerColor bool

const (
	PlayerWhite PlayerColor = true
	PlayerBlack PlayerColor = false
)

// PlayerPiece represents one player's piece
type PlayerPiece struct {
	Color PlayerColor
	Piece Piece
}
