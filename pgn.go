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

func pieceToPGNPiece(pgn *string, piece Piece) {
	switch piece {
	case Knight:
		*pgn += "N"
	case King:
		*pgn += "K"
	case Queen:
		*pgn += "Q"
	case Rook:
		*pgn += "R"
	case Bishop:
		*pgn += "B"
	}
}

func HistoryToPGN(history []HistoryEntry) string {
	pgn := ""
	white := true
	i := 1
	for j, entry := range history {
		if white {
			pgn += fmt.Sprintf("%d. ", i)
		} else {
			i++
		}
		if entry.KingsideCastle {
			pgn += "O-O"
		} else if entry.QueensideCastle {
			pgn += "O-O-O"
		} else {
			pieceToPGNPiece(&pgn, entry.Piece)

			pgn += fmt.Sprintf("%s%s", fileByFileView[entry.ToCoord.File], rankByRankView[entry.ToCoord.Rank])
			if entry.Check {
				pgn += "+"
			} else if entry.Checkmate {
				pgn += "#"
				break
			}
			if entry.Promotion {
				pgn += "="
				pieceToPGNPiece(&pgn, entry.PromotedPiece)
			}
		}
		if j < len(history)-1 {
			pgn += " "
		}

		white = !white
	}
	return pgn
}
