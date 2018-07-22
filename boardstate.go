package chess

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	return BoardState{
		Coord{Rank8, FileA}: OnBoardData{Color: PlayerBlack, Piece: Rook},
		Coord{Rank8, FileB}: OnBoardData{Color: PlayerBlack, Piece: Knight},
		Coord{Rank8, FileC}: OnBoardData{Color: PlayerBlack, Piece: Bishop},
		Coord{Rank8, FileD}: OnBoardData{Color: PlayerBlack, Piece: Queen},
		Coord{Rank8, FileE}: OnBoardData{Color: PlayerBlack, Piece: King},
		Coord{Rank8, FileF}: OnBoardData{Color: PlayerBlack, Piece: Bishop},
		Coord{Rank8, FileG}: OnBoardData{Color: PlayerBlack, Piece: Knight},
		Coord{Rank8, FileH}: OnBoardData{Color: PlayerBlack, Piece: Rook},

		Coord{Rank1, FileA}: OnBoardData{Color: PlayerWhite, Piece: Rook},
		Coord{Rank1, FileB}: OnBoardData{Color: PlayerWhite, Piece: Knight},
		Coord{Rank1, FileC}: OnBoardData{Color: PlayerWhite, Piece: Bishop},
		Coord{Rank1, FileD}: OnBoardData{Color: PlayerWhite, Piece: Queen},
		Coord{Rank1, FileE}: OnBoardData{Color: PlayerWhite, Piece: King},
		Coord{Rank1, FileF}: OnBoardData{Color: PlayerWhite, Piece: Bishop},
		Coord{Rank1, FileG}: OnBoardData{Color: PlayerWhite, Piece: Knight},
		Coord{Rank1, FileH}: OnBoardData{Color: PlayerWhite, Piece: Rook},
	}
}
