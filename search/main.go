package search

import (
	"fmt"
	"math/rand"
	"time"
)

func run(name string, search_fn SearchFunction, arr []int, key int) {
	start := time.Now()
	result := search_fn(arr, key)
	exec_time := time.Since(start)
	
	fmt.Println(">", name + " (execute in", exec_time.String() + "): index", result)
}


func generateArray(n int, limited int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(limited)
	}

	return arr
}

func Main() {
	fmt.Println("----------------------")
	fmt.Println("-       Search       -")
	fmt.Println("----------------------")
	
	arr := generateArray(10, 100)
	fmt.Println("Generate array:", arr)

	var key int
	fmt.Print("Input key you wanna search: ")
	fmt.Scanln(&key)

	fmt.Println("\n--------------------------")
	fmt.Println("Search:", key)
	run("Linear search", linearSearch, arr, key)
	run("Binary search", binarySearch, arr, key)
}