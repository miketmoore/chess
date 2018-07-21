package chess

// FindInSliceString finds a string in a slice of strings
func FindInSliceString(slice []string, needle string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == needle {
			return true
		}
	}
	return false
}
