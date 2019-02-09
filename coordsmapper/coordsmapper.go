package coordsmapper

import "github.com/miketmoore/chess-api/model"

var rankToY = map[model.Rank]int{
	model.Rank1: 0,
	model.Rank2: 1,
	model.Rank3: 2,
	model.Rank4: 3,
	model.Rank5: 4,
	model.Rank6: 5,
	model.Rank7: 6,
	model.Rank8: 7,
}

var yToRank = map[int]model.Rank{
	0: model.Rank1,
	1: model.Rank2,
	2: model.Rank3,
	3: model.Rank4,
	4: model.Rank5,
	5: model.Rank6,
	6: model.Rank7,
	7: model.Rank8,
}

var fileToX = map[model.File]int{
	model.FileA: 0,
	model.FileB: 1,
	model.FileC: 2,
	model.FileD: 3,
	model.FileE: 4,
	model.FileF: 5,
	model.FileG: 6,
	model.FileH: 7,
}

var xToFile = map[int]model.File{
	0: model.FileA,
	1: model.FileB,
	2: model.FileC,
	3: model.FileD,
	4: model.FileE,
	5: model.FileF,
	6: model.FileG,
	7: model.FileH,
}

// TranslateRankFileToXY translates [model.File,model.Rank] coordinates to [x,y] coordinates
func TranslateRankFileToXY(coord model.Coord) (int, int) {
	return fileToX[coord.File], rankToY[coord.Rank]
}

// TranslateXYToRankFile translates [x,y] coordinates to [model.File,model.Rank] coordinates
func TranslateXYToRankFile(x, y int) model.Coord {
	return model.Coord{File: xToFile[x], Rank: yToRank[y]}
}

// Getstate.CoordsBySlopeAndDistance gets a list of coordinates (file,rank)
func GetCoordsBySlopeAndDistance(start model.Coord, yChange, xChange, distance int) []model.Coord {
	x, y := TranslateRankFileToXY(start)

	coords := []model.Coord{}

	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := model.ValidCoords[coord]
		if ok {
			coords = append(coords, coord)
		}
		y += yChange
		x += xChange
		d++
	}

	return coords
}

func GetCoordBySlopeAndDistance(start model.Coord, yChange, xChange int) (model.Coord, bool) {
	x, y := TranslateRankFileToXY(start)

	distance := 1
	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := model.ValidCoords[coord]
		if ok {
			return coord, true
		}
		y += yChange
		x += xChange
		d++
	}
	return model.Coord{}, false
}

func GetCoordsBySlopeAndDistanceAll(start model.Coord, distance int) []model.Coord {
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

	coords := []model.Coord{}

	for _, slope := range slopes {
		coords = append(coords, GetCoordsBySlopeAndDistance(start, slope[0], slope[1], distance)...)
	}

	return coords
}

// GetCoordByXY gets algebraic notation for a set of
// rank (y) and file (x) coordinates
func GetCoordByXY(
	squareOriginByCoords map[model.Coord][]float64,
	x, y float64,
) (model.Coord, bool) {
	for coord, xy := range squareOriginByCoords {
		if xy[0] == x && xy[1] == y {
			return coord, true
		}
	}
	return model.Coord{}, false
}
