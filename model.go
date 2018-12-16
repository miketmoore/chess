package chess

// Model contains data used for the game
type Model struct {
	BoardState           BoardState
	PieceToMove          OnBoardData
	MoveStartCoord       Coord
	MoveDestinationCoord Coord
	Draw                 bool
	WhitesMove           bool
	CurrentState         State
}

// CurrentPlayerColor returns the current player color
func (m *Model) CurrentPlayerColor() PlayerColor {
	if m.WhitesMove {
		return PlayerWhite
	}
	return PlayerBlack
}

// EnemyPlayerColor returns the enemy player color
func (m *Model) EnemyPlayerColor() PlayerColor {
	if m.WhitesMove {
		return PlayerBlack
	}
	return PlayerWhite
}
