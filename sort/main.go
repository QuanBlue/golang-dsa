package sort

import (
	"fmt"
	"math/rand"
	"time"
)




func copyArray(arr []int) []int{
	new_arr := make([]int, len(arr))
	copy(new_arr, arr)

	return new_arr
}

func run(name string, sort_fn SortFunction, arr []int) {
	start := time.Now()
	sorted_arr := sort_fn(arr)
	exec_time := time.Since(start)
	
	fmt.Println(name, "- execute in", exec_time)
	fmt.Printf("Sorted array: %v\n\n", sorted_arr)
}


func generateArray(n int, limited int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(limited)
	}

	return arr
}


func Main(){
	fmt.Println("----------------------")
	fmt.Println("-        Sort        -")
	fmt.Println("----------------------")


	arr := generateArray(20, 100)
	fmt.Println("Generate array:", arr)


	fmt.Println("\n--------------------------")
	run("Bubble sort", bubbleSort, copyArray(arr))
	run("Selection sort", selectionSort, copyArray(arr))
	run("Insertion sort", insertionSort, copyArray(arr))
	run("Interchange sort", interchangeSort, copyArray(arr))
	run("Heap sort", heapSort, copyArray(arr))
	run("Quick sort", quickSort, copyArray(arr))
	run("Merge sort", mergeSort, copyArray(arr))
	run("Counting sort", countingSort, copyArray(arr))
	run("Radix sort", radixSort, copyArray(arr))

}