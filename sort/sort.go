package sort

import (
	"slices"
)

// Bubble sort
func bubbleSort(arr []int) []int{
	for i := 1; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// Selection sort
func selectionSort(arr []int) []int{
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

// Insertion sort
func insertionSort(arr []int) []int{
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j--
		}

		arr[j + 1] = key
	}

	return arr
}

// Interchange sort
func interchangeSort(arr []int) []int{
	for i := 0; i < len(arr) - 1; i++ {
		for j :=i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}

// Heap sort
func heapSort(arr []int) []int{
	// max-heap
	var heapify func(a []int, len, i int)
	heapify = func(a []int, len, i int) {
		largest := i
		left_child := i * 2 + 1
		right_child := i * 2 + 2

		if left_child < len && a[left_child] > a[largest] {
			largest = left_child
		} 

		if right_child < len && a[right_child] > a[largest] {
			largest = right_child
		} 

		if largest != i {
			a[largest], a[i] = a[i], a[largest]
			heapify(a, len, largest)
		}
	}

	// Create heap tree
	for i := len(arr)/2-1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	// Sort
	for i := len(arr)-1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

// Quick sort
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	partition := func(a []int) int{
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
			if arr[j] < pivot {
				i++

				arr[i], arr[j] = arr[j], arr[i]
			}
		}

		arr[i+1], arr[high] = arr[high], arr[i+1]
		
		return i + 1
	}

	pivot_idx := partition(arr)

	quickSort(arr[:pivot_idx]) // left side
	quickSort(arr[pivot_idx + 1:]) // left side

	return arr
}

// Merge sort
func mergeSort(arr []int) []int{
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])

		// merge left and right array to a single sorted array
		merge := func() []int {
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

		return merge()
	}

	return arr
}

// Counting sort
func countingSort(arr []int) []int{
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

// Radix sort
func radixSort(arr []int) []int {
	// counting sort
	countingSort := func (arr []int, exp int) []int{
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

	// radix sort
	max := slices.Max(arr)
	exp := 1

	for max / exp > 0 {
		arr = countingSort(arr, exp)
		exp *= 10
	}

	return arr
}