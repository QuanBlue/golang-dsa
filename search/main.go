package search

import (
	"fmt"
	"time"
)

type SearchFunction func([]int, int) (int)

func run(name string, search_fn SearchFunction, arr []int, key int) {
	start := time.Now()
	result := search_fn(arr, key)
	exec_time := time.Since(start)
	
	fmt.Println(name+":","Search", key,"at index",result,", execute in",exec_time)
}

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

	fmt.Println("\n--------------------------")
	run("Linear search", linearSearch, arr, key)
	run("Binary search", binarySearch, arr, key)
}