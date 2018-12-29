package chess

func isCoordStartPosition(playerColor PlayerColor, piece Piece, rank Rank) bool {

	if playerColor == PlayerWhite {
		// white
		if piece == Pawn {
			return rank == Rank2
		}
	} else {
		// black
		if piece == Pawn {
			return rank == Rank7
		}
	}

	return false
}

// Direction is a custom type used to describe a direction from a square on the board
type Direction int

const (
	North Direction = iota
	NorthWest
	NorthEast
	South
	SouthWest
	SouthEast
	East
	West
)

type DirectionSlope []int

var (
	SlopeNorth     DirectionSlope = []int{1, 0}
	SlopeNorthEast DirectionSlope = []int{1, 1}
	SlopeEast      DirectionSlope = []int{0, 1}
	SlopeSouthEast DirectionSlope = []int{-1, 1}
	SlopeSouth     DirectionSlope = []int{-1, 0}
	SlopeSouthWest DirectionSlope = []int{-1, -1}
	SlopeWest      DirectionSlope = []int{0, -1}
	SlopeNorthWest DirectionSlope = []int{1, -1}
)

// Rank is a custom type that represents a horizontal row (rank) on the chess board
type Rank uint8

// File is a custom type that represents a vertical column (file) on the chess board
type File uint8

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

const (
	FileNone File = iota
	FileA
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// Coord represents a board square by rank and file
type Coord struct {
	Rank Rank
	File File
}

type ValidMoves map[Coord]uint8

// GetRankFile returns the rank and file
func (c Coord) GetRankFile() (Rank, File) {
	return c.Rank, c.File
}

// GetValidMoves returns a list of valid coordinates the piece can be moved to
func GetValidMoves(playerColor PlayerColor, piece Piece, boardState BoardState, coord Coord) ValidMoves {
	switch piece {
	case Pawn:
		return getValidMovesPawn(playerColor, boardState, coord)
	case King:
		return getValidMovesKing(playerColor, boardState, coord)
	case Knight:
		return getValidMovesKnight(playerColor, boardState, coord)
	case Rook:
		return getValidMovesRook(playerColor, boardState, coord)
	case Bishop:
		return getValidMovesForBishop(playerColor, boardState, coord)
	case Queen:
		return getValidMovesForQueen(playerColor, boardState, coord)
	}
	return ValidMoves{}
}

func getValidMovesPawn(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {
	enemyColor := GetOppositeColor(playerColor)
	valid := ValidMoves{}

	yChange := 1
	if playerColor == PlayerBlack {
		yChange = -1
	}

	// get two spaces north or south
	coords := GetCoordsBySlopeAndDistance(currCoord, yChange, 0, 2)
	if !isOccupied(boardState, coords[0]) {
		valid[coords[0]] = 1
		if isCoordStartPosition(playerColor, Pawn, currCoord.Rank) && !isOccupied(boardState, coords[1]) {
			valid[coords[1]] = 1
		}
	}

	// pawn attack moves
	if playerColor == PlayerWhite {
		// NW
		coord, ok := GetCoordBySlopeAndDistance(currCoord, 1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}

		// NE
		coord, ok = GetCoordBySlopeAndDistance(currCoord, 1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	} else {
		// SW
		coord, ok := GetCoordBySlopeAndDistance(currCoord, -1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}

		// SE
		coord, ok = GetCoordBySlopeAndDistance(currCoord, -1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	}

	return valid
}

func getValidMovesKing(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {
	enemyColor := GetOppositeColor(playerColor)
	valid := ValidMoves{}

	coords := GetCoordsBySlopeAndDistanceAll(currCoord, 1)
	for _, coord := range coords {
		if !isOccupied(boardState, coord) || isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	}

	return valid
}

func getValidMovesRook(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {
	enemyColor := GetOppositeColor(playerColor)
	valid := ValidMoves{}

	slopes := []DirectionSlope{
		SlopeNorth,
		SlopeEast,
		SlopeSouth,
		SlopeWest,
	}

	for _, slope := range slopes {
		yChange := slope[0]
		xChange := slope[1]
		for i := 0; i < 8; i++ {
			coords := GetCoordsBySlopeAndDistance(currCoord, yChange, xChange, i)
			for _, coord := range coords {
				if !isOccupied(boardState, coord) {
					valid[coord] = 1
				} else if isOccupiedByColor(boardState, coord, enemyColor) {
					valid[coord] = 1
					break
				} else {
					break
				}
			}
		}
	}

	return valid
}

func getValidMovesKnight(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {

	valid := ValidMoves{}

	var slopes = [][]int{
		[]int{2, -1},
		[]int{2, 1},
		[]int{1, 2},
		[]int{-1, 2},
		[]int{-2, 1},
		[]int{-2, -1},
		[]int{1, -2},
		[]int{-1, -2},
	}

	for _, slope := range slopes {
		coords := GetCoordsBySlopeAndDistance(currCoord, slope[0], slope[1], 1)
		for _, coord := range coords {
			if !isOccupied(boardState, coord) || isOccupiedByColor(boardState, coord, GetOppositeColor(playerColor)) {
				valid[coord] = 1
			}
		}

	}

	return valid
}

func getValidMovesForBishop(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {
	enemyColor := GetOppositeColor(playerColor)
	valid := ValidMoves{}

	slopes := []DirectionSlope{
		SlopeNorthEast,
		SlopeSouthEast,
		SlopeSouthWest,
		SlopeNorthWest,
	}

	for _, slope := range slopes {
		yChange := slope[0]
		xChange := slope[1]
		for i := 0; i < 8; i++ {
			coords := GetCoordsBySlopeAndDistance(currCoord, yChange, xChange, i)
			for _, coord := range coords {
				if !isOccupied(boardState, coord) {
					valid[coord] = 1
				} else if isOccupiedByColor(boardState, coord, enemyColor) {
					valid[coord] = 1
					break
				} else {
					break
				}
			}
		}
	}

	return valid
}

func getValidMovesForQueen(playerColor PlayerColor, boardState BoardState, currCoord Coord) ValidMoves {
	diagonals := getValidMovesForBishop(playerColor, boardState, currCoord)
	horizontals := getValidMovesRook(playerColor, boardState, currCoord)
	valid := ValidMoves{}
	for coord := range diagonals {
		valid[coord] = 1
	}
	for coord := range horizontals {
		valid[coord] = 1
	}
	return valid
}

// IsDestinationValid checks if the specified color can move to the
func IsDestinationValid(whitesMove bool, isOccupied bool, occupant PlayerPiece) bool {
	if isOccupied {
		if whitesMove && occupant.Color == PlayerBlack {
			return true
		} else if !whitesMove && occupant.Color == PlayerWhite {
			return true
		}
	} else {
		return true
	}
	return false
}
