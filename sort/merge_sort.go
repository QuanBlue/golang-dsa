package sort

// -------------------------------------------------------------- //
// Merge
// > merge left and right array to a single sorted array
// -------------------------------------------------------------- //
func merge(left, right []int) []int {
	l, r := 0, 0
	var sorted_arr []int

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			sorted_arr = append(sorted_arr, right[r])
			r++
		} else {
			sorted_arr = append(sorted_arr, left[l])
			l++
		}
	}

	// Append any remaining elements from both slices (if any)
	sorted_arr = append(sorted_arr, left[l:]...)
	sorted_arr = append(sorted_arr, right[r:]...)

	return sorted_arr
}

// -------------------------------------------------------------- //
// Merge sort
// -------------------------------------------------------------- //
func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])

		return merge(left, right)
	}

	return arr
}
