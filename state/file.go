package state

// File is a custom type that represents a vertical column (file) on the chess board
type File uint8

const (
	FileNone File = iota
	FileA
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)
