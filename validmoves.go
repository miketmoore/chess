package chess

import "fmt"

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
type Direction string

const (
	North     Direction = "north"
	NorthWest Direction = "northwest"
	NorthEast Direction = "northeast"
	South     Direction = "south"
	SouthWest Direction = "southwest"
	SouthEast Direction = "southeast"
	East      Direction = "east"
	West      Direction = "west"
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

// GetPreviousFile gets the previous file as a string
func GetPreviousFile(file File) (File, bool) {
	for i, f := range FilesOrder {
		if f == file && i-1 >= 0 {
			return FilesOrder[i-1], true
		}
	}
	return FileNone, false
}

// GetNextFile gets the next file as a string
func GetNextFile(file File) (File, bool) {
	for i, f := range FilesOrder {
		if f == file && len(FilesOrder) > i+1 {
			return FilesOrder[i+1], true
		}
	}
	return FileNone, false
}

// GetRelativeCoord gets a rank+file coordinate relative to specified by direction and distance
func GetRelativeCoord(rank Rank, file File, direction Direction, distance int) (Coord, bool) {
	switch direction {
	case North:
		newRank := rank + Rank(distance)
		coord := NewCoord(file, newRank)
		_, ok := validCoords[coord]
		return coord, ok
	case NorthWest:
		newRank := rank + Rank(distance)
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := NewCoord(newFile, newRank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case NorthEast:
		newRank := rank + Rank(distance)
		newFile, ok := GetNextFile(file)
		if ok {
			coord := NewCoord(newFile, newRank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case South:
		newRank := rank - Rank(distance)
		coord := NewCoord(file, newRank)
		_, ok := validCoords[coord]
		return coord, ok
	case East:
		var newFile = file
		var ok bool
		for i := 0; i < distance; i++ {
			newFile, ok = GetNextFile(newFile)
		}
		if ok {
			coord := NewCoord(newFile, rank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case West:
		var newFile = file
		var ok bool
		for i := distance; i > 0; i-- {
			newFile, ok = GetPreviousFile(newFile)
		}
		if ok {
			coord := NewCoord(newFile, rank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case SouthWest:
		newRank := rank - Rank(distance)
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := NewCoord(newFile, newRank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case SouthEast:
		newRank := rank - Rank(distance)
		newFile, ok := GetNextFile(file)
		if ok {
			coord := NewCoord(newFile, newRank)
			_, ok := validCoords[coord]
			return coord, ok
		}
	}
	return Coord{}, false
}

// GetValidMoves returns a list of valid coordinates the piece can be moved to
func GetValidMoves(playerColor PlayerColor, piece Piece, boardState BoardState, coord Coord) []Coord {
	switch piece {
	case Pawn:
		return canPawnMove(playerColor, boardState, coord)
	case King:
		return canKingMove(boardState, coord)
	case Knight:
		return canKnightMove(boardState, coord)
	case Rook:
		return canRookMove(boardState, coord)
	case Bishop:
		return getValidMovesForBishop(boardState, coord)
	}
	return []Coord{}
}

func canPawnMove(playerColor PlayerColor, boardState BoardState, currCoord Coord) []Coord {
	rank, file := currCoord.GetRankFile()

	// if pawn is on starting square, it is elligible for moving one or two spaces

	// build hash of valid board destinations
	valid := []Coord{}

	yChange := 1
	if playerColor == PlayerBlack {
		yChange = -1
	}

	// get two spaces north
	coords := GetCoordsBySlopeAndDistance(currCoord, yChange, 0, 2)
	valid = append(valid, coords...)

	// pawn attack moves
	if playerColor == PlayerWhite {
		// NW
		coords := GetCoordsBySlopeAndDistance(currCoord, 1, 1, 1)
		valid = append(valid, coords...)

		// NE
		coords = GetCoordsBySlopeAndDistance(currCoord, 1, -1, 1)
		valid = append(valid, coords...)
	} else {
		// is SW occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, SouthWest, 1); ok {
			if occupant, isOccupied := boardState[coord]; isOccupied {
				if occupant.Color == PlayerWhite {
					valid = append(valid, coord)
				}
			}
		}
		// is SE occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, SouthEast, 1); ok {
			if occupant, isOccupied := boardState[coord]; isOccupied {
				if occupant.Color == PlayerWhite {
					valid = append(valid, coord)
				}
			}
		}
	}

	return valid
}

func canKingMove(boardState BoardState, currCoord Coord) []Coord {
	rank, file := currCoord.GetRankFile()

	valid := []Coord{}

	directions := []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
	for _, direction := range directions {
		if coord, ok, _ := IsRelCoordValid(boardState, rank, file, direction, 1); ok {
			valid = append(valid, coord)
		}
	}

	return valid
}

func canRookMove(boardState BoardState, currCoord Coord) []Coord {
	rank, file := currCoord.GetRankFile()
	valid := []Coord{}

	directions := []Direction{North, East, South, West}
	for _, direction := range directions {
		for i := 0; i < 8; i++ {
			if coord, ok, _ := IsRelCoordValid(boardState, rank, file, direction, i+1); ok {
				valid = append(valid, coord)
			} else {
				break
			}
		}
	}

	return valid
}

type pieceMove struct {
	Direction Direction
	Distance  int
}

func newMove(direction Direction, distance int) pieceMove {
	return pieceMove{
		direction,
		distance,
	}
}

var knightMoves = [][]pieceMove{
	[]pieceMove{newMove(North, 2), newMove(West, 1)},
	[]pieceMove{newMove(North, 2), newMove(East, 1)},
	[]pieceMove{newMove(East, 2), newMove(North, 1)},
	[]pieceMove{newMove(East, 2), newMove(South, 1)},
	[]pieceMove{newMove(South, 2), newMove(East, 1)},
	[]pieceMove{newMove(South, 2), newMove(West, 1)},
	[]pieceMove{newMove(West, 2), newMove(South, 1)},
	[]pieceMove{newMove(West, 2), newMove(North, 1)},
}

func canKnightMove(boardState BoardState, currCoord Coord) []Coord {
	rank, file := currCoord.GetRankFile()

	valid := []Coord{}

	for _, moves := range knightMoves {
		if coord, ok := checkKnightMove(boardState, rank, file, moves); ok {
			valid = append(valid, coord)
		}
	}

	return valid
}

func checkKnightMove(boardState BoardState, rank Rank, file File, moves []pieceMove) (Coord, bool) {
	coord, ok := GetRelativeCoord(rank, file, moves[0].Direction, moves[0].Distance)
	if !ok {
		return Coord{}, false
	}
	coord, ok, _ = IsRelCoordValid(boardState, coord.Rank, coord.File, moves[1].Direction, moves[1].Distance)
	if ok {
		return coord, true
	}
	return Coord{}, false
}

func getValidMovesForBishop(boardState BoardState, currCoord Coord) []Coord {
	// rank, file := currCoord.GetRankFile()

	valid := []Coord{}

	// NorthEast slope: +1/+1
	// follow slope from current coordinate
	// x, y := TranslateRankFileToXY(currCoord)
	// fmt.Println(">>> ", x, y)
	valid = append(valid, GetCoordsBySlope(currCoord, 1, 1)...)
	// valid = append(valid, getCoordsBySlope(currCoord, 1, -1)...)
	// valid = append(valid, getCoordsBySlope(currCoord, -1, 1)...)
	// valid = append(valid, getCoordsBySlope(currCoord, -1, -1)...)

	// ranks := append([]Rank{rank}, GetNextRanks(rank)...)

	/*
		   -  e3	north 1, east 1 (from d2)
		-  d2		north 1, east 1 (from c1)
		c1			start
	*/

	// for i, r := range ranks {
	// 	coordA, _ := GetRelativeCoord(r, file, North, 1)
	// 	coordB, _ := GetRelativeCoord(coordA.Rank, coordA.File, East, i+1)
	// 	if coordB.Rank != 0 {
	// 		valid = append(valid, coordB)
	// 		fStr := fileByFileView[coordB.File]
	// 		rStr := rankByRankView[coordB.Rank]
	// 		fmt.Println(fStr, rStr)
	// 	}

	// }

	// for i, r := range ranks {
	// 	coordA, _ := GetRelativeCoord(r, file, North, 1)
	// 	coordB, _ := GetRelativeCoord(coordA.Rank, coordA.File, West, i+1)
	// 	if coordB.Rank != 0 {
	// 		valid = append(valid, coordB)
	// 		fStr := fileByFileView[coordB.File]
	// 		rStr := rankByRankView[coordB.Rank]
	// 		fmt.Println(fStr, rStr)
	// 	}
	// }

	// ranks = append([]Rank{rank}, GetPreviousRanks(rank)...)

	// for i, r := range ranks {
	// 	coordA, _ := GetRelativeCoord(r, file, South, 1)
	// 	coordB, _ := GetRelativeCoord(coordA.Rank, coordA.File, East, i+1)
	// 	if coordB.Rank != 0 {
	// 		valid = append(valid, coordB)
	// 		fStr := fileByFileView[coordB.File]
	// 		rStr := rankByRankView[coordB.Rank]
	// 		fmt.Println(fStr, rStr)
	// 	}

	// }

	fmt.Println(valid)

	return valid
}

// GetNextRanks gets the series of ranksOrder after
func GetNextRanks(rank Rank) []Rank {
	resp := []Rank{}
	collect := false
	for _, r := range ranksOrder {
		if !collect && r == rank {
			collect = true
		} else if collect {
			resp = append(resp, r)
		}
	}
	return resp
}

// GetPreviousRanks gets the seris of ranks before
func GetPreviousRanks(rank Rank) []Rank {
	resp := []Rank{}
	collect := true
	for _, r := range ranksOrder {
		if collect {
			if r != rank {
				resp = append(resp, r)
			} else {
				collect = false
			}
		}
	}
	return resp
}

// IsRelCoordValid checks if the specified coordinate is valid
// It is valid if it exists and not occupied
func IsRelCoordValid(boardState BoardState, rank Rank, file File, direction Direction, n int) (Coord, bool, OnBoardData) {
	coord, ok := GetRelativeCoord(rank, file, direction, n)
	if ok {
		occupant, isOccupied := boardState[coord]
		if !isOccupied {
			return coord, true, occupant
		}
	}
	return Coord{}, false, OnBoardData{}
}
