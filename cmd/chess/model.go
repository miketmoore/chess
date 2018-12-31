package main

import (
	"github.com/miketmoore/chess/gamestate"
	"github.com/miketmoore/chess/model"
)

// Model contains data used for the game
type Model struct {
	BoardState           model.BoardState
	PieceToMove          model.PlayerPiece
	MoveStartCoord       model.Coord
	MoveDestinationCoord model.Coord
	Draw                 bool
	WhiteToMove          bool
	CurrentState         gamestate.GameState
}

// CurrentPlayerColor returns the current player color
func (m *Model) CurrentPlayerColor() model.PlayerColor {
	if m.WhiteToMove {
		return model.PlayerWhite
	}
	return model.PlayerBlack
}

// EnemyPlayerColor returns the enemy player color
func (m *Model) EnemyPlayerColor() model.PlayerColor {
	if m.WhiteToMove {
		return model.PlayerBlack
	}
	return model.PlayerWhite
}
