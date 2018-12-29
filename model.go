package chess

// Model contains data used for the game
type Model struct {
	BoardState           BoardState
	PieceToMove          PlayerPiece
	MoveStartCoord       Coord
	MoveDestinationCoord Coord
	Draw                 bool
	WhiteToMove          bool
	CurrentState         State
}

// CurrentPlayerColor returns the current player color
func (m *Model) CurrentPlayerColor() PlayerColor {
	if m.WhiteToMove {
		return PlayerWhite
	}
	return PlayerBlack
}

// EnemyPlayerColor returns the enemy player color
func (m *Model) EnemyPlayerColor() PlayerColor {
	if m.WhiteToMove {
		return PlayerBlack
	}
	return PlayerWhite
}
