package chess

// State is the type for the state enum
type State string

const (
	StateTitle             State = "title"
	StateDraw              State = "draw"
	StateSelectPiece       State = "selectSpace"
	StateSelectDestination State = "selectDestination"
	DrawValidMoves         State = "drawValidMoves"
)
