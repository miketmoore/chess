package chess

// PlayerColor is a type for player colors
type PlayerColor string

const (
	// PlayerWhite represents the white player
	PlayerWhite PlayerColor = "white"
	// PlayerBlack represents the black player
	PlayerBlack PlayerColor = "black"
)

func GetOppositeColor(color PlayerColor) PlayerColor {
	if color == PlayerWhite {
		return PlayerBlack
	}
	return PlayerWhite
}
