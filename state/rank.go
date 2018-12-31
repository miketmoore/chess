package state

// Rank is a custom type that represents a horizontal row (rank) on the chess board
type Rank uint8

const (
	RankNone Rank = iota
	Rank1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)
