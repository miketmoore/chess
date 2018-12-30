package chess

var validCoords = map[Coord]bool{
	Coord{Rank1, FileA}: true,
	Coord{Rank2, FileA}: true,
	Coord{Rank3, FileA}: true,
	Coord{Rank4, FileA}: true,
	Coord{Rank5, FileA}: true,
	Coord{Rank6, FileA}: true,
	Coord{Rank7, FileA}: true,
	Coord{Rank8, FileA}: true,

	Coord{Rank1, FileB}: true,
	Coord{Rank2, FileB}: true,
	Coord{Rank3, FileB}: true,
	Coord{Rank4, FileB}: true,
	Coord{Rank5, FileB}: true,
	Coord{Rank6, FileB}: true,
	Coord{Rank7, FileB}: true,
	Coord{Rank8, FileB}: true,

	Coord{Rank1, FileC}: true,
	Coord{Rank2, FileC}: true,
	Coord{Rank3, FileC}: true,
	Coord{Rank4, FileC}: true,
	Coord{Rank5, FileC}: true,
	Coord{Rank6, FileC}: true,
	Coord{Rank7, FileC}: true,
	Coord{Rank8, FileC}: true,

	Coord{Rank1, FileD}: true,
	Coord{Rank2, FileD}: true,
	Coord{Rank3, FileD}: true,
	Coord{Rank4, FileD}: true,
	Coord{Rank5, FileD}: true,
	Coord{Rank6, FileD}: true,
	Coord{Rank7, FileD}: true,
	Coord{Rank8, FileD}: true,

	Coord{Rank1, FileE}: true,
	Coord{Rank2, FileE}: true,
	Coord{Rank3, FileE}: true,
	Coord{Rank4, FileE}: true,
	Coord{Rank5, FileE}: true,
	Coord{Rank6, FileE}: true,
	Coord{Rank7, FileE}: true,
	Coord{Rank8, FileE}: true,

	Coord{Rank1, FileF}: true,
	Coord{Rank2, FileF}: true,
	Coord{Rank3, FileF}: true,
	Coord{Rank4, FileF}: true,
	Coord{Rank5, FileF}: true,
	Coord{Rank6, FileF}: true,
	Coord{Rank7, FileF}: true,
	Coord{Rank8, FileF}: true,

	Coord{Rank1, FileG}: true,
	Coord{Rank2, FileG}: true,
	Coord{Rank3, FileG}: true,
	Coord{Rank4, FileG}: true,
	Coord{Rank5, FileG}: true,
	Coord{Rank6, FileG}: true,
	Coord{Rank7, FileG}: true,
	Coord{Rank8, FileG}: true,

	Coord{Rank1, FileH}: true,
	Coord{Rank2, FileH}: true,
	Coord{Rank3, FileH}: true,
	Coord{Rank4, FileH}: true,
	Coord{Rank5, FileH}: true,
	Coord{Rank6, FileH}: true,
	Coord{Rank7, FileH}: true,
	Coord{Rank8, FileH}: true,
}

var rankToY = map[Rank]int{
	Rank1: 0,
	Rank2: 1,
	Rank3: 2,
	Rank4: 3,
	Rank5: 4,
	Rank6: 5,
	Rank7: 6,
	Rank8: 7,
}

var yToRank = map[int]Rank{
	0: Rank1,
	1: Rank2,
	2: Rank3,
	3: Rank4,
	4: Rank5,
	5: Rank6,
	6: Rank7,
	7: Rank8,
}

var fileToX = map[File]int{
	FileA: 0,
	FileB: 1,
	FileC: 2,
	FileD: 3,
	FileE: 4,
	FileF: 5,
	FileG: 6,
	FileH: 7,
}

var xToFile = map[int]File{
	0: FileA,
	1: FileB,
	2: FileC,
	3: FileD,
	4: FileE,
	5: FileF,
	6: FileG,
	7: FileH,
}

func TranslateRankFileToXY(coord Coord) (int, int) {
	return fileToX[coord.File], rankToY[coord.Rank]
}

func TranslateXYToRankFile(x, y int) Coord {
	return Coord{File: xToFile[x], Rank: yToRank[y]}
}

// GetCoordsBySlopeAndDistance gets a list of coordinates (rank,file)
func GetCoordsBySlopeAndDistance(start Coord, yChange, xChange, distance int) []Coord {
	x, y := TranslateRankFileToXY(start)

	coords := []Coord{}

	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := validCoords[coord]
		if ok {
			coords = append(coords, coord)
		}
		y += yChange
		x += xChange
		d++
	}

	return coords
}

func GetCoordBySlopeAndDistance(start Coord, yChange, xChange int) (Coord, bool) {
	x, y := TranslateRankFileToXY(start)

	distance := 1
	d := 0

	y += yChange
	x += xChange

	for d < distance && y < 8 && y >= 0 && x < 8 && x >= 0 {
		coord := TranslateXYToRankFile(x, y)
		_, ok := validCoords[coord]
		if ok {
			return coord, true
		}
		y += yChange
		x += xChange
		d++
	}
	return Coord{}, false
}

func GetCoordsBySlopeAndDistanceAll(start Coord, distance int) []Coord {
	slopes := [][]int{
		[]int{1, 0},   // n
		[]int{1, 1},   // ne
		[]int{0, 1},   // e
		[]int{-1, 1},  // se
		[]int{-1, 0},  // s
		[]int{-1, -1}, // sw
		[]int{0, -1},  // w
		[]int{1, -1},  // nw
	}

	coords := []Coord{}

	for _, slope := range slopes {
		coords = append(coords, GetCoordsBySlopeAndDistance(start, slope[0], slope[1], distance)...)
	}

	return coords
}

func isOccupied(boardState BoardState, coord Coord) bool {
	_, isOccupied := boardState[coord]
	return isOccupied
}

func isOccupiedByColor(boardState BoardState, coord Coord, color PlayerColor) bool {
	occupant, occupied := boardState[coord]
	return occupied && occupant.Color == color
}
