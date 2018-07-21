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

// CanPawnMove checks if a pawn can move given a board state
func CanPawnMove(model Model, squareName string) []string {
	rank := string(squareName[1])
	file := string(squareName[0])
	fmt.Printf("CanPawnMove from %s\n", squareName)
	// if pawn is on starting square, it is elligible for moving one or two spaces
	playerColor := model.CurrentPlayerColor()
	onStart := isCoordStartPosition(playerColor, Pawn, rank, file)
	fmt.Printf("\tisStart:%v\n", onStart)

	// build hash of valid spaces to move
	validDestinations := []string{}
	if onStart {

		// is one space ahead vacant?
		rankOneAhead, oneAheadExists := getRankAhead(rank, 1)
		if oneAheadExists {
			fmt.Printf("rank ahead: %s\n", rankOneAhead)
			// is rankOneAhead occupied?
			_, isOccupied := model.BoardState[file+rankOneAhead]
			fmt.Printf("is occupied: %t\n", isOccupied)
			if !isOccupied {
				validDestinations = append(validDestinations, file+rankOneAhead)
			}
		}

		// is two spaces ahead vacant?
		rankTwoAhead, twoAheadExists := getRankAhead(rank, 2)
		if twoAheadExists {
			fmt.Printf("rank two ahead: %s\n", rankTwoAhead)
			_, isOccupied := model.BoardState[file+rankTwoAhead]
			fmt.Printf("is occupied: %t\n", isOccupied)
			if !isOccupied {
				validDestinations = append(validDestinations, file+rankTwoAhead)
			}
		}
	}
	fmt.Println(validDestinations)
	return validDestinations
}
