package chess

// PlayerColor is a type for player colors
type PlayerColor uint8

const (
	PlayerWhite PlayerColor = iota
	PlayerBlack
)

func GetOppositeColor(color PlayerColor) PlayerColor {
	if color == PlayerWhite {
		return PlayerBlack
	}
	return PlayerWhite
}
