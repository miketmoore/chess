package chess

import (
	"fmt"
	"strconv"
)

func isCoordStartPosition(playerColor PlayerColor, piece Piece, rank, file string) bool {

	if playerColor == PlayerWhite {
		// white
		if piece == Pawn {
			return rank == "2"
		}
	} else {
		// black
		if piece == Pawn {
			return rank == "7"
		}
	}

	return false
}

// Is the rank n spaces north a real board location?
func getRankAhead(rank string, n int) (string, bool) {
	if s, err := strconv.Atoi(rank); err == nil {
		fmt.Printf("%T, %v", s, s)
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

var ranks = []int{1, 2, 3, 4, 5, 6, 7, 8}
var files = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

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

func coordFromRankFile(rank int, file string) string {
	return fmt.Sprintf("%s%d", file, rank)
}

// GetPreviousFile gets the previous file as a string
func GetPreviousFile(file string) (string, bool) {
	for i, f := range files {
		if f == file && i-1 >= 0 {
			return files[i-1], true
		}
	}
	return "", false
}

// GetNextFile gets the next file as a string
func GetNextFile(file string) (string, bool) {
	for i, f := range files {
		if f == file && len(files) > i+1 {
			return files[i+1], true
		}
	}
	return "", false
}

// GetRelativeCoord gets a rank+file coordinate relative to specified by direction and distance
func GetRelativeCoord(rank, file string, direction Direction, distance int) (string, bool) {
	var rankInt int
	var err error
	if rankInt, err = strconv.Atoi(rank); err != nil {
		return "", false
	}
	switch direction {
	case North:
		newRank := rankInt + distance
		coord := coordFromRankFile(newRank, file)
		_, ok := validCoords[coord]
		return coord, ok
	case NorthWest:
		newRank := rankInt + distance
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case NorthEast:
		newRank := rankInt + distance
		newFile, ok := GetNextFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case South:
		newRank := rankInt - distance
		coord := coordFromRankFile(newRank, file)
		_, ok := validCoords[coord]
		return coord, ok
	case SouthWest:
		newRank := rankInt - distance
		newFile, ok := GetPreviousFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	case SouthEast:
		newRank := rankInt - distance
		newFile, ok := GetNextFile(file)
		if ok {
			coord := coordFromRankFile(newRank, newFile)
			_, ok := validCoords[coord]
			return coord, ok
		}
	}
	return "", false
}

// getRankAndFileFromSquareName converts a square name (example: d3) to rank (3)
// and file(d) strings
func getRankAndFileFromSquareName(squareName string) (rank, file string) {
	return string(squareName[1]), string(squareName[0])
}

// CanPawnMove checks if a pawn can move given a board state
// If a pawn is on it's starting square, then it is elligible to move one or two spaces forward.
// A pawn can move if an opposing piece is NW or NE.
func CanPawnMove(model Model, squareName string) []string {
	fmt.Println("can pawn move ", squareName)
	rank, file := getRankAndFileFromSquareName(squareName)

	// if pawn is on starting square, it is elligible for moving one or two spaces
	playerColor := model.CurrentPlayerColor()

	// build hash of valid board destinations
	valid := []string{}

	direction := North
	if playerColor == PlayerBlack {
		direction = South
	}

	// is one space ahead vacant?
	if coord, ok, _ := isRelCoordValid(model.BoardState, rank, file, direction, 1); ok {
		valid = append(valid, coord)
	}

	if isCoordStartPosition(playerColor, Pawn, rank, file) {

		// is two spaces ahead vacant?
		if coord, ok, _ := isRelCoordValid(model.BoardState, rank, file, direction, 2); ok {
			valid = append(valid, coord)
		}

	}

	// pawn attack moves
	if playerColor == PlayerWhite {
		// is NW occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, NorthWest, 1); ok {
			if occupant, isOccupied := model.BoardState[coord]; isOccupied {
				if occupant.Color == PlayerBlack {
					valid = append(valid, coord)
				}
			}
		}
		// // is NE occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, NorthEast, 1); ok {
			if occupant, isOccupied := model.BoardState[coord]; isOccupied {
				if occupant.Color == PlayerBlack {
					valid = append(valid, coord)
				}
			}
		}
	} else {
		// is SW occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, SouthWest, 1); ok {
			if occupant, isOccupied := model.BoardState[coord]; isOccupied {
				if occupant.Color == PlayerWhite {
					valid = append(valid, coord)
				}
			}
		}
		// is SE occupied by the enemy? if so, it is a valid move
		if coord, ok := GetRelativeCoord(rank, file, SouthEast, 1); ok {
			if occupant, isOccupied := model.BoardState[coord]; isOccupied {
				if occupant.Color == PlayerWhite {
					valid = append(valid, coord)
				}
			}
		}
	}

	fmt.Println(valid)
	return valid
}

// CanKingMove determines all valid moves for the King
func CanKingMove(model Model, squareName string) []string {
	fmt.Println("can king move ", squareName)
	rank, file := getRankAndFileFromSquareName(squareName)

	valid := []string{}

	directions := []Direction{North, South, East, West}
	for _, direction := range directions {
		if coord, ok, _ := isRelCoordValid(model.BoardState, rank, file, direction, 1); ok {
			valid = append(valid, coord)
		}
	}

	return valid
}

// GetNextRanks gets the series of ranks after
func GetNextRanks(rankStr string) []string {
	resp := []string{}
	if rankInt, err := strconv.Atoi(rankStr); err == nil {
		collect := false
		for _, r := range ranks {
			if !collect && r == rankInt {
				collect = true
			} else if collect {
				resp = append(resp, fmt.Sprintf("%d", r))
			}
		}
		return resp
	}
	return resp
}

// CanRookMove determines all valid moves for the Rook
func CanRookMove(model Model, squareName string) []string {
	fmt.Println("can rook move ", squareName)
	rank, file := getRankAndFileFromSquareName(squareName)

	valid := []string{}

	ranks := GetNextRanks(rank)
	fmt.Println("Next ranks: ", ranks)

	if coord, ok, _ := isRelCoordValid(model.BoardState, rank, file, North, 1); ok {
		valid = append(valid, coord)
	}

	return valid
}

func isRelCoordValid(boardState BoardState, rank, file string, direction Direction, n int) (string, bool, OnBoardData) {
	if coord, ok := GetRelativeCoord(rank, file, direction, n); ok {
		if occupant, isOccupied := boardState[coord]; !isOccupied {
			return coord, true, occupant
		}
	}
	return "", false, OnBoardData{}
}
