package sort

import "slices"

// -------------------------------------------------------------- //
// Radix Counting sort
// -------------------------------------------------------------- //
func rCountingSort(arr []int, exp int) []int {
	count_arr := make([]int, 10)
	
	// count
	for _, v := range arr {
		index := v / exp
		count_arr[index % 10]++
	}

	// cumulate sum
	for i := 1; i < 10; i++ {
		count_arr[i] += count_arr[i - 1]
	}
	
	// sort
	sorted_arr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] / exp

		sorted_arr[count_arr[index % 10] - 1] = arr[i]
		count_arr[index % 10]--
	}

	return sorted_arr
}

// -------------------------------------------------------------- //
// Radix sort
// -------------------------------------------------------------- //
func radixSort(arr []int) []int {
	max := slices.Max(arr)
	exp := 1

	for max / exp > 0 {
		arr = rCountingSort(arr, exp)
		exp *= 10
	}

	return arr
}