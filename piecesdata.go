package chess

// Piece represents the piece type
type Piece string

const (
	// Pawn represents the "pawn" piece
	Pawn Piece = "pawn"
	// Rook represents the "rook" piece
	Rook Piece = "rook"
	// Knight represents the "knight" piece
	Knight Piece = "knight"
	// Bishop represents the "bishop" piece
	Bishop Piece = "bishop"
	// Queen represents the "queen" piece
	Queen Piece = "queen"
	// King represents the "king" piece
	King Piece = "king"
)

// OnBoardData represents one player's piece
type OnBoardData struct {
	Color PlayerColor
	Piece Piece
}

// OnBoard is a data structure used to track pieces currently on the board
type OnBoard map[string]OnBoardData
