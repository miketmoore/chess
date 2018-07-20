package chess

// State is the type for the state enum
type State string

const (
	// StateTitle is the title screen state
	StateTitle State = "title"
	// StateDraw is the "draw current board state" state
	StateDraw State = "draw"
	// StateSelectPiece is the "select piece to move" state
	StateSelectPiece State = "selectSpace"
	// StateSelectDestination is the "select destination for piece to move" state
	StateSelectDestination State = "selectDestination"
)
