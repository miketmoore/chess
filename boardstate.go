package chess

func newOnBoardData(color PlayerColor, piece Piece) OnBoardData {
	return OnBoardData{color, piece}
}

// InitialOnBoardState returns the board state for a new game of chess
func InitialOnBoardState() BoardState {
	boardState := BoardState{
		NewCoord(FileA, Rank8): newOnBoardData(PlayerBlack, Rook),
		NewCoord(FileB, Rank8): newOnBoardData(PlayerBlack, Knight),
		NewCoord(FileC, Rank8): newOnBoardData(PlayerBlack, Bishop),
		NewCoord(FileD, Rank8): newOnBoardData(PlayerBlack, Queen),
		NewCoord(FileE, Rank8): newOnBoardData(PlayerBlack, King),
		NewCoord(FileF, Rank8): newOnBoardData(PlayerBlack, Bishop),
		NewCoord(FileG, Rank8): newOnBoardData(PlayerBlack, Knight),
		NewCoord(FileH, Rank8): newOnBoardData(PlayerBlack, Rook),

		NewCoord(FileA, Rank1): newOnBoardData(PlayerWhite, Rook),
		NewCoord(FileB, Rank1): newOnBoardData(PlayerWhite, Knight),
		NewCoord(FileC, Rank1): newOnBoardData(PlayerWhite, Bishop),
		NewCoord(FileD, Rank1): newOnBoardData(PlayerWhite, Queen),
		NewCoord(FileE, Rank1): newOnBoardData(PlayerWhite, King),
		NewCoord(FileF, Rank1): newOnBoardData(PlayerWhite, Bishop),
		NewCoord(FileG, Rank1): newOnBoardData(PlayerWhite, Knight),
		NewCoord(FileH, Rank1): newOnBoardData(PlayerWhite, Rook),
	}

	for _, file := range FilesOrder {
		coord := NewCoord(file, Rank7)
		boardState[coord] = newOnBoardData(PlayerBlack, Pawn)
	}

	for _, file := range FilesOrder {
		coord := NewCoord(file, Rank2)
		boardState[coord] = newOnBoardData(PlayerWhite, Pawn)
	}

	return boardState
}
