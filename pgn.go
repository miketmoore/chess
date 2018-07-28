package chess

import "fmt"

/*

https://www.chessclub.com/help/PGN-spec

P pawn
N knight
B bishop
R rook
Q queen
K king

Move construction

* list moving piece letter (omitted for pawns), followed by destination square (file, rank).File
* a4 = pawn moves to a4
* Na5 = knight moves to a5

kingside castling O-O
queenside castling O-O-O

promotion: a4=Q  ... pawn to a4, promoted to queen
en passant: no special notation. formed as if the captured pawn were on the capturing pawn's destination square

check move: add + as suffix
checkmate move: add # as suffix
*/

func HistoryToPGN(history []HistoryEntry) string {
	pgn := ""
	white := true
	i := 1
	for _, entry := range history {
		if white {
			pgn += fmt.Sprintf("%d. ", i)
		} else {
			i++
		}
		switch entry.Piece {
		case Knight:
			pgn += "N"
		case King:
			pgn += "K"
		case Queen:
			pgn += "Q"
		case Rook:
			pgn += "R"
		case Bishop:
			pgn += "B"
		}
		pgn += fmt.Sprintf("%s%s", fileByFileView[entry.ToCoord.File], rankByRankView[entry.ToCoord.Rank])
		if entry.Check {
			pgn += "+"
		}
		pgn += " "
		white = !white
	}
	return pgn
}
