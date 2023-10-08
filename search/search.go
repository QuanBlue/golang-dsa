package search

import "time"


// linear search
func linearSearch(arr []int, key int) (position int, exec_time time.Duration) {
	start := time.Now()
	
	for idx, val := range arr {
		if val == key {
			position = idx
			exec_time = time.Since(start)
			return 
		}
	}

	position = -1
	exec_time = time.Since(start)
	return
}

// binary search
func binarySearch(arr []int, key int) (position int, exec_time time.Duration) {
	start := time.Now()
	
	arr_len := len(arr)
	start_idx := 0
	end_idx := arr_len - 1
	
	for start_idx <= end_idx {
		mid := (start_idx + end_idx) / 2

		if arr[mid] == key {
			position = mid
			exec_time = time.Since(start)
			return
		} else if key > arr[mid] {
			start_idx = mid + 1
		} else {
			end_idx = mid - 1
		}	
	} 

	position = -1
	exec_time = time.Since(start)
	return
}
