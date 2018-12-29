package chess

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	boardState := BoardState{
		Coord{File: FileA, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Rook},
		Coord{File: FileB, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Knight},
		Coord{File: FileC, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Bishop},
		Coord{File: FileD, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Queen},
		Coord{File: FileE, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: King},
		Coord{File: FileF, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Bishop},
		Coord{File: FileG, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Knight},
		Coord{File: FileH, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: Rook},

		Coord{File: FileA, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Rook},
		Coord{File: FileB, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Knight},
		Coord{File: FileC, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Bishop},
		Coord{File: FileD, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Queen},
		Coord{File: FileE, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: King},
		Coord{File: FileF, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Bishop},
		Coord{File: FileG, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Knight},
		Coord{File: FileH, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: Rook},
	}

	var file File

	for file = 1; file <= 8; file++ {
		coord := Coord{File: file, Rank: Rank7}
		boardState[coord] = PlayerPiece{Color: PlayerBlack, Piece: Pawn}
	}

	for file = 1; file <= 8; file++ {
		coord := Coord{File: file, Rank: Rank2}
		boardState[coord] = PlayerPiece{Color: PlayerWhite, Piece: Pawn}
	}

	return boardState
}
