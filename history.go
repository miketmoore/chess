package chess

// HistoryEntry represents one log of game history
type HistoryEntry struct {
	WhitesMove bool
	Piece      Piece
	FromCoord  string
	ToCoord    string
}
