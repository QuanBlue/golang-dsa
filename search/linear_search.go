package search

// -------------------------------------------------------------- //
// Linear search
// -------------------------------------------------------------- //
func linearSearch(arr []int, key int) (position int) {
	for idx, val := range arr {
		if val == key {
			position = idx
			return 
		}
	}

	position = -1
	return
}
