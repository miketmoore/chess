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
type Rank int

// File is a custom type that represents a vertical column (file) on the chess board
type File int

const (
	RankNone Rank = 0
	Rank1    Rank = 1
	Rank2    Rank = 2
	Rank3    Rank = 3
	Rank4    Rank = 4
	Rank5    Rank = 5
	Rank6    Rank = 6
	Rank7    Rank = 7
	Rank8    Rank = 8

	FileNone File = 0
	FileA    File = 1
	FileB    File = 2
	FileC    File = 3
	FileD    File = 4
	FileE    File = 5
	FileF    File = 6
	FileG    File = 7
	FileH    File = 8
)

var ranksOrder = []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}
var FilesOrder = []File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

var rankByRankView = map[Rank]string{
	Rank1: "1",
	Rank2: "2",
	Rank3: "3",
	Rank4: "4",
	Rank5: "5",
	Rank6: "6",
	Rank7: "7",
	Rank8: "8",
}

var fileByFileView = map[File]string{
	FileA: "a",
	FileB: "b",
	FileC: "c",
	FileD: "d",
	FileE: "e",
	FileF: "f",
	FileG: "g",
	FileH: "h",
}

var rankViewByRank = map[string]Rank{
	"1": Rank1,
	"2": Rank2,
	"3": Rank3,
	"4": Rank4,
	"5": Rank5,
	"6": Rank6,
	"7": Rank7,
	"8": Rank8,
}

var fileViewByFile = map[string]File{
	"a": FileA,
	"b": FileB,
	"c": FileC,
	"d": FileD,
	"e": FileE,
	"f": FileF,
	"g": FileG,
	"h": FileH,
}

// Coord represents a board square by rank and file
type Coord struct {
	Rank Rank
	File File
}

// GetRankFile returns the rank and file
func (c Coord) GetRankFile() (Rank, File) {
	return c.Rank, c.File
}

// NewCoord returns a new rank and file coordinate
func NewCoord(file File, rank Rank) Coord {
	return Coord{
		File: file,
		Rank: rank,
	}
}

var validCoords = map[Coord]bool{
	Coord{Rank1, FileA}: true,
	Coord{Rank2, FileA}: true,
	Coord{Rank3, FileA}: true,
	Coord{Rank4, FileA}: true,
	Coord{Rank5, FileA}: true,
	Coord{Rank6, FileA}: true,
	Coord{Rank7, FileA}: true,
	Coord{Rank8, FileA}: true,

	Coord{Rank1, FileB}: true,
	Coord{Rank2, FileB}: true,
	Coord{Rank3, FileB}: true,
	Coord{Rank4, FileB}: true,
	Coord{Rank5, FileB}: true,
	Coord{Rank6, FileB}: true,
	Coord{Rank7, FileB}: true,
	Coord{Rank8, FileB}: true,

	Coord{Rank1, FileC}: true,
	Coord{Rank2, FileC}: true,
	Coord{Rank3, FileC}: true,
	Coord{Rank4, FileC}: true,
	Coord{Rank5, FileC}: true,
	Coord{Rank6, FileC}: true,
	Coord{Rank7, FileC}: true,
	Coord{Rank8, FileC}: true,

	Coord{Rank1, FileD}: true,
	Coord{Rank2, FileD}: true,
	Coord{Rank3, FileD}: true,
	Coord{Rank4, FileD}: true,
	Coord{Rank5, FileD}: true,
	Coord{Rank6, FileD}: true,
	Coord{Rank7, FileD}: true,
	Coord{Rank8, FileD}: true,

	Coord{Rank1, FileE}: true,
	Coord{Rank2, FileE}: true,
	Coord{Rank3, FileE}: true,
	Coord{Rank4, FileE}: true,
	Coord{Rank5, FileE}: true,
	Coord{Rank6, FileE}: true,
	Coord{Rank7, FileE}: true,
	Coord{Rank8, FileE}: true,

	Coord{Rank1, FileF}: true,
	Coord{Rank2, FileF}: true,
	Coord{Rank3, FileF}: true,
	Coord{Rank4, FileF}: true,
	Coord{Rank5, FileF}: true,
	Coord{Rank6, FileF}: true,
	Coord{Rank7, FileF}: true,
	Coord{Rank8, FileF}: true,

	Coord{Rank1, FileG}: true,
	Coord{Rank2, FileG}: true,
	Coord{Rank3, FileG}: true,
	Coord{Rank4, FileG}: true,
	Coord{Rank5, FileG}: true,
	Coord{Rank6, FileG}: true,
	Coord{Rank7, FileG}: true,
	Coord{Rank8, FileG}: true,

	Coord{Rank1, FileH}: true,
	Coord{Rank2, FileH}: true,
	Coord{Rank3, FileH}: true,
	Coord{Rank4, FileH}: true,
	Coord{Rank5, FileH}: true,
	Coord{Rank6, FileH}: true,
	Coord{Rank7, FileH}: true,
	Coord{Rank8, FileH}: true,
}

// GetValidMoves returns a list of valid coordinates the piece can be moved to
func GetValidMoves(playerColor PlayerColor, piece Piece, boardState BoardState, coord Coord) []Coord {
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
	return []Coord{}
}

func getValidMovesPawn(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	enemyColor := GetOppositeColor(playerColor)
	valid := []Coord{}

	yChange := 1
	if playerColor == PlayerBlack {
		yChange = -1
	}

	// get two spaces north or south
	coords := GetCoordsBySlopeAndDistance(currCoord, yChange, 0, 2)
	if !isOccupied(boardState, coords[0]) {
		valid = append(valid, coords[0])
		if isCoordStartPosition(playerColor, Pawn, currCoord.Rank) && !isOccupied(boardState, coords[1]) {
			valid = append(valid, coords[1])
		}
	}

	// pawn attack moves
	if playerColor == PlayerWhite {
		// NW
		coord, ok := GetCoordBySlopeAndDistance(currCoord, 1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid = append(valid, coord)
		}

		// NE
		coord, ok = GetCoordBySlopeAndDistance(currCoord, 1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid = append(valid, coord)
		}
	} else {
		// SW
		coord, ok := GetCoordBySlopeAndDistance(currCoord, -1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid = append(valid, coord)
		}

		// SE
		coord, ok = GetCoordBySlopeAndDistance(currCoord, -1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid = append(valid, coord)
		}
	}

	return valid
}

func getValidMovesKing(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	enemyColor := GetOppositeColor(playerColor)
	valid := []Coord{}

	coords := GetCoordsBySlopeAndDistanceAll(currCoord, 1)
	for _, coord := range coords {
		if !isOccupied(boardState, coord) || isOccupiedByColor(boardState, coord, enemyColor) {
			valid = append(valid, coord)
		}
	}

	return valid
}

func getValidMovesRook(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	enemyColor := GetOppositeColor(playerColor)
	valid := []Coord{}

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
					valid = append(valid, coord)
				} else if isOccupiedByColor(boardState, coord, enemyColor) {
					valid = append(valid, coord)
					break
				} else {
					break
				}
			}
		}
	}

	return valid
}

func getValidMovesKnight(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {

	valid := []Coord{}

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
				valid = append(valid, coord)
			}
		}

	}

	return valid
}

func getValidMovesForBishop(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	enemyColor := GetOppositeColor(playerColor)
	valid := []Coord{}

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
					valid = append(valid, coord)
				} else if isOccupiedByColor(boardState, coord, enemyColor) {
					valid = append(valid, coord)
					break
				} else {
					break
				}
			}
		}
	}

	return valid
}

func getValidMovesForQueen(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	diagonals := getValidMovesForBishop(playerColor, boardState, currCoord)
	horizontals := getValidMovesRook(playerColor, boardState, currCoord)
	valid := []Coord{}
	valid = append(valid, diagonals...)
	valid = append(valid, horizontals...)
	return valid
}

// IsDestinationValid checks if the specified color can move to the
func IsDestinationValid(whitesMove bool, isOccupied bool, occupant OnBoardData) bool {
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
