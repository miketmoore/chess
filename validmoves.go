package chess

import (
	"fmt"
	"strconv"
)

func isCoordStartPosition(playerColor PlayerColor, piece Piece, rank Rank, file File) bool {

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

// Is the rank n spaces north a real board location?
func getRankAhead(rank string, n int) (string, bool) {
	if s, err := strconv.Atoi(rank); err == nil {
		return fmt.Sprintf("%d", s+n), true
	}
	return "", false
}

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

type Rank int
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

var ranks = []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}
var files = []File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

var validCoords = map[string]bool{
	"a1": true,
	"a2": true,
	"a3": true,
	"a4": true,
	"a5": true,
	"a6": true,
	"a7": true,
	"a8": true,

	"b1": true,
	"b2": true,
	"b3": true,
	"b4": true,
	"b5": true,
	"b6": true,
	"b7": true,
	"b8": true,

	"c1": true,
	"c2": true,
	"c3": true,
	"c4": true,
	"c5": true,
	"c6": true,
	"c7": true,
	"c8": true,

	"d1": true,
	"d2": true,
	"d3": true,
	"d4": true,
	"d5": true,
	"d6": true,
	"d7": true,
	"d8": true,

	"e1": true,
	"e2": true,
	"e3": true,
	"e4": true,
	"e5": true,
	"e6": true,
	"e7": true,
	"e8": true,

	"f1": true,
	"f2": true,
	"f3": true,
	"f4": true,
	"f5": true,
	"f6": true,
	"f7": true,
	"f8": true,

	"g1": true,
	"g2": true,
	"g3": true,
	"g4": true,
	"g5": true,
	"g6": true,
	"g7": true,
	"g8": true,

	"h1": true,
	"h2": true,
	"h3": true,
	"h4": true,
	"h5": true,
	"h6": true,
	"h7": true,
	"h8": true,
}

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

func coordFromRankFile(rank Rank, file File) string {
	return fmt.Sprintf("%s%s", fileByFileView[file], rankByRankView[rank])
}

// GetPreviousFile gets the previous file as a string
func GetPreviousFile(file File) (File, bool) {
	for i, f := range files {
		if f == file && i-1 >= 0 {
			return files[i-1], true
		}
	}
	return FileNone, false
}

// GetNextFile gets the next file as a string
func GetNextFile(file File) (File, bool) {
	for i, f := range files {
		if f == file && len(files) > i+1 {
			return files[i+1], true
		}
	}
	return FileNone, false
}

// GetRelativeCoord gets a rank+file coordinate relative to specified by direction and distance
func GetRelativeCoord(rank Rank, file File, direction Direction, distance int) (string, bool) {
	switch direction {
	case North:
		newRank := rank + Rank(distance)
		coord := coordFromRankFile(newRank, file)
		_, ok := validCoords[coord]
		return coord, ok
	case NorthWest:
		newRank := rank + Rank(distance)
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case NorthEast:
		newRank := rank + Rank(distance)
		newFile, ok := GetNextFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case South:
		newRank := rank - Rank(distance)
		coord := coordFromRankFile(newRank, file)
		_, ok := validCoords[coord]
		return coord, ok
	case East:
		var newFile = file
		var ok bool
		for i := 0; i < distance; i++ {
			newFile, ok = GetNextFile(newFile)
		}
		if ok {
			coord := coordFromRankFile(rank, newFile)
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
			coord := coordFromRankFile(rank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case SouthWest:
		newRank := rank - Rank(distance)
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case SouthEast:
		newRank := rank - Rank(distance)
		newFile, ok := GetNextFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	}
	return "", false
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

// getRankAndFileFromSquareName converts a square name (example: d3) to rank (3)
// and file(d) strings
func getRankAndFileFromSquareName(squareName string) (Rank, File) {
	rankStr := string(squareName[1])
	fileStr := string(squareName[0])

	return rankViewByRank[rankStr], fileViewByFile[fileStr]
}

// GetValidMoves returns a list of valid coordinates the piece can be moved to
func GetValidMoves(playerColor PlayerColor, piece Piece, boardState BoardState, squareName string) []string {
	switch piece {
	case Pawn:
		return canPawnMove(playerColor, boardState, squareName)
	case King:
		return canKingMove(boardState, squareName)
	case Knight:
		return canKnightMove(boardState, squareName)
	case Rook:
		return canRookMove(boardState, squareName)
	}
	return []string{}
}

func canPawnMove(playerColor PlayerColor, boardState BoardState, squareName string) []string {
	rank, file := getRankAndFileFromSquareName(squareName)

	// if pawn is on starting square, it is elligible for moving one or two spaces

	// build hash of valid board destinations
	valid := []string{}

	direction := North
	if playerColor == PlayerBlack {
		direction = South
	}

	// is one space ahead vacant?
	if coord, ok, _ := IsRelCoordValid(boardState, rank, file, direction, 1); ok {
		valid = append(valid, coord)
	}

	if isCoordStartPosition(playerColor, Pawn, rank, file) {

		// is two spaces ahead vacant?
		if coord, ok, _ := IsRelCoordValid(boardState, rank, file, direction, 2); ok {
			valid = append(valid, coord)
		}

	}

	// pawn attack moves
	if playerColor == PlayerWhite {
		// is NW occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, NorthWest, 1); ok {
			if occupant, isOccupied := boardState[coord]; isOccupied {
				if occupant.Color == PlayerBlack {
					valid = append(valid, coord)
				}
			}
		}
		// // is NE occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, NorthEast, 1); ok {
			if occupant, isOccupied := boardState[coord]; isOccupied {
				if occupant.Color == PlayerBlack {
					valid = append(valid, coord)
				}
			}
		}
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

func canKingMove(boardState BoardState, squareName string) []string {
	rank, file := getRankAndFileFromSquareName(squareName)

	valid := []string{}

	directions := []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
	for _, direction := range directions {
		if coord, ok, _ := IsRelCoordValid(boardState, rank, file, direction, 1); ok {
			valid = append(valid, coord)
		}
	}

	return valid
}

func canRookMove(boardState BoardState, squareName string) []string {
	rank, file := getRankAndFileFromSquareName(squareName)
	valid := []string{}

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

func canKnightMove(boardState BoardState, squareName string) []string {
	rank, file := getRankAndFileFromSquareName(squareName)

	valid := []string{}

	all := [][]pieceMove{
		[]pieceMove{
			pieceMove{North, 2},
			pieceMove{West, 1},
		},
		[]pieceMove{
			pieceMove{North, 2},
			pieceMove{East, 1},
		},
		[]pieceMove{
			pieceMove{East, 2},
			pieceMove{North, 1},
		},
		[]pieceMove{
			pieceMove{East, 2},
			pieceMove{South, 1},
		},
		[]pieceMove{
			pieceMove{South, 2},
			pieceMove{East, 1},
		},
		[]pieceMove{
			pieceMove{South, 2},
			pieceMove{West, 1},
		},
		[]pieceMove{
			pieceMove{West, 2},
			pieceMove{South, 1},
		},
		[]pieceMove{
			pieceMove{West, 2},
			pieceMove{North, 1},
		},
	}

	for _, moves := range all {
		if coord, ok := checkKnightMove(boardState, rank, file, moves); ok {
			valid = append(valid, coord)
		}
	}

	return valid
}

func checkKnightMove(boardState BoardState, rank Rank, file File, moves []pieceMove) (string, bool) {
	if coord, ok := GetRelativeCoord(rank, file, moves[0].Direction, moves[0].Distance); ok {
		rank, file := getRankAndFileFromSquareName(coord)
		if coord, ok, _ := IsRelCoordValid(boardState, rank, file, moves[1].Direction, moves[1].Distance); ok {
			return coord, true
		}
	}
	return "", false
}

// GetNextRanks gets the series of ranks after
func GetNextRanks(rank Rank) []Rank {
	resp := []Rank{}
	collect := false
	for _, r := range ranks {
		if !collect && r == rank {
			collect = true
		} else if collect {
			resp = append(resp, r)
		}
	}
	return resp
}

// IsRelCoordValid checks if the specified coordinate is valid
// It is valid if it exists and not occupied
func IsRelCoordValid(boardState BoardState, rank Rank, file File, direction Direction, n int) (string, bool, OnBoardData) {
	if coord, ok := GetRelativeCoord(rank, file, direction, n); ok {
		occupant, isOccupied := boardState[coord]
		if !isOccupied {
			return coord, true, occupant
		}
	}
	return "", false, OnBoardData{}
}
