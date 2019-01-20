package load

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/miketmoore/chess/model"
)

func Load(path string) (model.PlayerColor, model.BoardState, error) {
	var currentColor model.PlayerColor
	boardState := model.BoardState{}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file")
		return currentColor, boardState, err
	}

	s := string(b)

	count := 0
	playerPiece := model.PlayerPiece{}
	coord := model.Coord{}

	for char, index := range s {
		if index == 0 {
			//  first char (current player color)
			if char == '0' {
				currentColor = model.PlayerWhite
			} else if char == '1' {
				currentColor = model.PlayerBlack
			} else {
				return currentColor, boardState, errors.New("Unexpected first character. Should be 0 or 1 to indicate current player color.")
			}
		} else if count == 0 {
			// piece color
			if char == '0' {
				playerPiece.Color = model.PlayerWhite
			} else if char == '1' {
				playerPiece.Color = model.PlayerBlack
			} else {
				return currentColor, boardState, errors.New("Expected 0 or 1 to indicate piece color.")
			}
			count++
		} else if count == 1 {
			d := model.Piece(char)
			if d < model.PiecePawn || d > model.PieceKing {
				return currentColor, boardState, errors.New("Unexpected value for piece.")
			} else {
				playerPiece.Piece = d
			}
			count++
		} else if count == 2 {
			d := model.Rank(char)
			if d < model.RankNone || d > model.Rank8 {
				return currentColor, boardState, errors.New("Unexpected value for piece rank.")
			} else {
				coord.Rank = d
			}
			count++
		} else if count == 3 {
			d := model.File(char)
			if d < model.FileNone || d > model.FileH {
				return currentColor, boardState, errors.New("Unexpected value for piece file.")
			} else {
				coord.File = d
			}
			boardState[coord] = playerPiece
			playerPiece = model.PlayerPiece{}
			coord = model.Coord{}
			count = 0
		}
	}

	return currentColor, boardState, nil
}
