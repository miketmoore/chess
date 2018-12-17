package chess

func newPlayerPiece(color PlayerColor, piece Piece) PlayerPiece {
	return PlayerPiece{color, piece}
}

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	boardState := BoardState{
		NewCoord(FileA, Rank8): newPlayerPiece(PlayerBlack, Rook),
		NewCoord(FileB, Rank8): newPlayerPiece(PlayerBlack, Knight),
		NewCoord(FileC, Rank8): newPlayerPiece(PlayerBlack, Bishop),
		NewCoord(FileD, Rank8): newPlayerPiece(PlayerBlack, Queen),
		NewCoord(FileE, Rank8): newPlayerPiece(PlayerBlack, King),
		NewCoord(FileF, Rank8): newPlayerPiece(PlayerBlack, Bishop),
		NewCoord(FileG, Rank8): newPlayerPiece(PlayerBlack, Knight),
		NewCoord(FileH, Rank8): newPlayerPiece(PlayerBlack, Rook),

		NewCoord(FileA, Rank1): newPlayerPiece(PlayerWhite, Rook),
		NewCoord(FileB, Rank1): newPlayerPiece(PlayerWhite, Knight),
		NewCoord(FileC, Rank1): newPlayerPiece(PlayerWhite, Bishop),
		NewCoord(FileD, Rank1): newPlayerPiece(PlayerWhite, Queen),
		NewCoord(FileE, Rank1): newPlayerPiece(PlayerWhite, King),
		NewCoord(FileF, Rank1): newPlayerPiece(PlayerWhite, Bishop),
		NewCoord(FileG, Rank1): newPlayerPiece(PlayerWhite, Knight),
		NewCoord(FileH, Rank1): newPlayerPiece(PlayerWhite, Rook),
	}

	for _, file := range FilesOrder {
		coord := NewCoord(file, Rank7)
		boardState[coord] = newPlayerPiece(PlayerBlack, Pawn)
	}

	for _, file := range FilesOrder {
		coord := NewCoord(file, Rank2)
		boardState[coord] = newPlayerPiece(PlayerWhite, Pawn)
	}

	return boardState
}
