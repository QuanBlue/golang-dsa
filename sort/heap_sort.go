package sort

// -------------------------------------------------------------- //
// Heapify
// > max-heap
// -------------------------------------------------------------- //
func heapify(a []int, len, i int) {
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

// -------------------------------------------------------------- //
// Heap sort
// -------------------------------------------------------------- //
func heapSort(arr []int) []int {
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