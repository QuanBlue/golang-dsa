package sort

// -------------------------------------------------------------- //
// Selection sort
// -------------------------------------------------------------- //
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		min_idx := i
		
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		
		arr[i], arr[min_idx] = arr[min_idx], arr[i]
	}

	return arr
}
