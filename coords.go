package chess

import chessapi "github.com/miketmoore/chess-api"

// GetFileRankByXY a coordinate for a set of rank (y) and file (x) coordinates
func GetFileRankByXY(
	squareOriginByCoords map[chessapi.Coord][]float64,
	x, y float64,
) (chessapi.Coord, bool) {
	for coord, xy := range squareOriginByCoords {
		if xy[0] == x && xy[1] == y {
			return coord, true
		}
	}
	return chessapi.Coord{}, false
}
