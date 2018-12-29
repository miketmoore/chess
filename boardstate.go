package chess

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	boardState := BoardState{
		Coord{File: FileA, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceRook},
		Coord{File: FileB, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceKnight},
		Coord{File: FileC, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceBishop},
		Coord{File: FileD, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceQueen},
		Coord{File: FileE, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceKing},
		Coord{File: FileF, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceBishop},
		Coord{File: FileG, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceKnight},
		Coord{File: FileH, Rank: Rank8}: PlayerPiece{Color: PlayerBlack, Piece: PieceRook},

		Coord{File: FileA, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceRook},
		Coord{File: FileB, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceKnight},
		Coord{File: FileC, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceBishop},
		Coord{File: FileD, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceQueen},
		Coord{File: FileE, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceKing},
		Coord{File: FileF, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceBishop},
		Coord{File: FileG, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceKnight},
		Coord{File: FileH, Rank: Rank1}: PlayerPiece{Color: PlayerWhite, Piece: PieceRook},
	}

	var file File

	for file = 1; file <= 8; file++ {
		coord := Coord{File: file, Rank: Rank7}
		boardState[coord] = PlayerPiece{Color: PlayerBlack, Piece: PiecePawn}
	}

	for file = 1; file <= 8; file++ {
		coord := Coord{File: file, Rank: Rank2}
		boardState[coord] = PlayerPiece{Color: PlayerWhite, Piece: PiecePawn}
	}

	return boardState
}
