package state

// Type is the type for the state enum
type Type string

const (
	// Title is the title screen state
	Title Type = "title"
	// Draw is the "draw current board state" state
	Draw Type = "draw"
	// SelectPiece is the "select piece to move" state
	SelectPiece Type = "selectSpace"
	// SelectDestination is the "select destination for piece to move" state
	SelectDestination Type = "selectDestination"
	// DrawMove is the state for drawing/animating the piece to move
	DrawMove Type = "drawMove"
)
