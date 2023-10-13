package graph

import (
	"fmt"
	"time"
)


func run(name string, search_fn SearchFunction, graph Graph, start_node, goal_node string) {
	start := time.Now()
	isFound, weight, path := search_fn(graph, start_node, goal_node)
	exec_time := time.Since(start)

	fmt.Println(">", name + " (execute in", exec_time.String() + "):") 
	fmt.Println("  -", goal_node, "is Found:", isFound)
	fmt.Println("  - Weight:", weight)
	fmt.Print("  - Path: ")

	for i, v := range path {
		fmt.Print(v.name)

		if i == len(path) - 1 {
			fmt.Print("\n\n")
		} else {
			fmt.Print(" -> ")
		}
	}

}

func Main() {
	fmt.Println("----------------------")
	fmt.Println("-       Graph       -")
	fmt.Println("----------------------")
	graph := new(Graph)

	graph.addEdge("A", "B", 0)
	graph.addEdge("A", "C", 0)
	
	graph.addEdge("B", "D", 0)
	graph.addEdge("B", "E", 0)
	graph.addEdge("B", "F", 0)

	graph.addEdge("C", "G", 0)

	graph.addEdge("F", "H", 0)

	graph.addEdge("G", "I", 0)

	graph.printVertices()


	run("BFS", breadthFirstSearch, *graph, "A", "F")
	run("BFS", breadthFirstSearch, *graph, "A", "I")
	run("BFS", breadthFirstSearch, *graph, "A", "X")


}