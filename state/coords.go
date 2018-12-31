package state

// Coord represents a board square by rank and file
type Coord struct {
	Rank Rank
	File File
}

// GetRankFile returns the rank and file
func (c Coord) GetRankFile() (Rank, File) {
	return c.Rank, c.File
}
