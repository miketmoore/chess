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

func getRankAhead(rank string, n int) (string, bool) {
	if s, err := strconv.Atoi(rank); err == nil {
		fmt.Printf("%T, %v", s, s)
		return fmt.Sprintf("%d", s+n), true
	}
	return "", false
}

// GetRankAndFileFromSquareName converts a square name (example: d3) to rank (3)
// and file(d) strings
func GetRankAndFileFromSquareName(squareName string) (rank, file string) {
	return string(squareName[1]), string(squareName[0])
}

// CanPawnMove checks if a pawn can move given a board state
// If a pawn is on it's starting square, then it is elligible to move one or two spaces forward.
// A pawn can move if an opposing piece is NW or NE.
func CanPawnMove(model Model, squareName string) []string {
	rank, file := GetRankAndFileFromSquareName(squareName)

	// if pawn is on starting square, it is elligible for moving one or two spaces
	playerColor := model.CurrentPlayerColor()

	// build hash of valid board destinations
	validDestinations := []string{}
	if isCoordStartPosition(playerColor, Pawn, rank, file) {

		// is one space ahead vacant?
		rankOneAhead, oneAheadExists := getRankAhead(rank, 1)
		if oneAheadExists {
			// is rankOneAhead occupied?
			_, isOccupied := model.BoardState[file+rankOneAhead]
			if !isOccupied {
				validDestinations = append(validDestinations, file+rankOneAhead)
			}
		}

		// is two spaces ahead vacant?
		rankTwoAhead, twoAheadExists := getRankAhead(rank, 2)
		if twoAheadExists {
			_, isOccupied := model.BoardState[file+rankTwoAhead]
			if !isOccupied {
				validDestinations = append(validDestinations, file+rankTwoAhead)
			}
		}
	}
	fmt.Println(validDestinations)
	return validDestinations
}
