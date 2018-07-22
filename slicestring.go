package chess

// FindInSliceCoord finds a Coord in a slice
func FindInSliceCoord(slice []Coord, needle Coord) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == needle {
			return true
		}
	}
	return false
}
