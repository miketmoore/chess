package gamemodel

import (
	"github.com/miketmoore/chess/gamestate"
	"github.com/miketmoore/chess/model"
)

// Model contains data used for the game
type GameModel struct {
	BoardState           model.BoardState
	PieceToMove          model.PlayerPiece
	MoveStartCoord       model.Coord
	MoveDestinationCoord model.Coord
	Draw                 bool
	WhiteToMove          bool
	CurrentState         gamestate.GameState
}

// CurrentPlayerColor returns the current player color
func (m *GameModel) CurrentPlayerColor() model.PlayerColor {
	if m.WhiteToMove {
		return model.PlayerWhite
	}
	return model.PlayerBlack
}

// EnemyPlayerColor returns the enemy player color
func (m *GameModel) EnemyPlayerColor() model.PlayerColor {
	if m.WhiteToMove {
		return model.PlayerBlack
	}
	return model.PlayerWhite
}

// Move moves the current player piece
func (m *GameModel) Move(destCoord model.Coord) {
	m.CurrentState = gamestate.Draw
	m.Draw = true
	m.MoveDestinationCoord = destCoord

	m.BoardState[destCoord] = m.PieceToMove
	delete(m.BoardState, m.MoveStartCoord)

	m.WhiteToMove = !m.WhiteToMove
}
