package sort

import "slices"

// -------------------------------------------------------------- //
// Counting sort
// -------------------------------------------------------------- //
func countingSort(arr []int) []int {
	min := slices.Min(arr)
	max := slices.Max(arr)
	element_range := max - min + 1

	// count
	count_arr := make([]int, element_range)
	for _, v := range arr {
		count_arr[v - min]++
	}

	// cumulative count
	for i := 1; i < element_range; i++ {
		count_arr[i] += count_arr[i - 1]
	}

	// sort
	sorted_arr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sorted_arr[count_arr[arr[i] - min] - 1] = arr[i]
		count_arr[arr[i] - min]--
	}

	return sorted_arr
}