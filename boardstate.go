package chess

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	return BoardState{
		"a8": OnBoardData{Color: PlayerBlack, Piece: Rook},
		"b8": OnBoardData{Color: PlayerBlack, Piece: Knight},
		"c8": OnBoardData{Color: PlayerBlack, Piece: Bishop},
		"d8": OnBoardData{Color: PlayerBlack, Piece: Queen},
		"e8": OnBoardData{Color: PlayerBlack, Piece: King},
		"f8": OnBoardData{Color: PlayerBlack, Piece: Bishop},
		"g8": OnBoardData{Color: PlayerBlack, Piece: Knight},
		"h8": OnBoardData{Color: PlayerBlack, Piece: Rook},

		"a1": OnBoardData{Color: PlayerWhite, Piece: Rook},
		"b1": OnBoardData{Color: PlayerWhite, Piece: Knight},
		"c1": OnBoardData{Color: PlayerWhite, Piece: Bishop},
		"d1": OnBoardData{Color: PlayerWhite, Piece: Queen},
		"e1": OnBoardData{Color: PlayerWhite, Piece: King},
		"f1": OnBoardData{Color: PlayerWhite, Piece: Bishop},
		"g1": OnBoardData{Color: PlayerWhite, Piece: Knight},
		"h1": OnBoardData{Color: PlayerWhite, Piece: Rook},
	}
}
