package chess

import "fmt"

/*
	(x,y)
	7
	6
	5
	4
	3
	2
	1  0,1  1,1  2,1  3,1  4,1  5,1  6,1  7,1
	0  0,0  1,0  2,0  3,0  4,0  5,0  6,0  7,0
	   0    1    2    3    4    5    6    7
*/
func TranslateRankFileToXY(coord Coord) (int, int) {
	y := int(coord.Rank) - 1
	x := int(coord.File) - 1
	return x, y
}

func TranslateXYToRankFile(x, y int) Coord {
	file := File(x + 1)
	rank := Rank(y + 1)
	return NewCoord(file, rank)
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
		fmt.Println(x, y)
		_, ok := validCoords[coord]
		fmt.Println("OK")
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
		fmt.Println(x, y)
		_, ok := validCoords[coord]
		fmt.Println("OK")
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
