package gamestate

// GameState is the type for the state enum
type GameState uint8

const (
	Title GameState = iota
	Draw
	SelectPiece
	SelectDestination
	DrawValidMoves
)
