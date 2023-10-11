package sort

// -------------------------------------------------------------- //
// Bubble sort
// -------------------------------------------------------------- //
func bubbleSort(arr []int) []int {
	for i := 1; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}