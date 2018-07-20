package piecesdata

import "github.com/miketmoore/chess"

// Type represents the piece type
type Type string

const (
	// Pawn represents the "pawn" piece
	Pawn Type = "pawn"
	// Rook represents the "rook" piece
	Rook Type = "rook"
	// Knight represents the "knight" piece
	Knight Type = "knight"
	// Bishop represents the "bishop" piece
	Bishop Type = "bishop"
	// Queen represents the "queen" piece
	Queen Type = "queen"
	// King represents the "king" piece
	King Type = "king"
)

// LiveData represents one player's piece
type LiveData struct {
	Color chess.PlayerColor
	Type  Type
}

// Live is a data structure used to track pieces currently on the board
type Live map[string]LiveData
