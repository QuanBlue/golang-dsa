package graph

import (
	"slices"
)

// -------------------------------------------------------------- //
// Breadth First Search (BFS)
// -------------------------------------------------------------- //
func breadthFirstSearch(g Graph, start_node, goal_node string) (bool, int, []Vertex) {
	visited := make([]string, 0)
	frontier := make([]GraphTraversal, 0)

	// add start node to the queue
	traversal := GraphTraversal{
		vertex: g.vertexes[start_node],
		totalWeight: 0,
		path: []Vertex{g.vertexes[start_node]},
	}
	frontier = append(frontier, traversal)

	// while queue not empty -> keep traversal
	for len(frontier) > 0 {
		// get first vertex in queue by pop vertex in frontier
		current_node := frontier[0]
		frontier = frontier[1:]

		// if found vertex -> return
		if current_node.vertex.name == goal_node {
			return true, current_node.totalWeight, current_node.path
		}

		// add all child node of current node to frontier
		if !slices.Contains(visited, current_node.vertex.name) {
			// add current vertex in visited 
			visited = append(visited, current_node.vertex.name)
	
			for _, child := range current_node.vertex.edges {
				traversal := GraphTraversal{
					vertex: child.vertex,
					totalWeight: current_node.totalWeight + child.weight, 
					path: append(current_node.path, child.vertex),
				}
				frontier = append(frontier, traversal)
			}
		}
	}

	return false, 0, make([]Vertex,0)
}