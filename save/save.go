package save

import (
	"fmt"
	"os"

	"github.com/miketmoore/chess/model"

	"github.com/miketmoore/chess/gamemodel"
)

func Save(game *gamemodel.GameModel) error {

	data := buildData(game)

	err := writeToFile(data)
	if err != nil {
		fmt.Println("Error writing save data to file")
		return err
	}

	return nil
}

func buildData(game *gamemodel.GameModel) string {
	currentPlayer := game.CurrentPlayerColor()
	boardState := game.BoardState

	var data string

	// Current player color (player to move)
	if currentPlayer == model.PlayerWhite {
		data = "0"
	} else {
		data = "1"
	}

	// Four digit string representing piece-color-rank-file
	for coord, playerPiece := range boardState {

		var color string // (0,1)
		if playerPiece.Color == model.PlayerWhite {
			color = "0"
		} else {
			color = "1"
		}

		// [0-1][0-5][0-8][0-8]
		s := fmt.Sprintf("%s%d%d%d", color, playerPiece.Piece, coord.Rank, coord.File)

		data = fmt.Sprintf("%s%s", data, s)
	}

	data = fmt.Sprintf("%s\n", data)

	return data
}

func writeToFile(data string) error {
	loc := "/tmp/chess_game"

	f, err := os.Create(loc)
	if err != nil {
		fmt.Println("Error creating file")
		return err
	}

	defer f.Close()

	_, err = f.WriteString(data)

	if err != nil {
		fmt.Println("Error writing file")
		return err
	}

	f.Sync()

	fmt.Printf("File written to: %s\n", loc)

	return nil
}
