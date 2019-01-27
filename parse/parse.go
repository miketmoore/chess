package parse

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/miketmoore/chess/model"
)

func Parse(data string) (model.PlayerColor, model.BoardState, error) {
	var currentColor model.PlayerColor
	boardState := model.BoardState{}

	count := 0
	playerPiece := model.PlayerPiece{}
	coord := model.Coord{}

	for index, r := range data {
		if r == '\n' {
			fmt.Println("Newline encountered, ending parse")
			break
		}

		parsed, err := strconv.ParseUint(string(r), 10, 32)
		if err != nil {
			fmt.Println(err)
			return currentColor, boardState, errors.New("Failed to parse rune into uint")
		}

		if index == 0 {
			//  first char (current player color)
			if r == '0' {
				currentColor = model.PlayerWhite
			} else if r == '1' {
				currentColor = model.PlayerBlack
			} else {
				return currentColor, boardState, errors.New("Unexpected first character. Should be 0 or 1 to indicate current player color.")
			}
		} else if count == 0 {
			// piece color
			if r == '0' {
				playerPiece.Color = model.PlayerWhite
			} else if r == '1' {
				playerPiece.Color = model.PlayerBlack
			} else {
				return currentColor, boardState, errors.New("Expected 0 or 1 to indicate piece color.")
			}
			count++
		} else if count == 1 {
			d := model.Piece(parsed)

			if d < model.PiecePawn || d > model.PieceKing {
				return currentColor, boardState, errors.New("Unexpected value for piece.")
			} else {
				playerPiece.Piece = d
			}
			count++
		} else if count == 2 {
			d := model.Rank(parsed)

			if d < model.RankNone || d > model.Rank8 {
				return currentColor, boardState, errors.New("Unexpected value for rank.")
			} else {
				coord.Rank = d
			}
			count++
		} else if count == 3 {
			d := model.File(parsed)

			if d < model.FileNone || d > model.FileH {
				return currentColor, boardState, errors.New("Unexpected value for file.")
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
