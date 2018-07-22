package chess

import "fmt"

// HistoryEntry represents one log of game history
type HistoryEntry struct {
	WhitesMove    bool
	Piece         Piece
	FromCoord     Coord
	ToCoord       Coord
	CapturedPiece Piece
}

// FromCoordString returns the Coord entry as a string
func (h HistoryEntry) FromCoordString() string {
	return coordToString(h.FromCoord)
}

// ToCoordString returns the Coord entry as a string
func (h HistoryEntry) ToCoordString() string {
	return coordToString(h.ToCoord)
}

func coordToString(coord Coord) string {
	rank := int(coord.Rank)
	file := int(coord.File)
	return fmt.Sprintf("%d%d", file, rank)
}
