package chess

// Model contains data used for the game
type Model struct {
	History              []HistoryEntry
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
