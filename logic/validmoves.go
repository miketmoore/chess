package logic

import (
	"github.com/miketmoore/chess/coordsmapper"
	"github.com/miketmoore/chess/state"
)

func isCoordStartPosition(playerColor state.PlayerColor, piece state.Piece, rank state.Rank) bool {

	if playerColor == state.PlayerWhite {
		// white
		if piece == state.PiecePawn {
			return rank == state.Rank2
		}
	} else {
		// black
		if piece == state.PiecePawn {
			return rank == state.Rank7
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
	SlopeNorth     = DirectionSlope{1, 0}
	SlopeNorthEast = DirectionSlope{1, 1}
	SlopeEast      = DirectionSlope{0, 1}
	SlopeSouthEast = DirectionSlope{-1, 1}
	SlopeSouth     = DirectionSlope{-1, 0}
	SlopeSouthWest = DirectionSlope{-1, -1}
	SlopeWest      = DirectionSlope{0, -1}
	SlopeNorthWest = DirectionSlope{1, -1}
)

type ValidMoves map[state.Coord]uint8

// GetValidMoves returns a list of valid coordinates the piece can be moved to
func GetValidMoves(playerColor state.PlayerColor, piece state.Piece, boardState state.BoardState,
	coord state.Coord) ValidMoves {
	switch piece {
	case state.PiecePawn:
		return getValidMovesPawn(playerColor, boardState, coord)
	case state.PieceKing:
		return getValidMovesKing(playerColor, boardState, coord)
	case state.PieceKnight:
		return getValidMovesKnight(playerColor, boardState, coord)
	case state.PieceRook:
		return getValidMovesRook(playerColor, boardState, coord)
	case state.PieceBishop:
		return getValidMovesForBishop(playerColor, boardState, coord)
	case state.PieceQueen:
		return getValidMovesForQueen(playerColor, boardState, coord)
	}
	return ValidMoves{}
}

func isOccupied(boardState state.BoardState, coord state.Coord) bool {
	_, isOccupied := boardState[coord]
	return isOccupied
}

func isOccupiedByColor(boardState state.BoardState, coord state.Coord, color state.PlayerColor) bool {
	occupant, occupied := boardState[coord]
	return occupied && occupant.Color == color
}

func getValidMovesPawn(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {
	enemyColor := !playerColor
	valid := ValidMoves{}

	yChange := 1
	if playerColor == state.PlayerBlack {
		yChange = -1
	}

	// get two spaces north or south
	coords := coordsmapper.GetCoordsBySlopeAndDistance(currCoord, yChange, 0, 2)
	if !isOccupied(boardState, coords[0]) {
		valid[coords[0]] = 1
		if isCoordStartPosition(playerColor, state.PiecePawn, currCoord.Rank) && !isOccupied(boardState, coords[1]) {
			valid[coords[1]] = 1
		}
	}

	// pawn attack moves
	if playerColor == state.PlayerWhite {
		// NW
		coord, ok := coordsmapper.GetCoordBySlopeAndDistance(currCoord, 1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}

		// NE
		coord, ok = coordsmapper.GetCoordBySlopeAndDistance(currCoord, 1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	} else {
		// SW
		coord, ok := coordsmapper.GetCoordBySlopeAndDistance(currCoord, -1, -1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}

		// SE
		coord, ok = coordsmapper.GetCoordBySlopeAndDistance(currCoord, -1, 1)
		if ok && isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	}

	return valid
}

func getValidMovesKing(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {
	enemyColor := !playerColor
	valid := ValidMoves{}

	coords := coordsmapper.GetCoordsBySlopeAndDistanceAll(currCoord, 1)
	for _, coord := range coords {
		if !isOccupied(boardState, coord) || isOccupiedByColor(boardState, coord, enemyColor) {
			valid[coord] = 1
		}
	}

	return valid
}

func getValidMovesRook(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {
	enemyColor := !playerColor
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
			coords := coordsmapper.GetCoordsBySlopeAndDistance(currCoord, yChange, xChange, i)
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

func getValidMovesKnight(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {

	valid := ValidMoves{}

	var slopes = [][]int{
		{2, -1},
		{2, 1},
		{1, 2},
		{-1, 2},
		{-2, 1},
		{-2, -1},
		{1, -2},
		{-1, -2},
	}

	for _, slope := range slopes {
		coords := coordsmapper.GetCoordsBySlopeAndDistance(currCoord, slope[0], slope[1], 1)
		for _, coord := range coords {
			if !isOccupied(boardState, coord) || isOccupiedByColor(boardState, coord, !playerColor) {
				valid[coord] = 1
			}
		}

	}

	return valid
}

func getValidMovesForBishop(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {
	enemyColor := !playerColor
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
			coords := coordsmapper.GetCoordsBySlopeAndDistance(currCoord, yChange, xChange, i)
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

func getValidMovesForQueen(playerColor state.PlayerColor, boardState state.BoardState, currCoord state.Coord) ValidMoves {
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
func IsDestinationValid(whiteToMove bool, isOccupied bool, occupant state.PlayerPiece) bool {
	if isOccupied {
		if whiteToMove && occupant.Color == state.PlayerBlack {
			return true
		} else if !whiteToMove && occupant.Color == state.PlayerWhite {
			return true
		}
	} else {
		return true
	}
	return false
}
