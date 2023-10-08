package search

import (
	"fmt"
)

func Main() {
	fmt.Println("----------------------")
	fmt.Println("-       Search       -")
	fmt.Println("----------------------")
	
	var n int
	fmt.Print("Input array length:")
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i:=0; i < n; i++ {
		fmt.Printf("> a[%d] = ", i)
		fmt.Scanln(&arr[i])
	}

	var key int
	fmt.Print("Input key you wanna search: ")
	fmt.Scanln(&key)

	// Linear search
	result, exec_time := linearSearch(arr, key)
	fmt.Println("\nLinear search",key,"at index",result,"in",exec_time)

	// Binary search
	result, exec_time = binarySearch(arr, key)
	fmt.Println("\nBinary search",key,"at index",result,"in",exec_time)
}