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

func GetCoordsBySlope(start Coord, xChange, yChange int) []Coord {
	fmt.Println(yChange, xChange)
	coords := []Coord{}
	x, y := TranslateRankFileToXY(start)
	fmt.Println("translated x,y: ", x, y)
	x++
	y++
	for x <= 7 && y <= 7 {
		fmt.Println("LOOP")
		coord := TranslateXYToRankFile(x, y)
		_, ok := validCoords[coord]
		if ok {
			// fmt.Printf("translated coord: %+v\n", coord)
			coords = append(coords, coord)

		}
		// fmt.Println("x...", x, xChange)
		// fmt.Println("y...", y, yChange)
		x += xChange
		y += yChange
		fmt.Println("x,y: ", x, y)
	}
	return coords
}

// GetCoordsBySlopeAndDistance gets a list of coordinates (rank,file)
func GetCoordsBySlopeAndDistance(start Coord, yChange, xChange, distance int) []Coord {
	x, y := TranslateRankFileToXY(start)
	fmt.Println(x, y)
	coords := []Coord{}

	d := 0
	if yChange == 1 && xChange == 0 {
		y++
		for d < distance && y < 8 {
			coord := TranslateXYToRankFile(x, y)
			_, ok := validCoords[coord]
			if ok {
				coords = append(coords, coord)
			}
			y++
			d++
		}
	} else if yChange == -1 && xChange == 0 {
		y--
		for d < distance && y >= 0 {
			coord := TranslateXYToRankFile(x, y)
			_, ok := validCoords[coord]
			if ok {
				coords = append(coords, coord)
			}
			y--
			d++
		}
	} else if yChange == 1 && xChange == 1 {
		y++
		x++
		for d < distance && y < 8 && x < 8 {
			coord := TranslateXYToRankFile(x, y)
			_, ok := validCoords[coord]
			if ok {
				coords = append(coords, coord)
			}
			y++
			x++
			d++
		}
	} else if yChange == 1 && xChange == -1 {
		y++
		x--
		for d < distance && y < 8 && x > 0 {
			coord := TranslateXYToRankFile(x, y)
			_, ok := validCoords[coord]
			if ok {
				coords = append(coords, coord)
			}
			y++
			x--
			d++
		}
	}

	return coords
}
