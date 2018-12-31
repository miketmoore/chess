package coordsmapper

import (
	"github.com/miketmoore/chess/state"
)

var rankToY = map[state.Rank]int{
	state.Rank1: 0,
	state.Rank2: 1,
	state.Rank3: 2,
	state.Rank4: 3,
	state.Rank5: 4,
	state.Rank6: 5,
	state.Rank7: 6,
	state.Rank8: 7,
}

var yToRank = map[int]state.Rank{
	0: state.Rank1,
	1: state.Rank2,
	2: state.Rank3,
	3: state.Rank4,
	4: state.Rank5,
	5: state.Rank6,
	6: state.Rank7,
	7: state.Rank8,
}

var fileToX = map[state.File]int{
	state.FileA: 0,
	state.FileB: 1,
	state.FileC: 2,
	state.FileD: 3,
	state.FileE: 4,
	state.FileF: 5,
	state.FileG: 6,
	state.FileH: 7,
}

var xToFile = map[int]state.File{
	0: state.FileA,
	1: state.FileB,
	2: state.FileC,
	3: state.FileD,
	4: state.FileE,
	5: state.FileF,
	6: state.FileG,
	7: state.FileH,
}

// TranslateRankFileToXY translates [state.File,state.Rank] coordinates to [x,y] coordinates
func TranslateRankFileToXY(coord state.Coord) (int, int) {
	return fileToX[coord.File], rankToY[coord.Rank]
}

// TranslateXYToRankFile translates [x,y] coordinates to [state.File,state.Rank] coordinates
func TranslateXYToRankFile(x, y int) state.Coord {
	return state.Coord{File: xToFile[x], Rank: yToRank[y]}
}

// Getstate.CoordsBySlopeAndDistance gets a list of coordinates (file,rank)
func GetCoordsBySlopeAndDistance(start state.Coord, yChange, xChange, distance int) []state.Coord {
	x, y := TranslateRankFileToXY(start)

	coords := []state.Coord{}

	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := state.ValidCoords[coord]
		if ok {
			coords = append(coords, coord)
		}
		y += yChange
		x += xChange
		d++
	}

	return coords
}

func GetCoordBySlopeAndDistance(start state.Coord, yChange, xChange int) (state.Coord, bool) {
	x, y := TranslateRankFileToXY(start)

	distance := 1
	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := state.ValidCoords[coord]
		if ok {
			return coord, true
		}
		y += yChange
		x += xChange
		d++
	}
	return state.Coord{}, false
}

func GetCoordsBySlopeAndDistanceAll(start state.Coord, distance int) []state.Coord {
	slopes := [][]int{
		{1, 0},   // n
		{1, 1},   // ne
		{0, 1},   // e
		{-1, 1},  // se
		{-1, 0},  // s
		{-1, -1}, // sw
		{0, -1},  // w
		{1, -1},  // nw
	}

	coords := []state.Coord{}

	for _, slope := range slopes {
		coords = append(coords, GetCoordsBySlopeAndDistance(start, slope[0], slope[1], distance)...)
	}

	return coords
}

// GetCoordByXY gets algebraic notation for a set of
// rank (y) and file (x) coordinates
func GetCoordByXY(
	squareOriginByCoords map[state.Coord][]float64,
	x, y float64,
) (state.Coord, bool) {
	for coord, xy := range squareOriginByCoords {
		if xy[0] == x && xy[1] == y {
			return coord, true
		}
	}
	return state.Coord{}, false
}
