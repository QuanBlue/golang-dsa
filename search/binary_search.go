package search

// -------------------------------------------------------------- //
// Binary search
// -------------------------------------------------------------- //
func binarySearch(arr []int, key int) (position int) {
	start_idx := 0
	end_idx := len(arr) - 1
	
	for start_idx <= end_idx {
		mid := (start_idx + end_idx) / 2

		if arr[mid] == key {
			position = mid
			return
		} else if key > arr[mid] {
			start_idx = mid + 1
		} else {
			end_idx = mid - 1
		}	
	} 

	position = -1
	return
}