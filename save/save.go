package save

import (
	"fmt"
	"os"
	"time"

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

	now := time.Now().String()

	data := fmt.Sprintf("%s\ncurrent=%t", now, currentPlayer)

	for coord, playerPiece := range boardState {
		x := fmt.Sprintf("color=%t;piece=%d;rank=%d;file=%d;", playerPiece.Color, playerPiece.Piece, coord.Rank, coord.File)
		data = fmt.Sprintf("%s\n%s", data, x)
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
