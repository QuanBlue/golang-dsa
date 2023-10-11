package sort

// -------------------------------------------------------------- //
// Partition
// -------------------------------------------------------------- //
func partition(a []int) int {
		//   Start from the leftmost element and keep track of the index of 
		// smaller (or equal) elements as i. While traversing, if we find 
		// a smaller element, we swap the current element with arr[i]. 
		// Otherwise, we ignore the current element.
		// Ref: https://www.geeksforgeeks.org/quick-sort/
		
		low := 0
		high := len(a) - 1
		
		pivot := a[high]
		i := low - 1 // a[:i+1] always less than pivot

		for j := low; j < high; j++ {
			if a[j] < pivot {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}

		a[i+1], a[high] = a[high], a[i+1]
		
		return i + 1
}

// -------------------------------------------------------------- //
// Quick sort
// -------------------------------------------------------------- //
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot_idx := partition(arr)

	quickSort(arr[:pivot_idx]) // left side
	quickSort(arr[pivot_idx + 1:]) // left side

	return arr
}