// Direction Graph

package graph

import (
	"fmt"
	"strconv"
)

type SearchFunction func(Graph, string, string) (bool, int, []Vertex)

// Graph Traversal represents current vertex 
// with path and cost from start vertex to it, 
type GraphTraversal struct {
	vertex Vertex
	totalWeight int
	path []Vertex
}

// Graph represents a set of vertices connected by edges.
type Graph struct {
	vertexes map[string]Vertex
}

// Vertex is a node in the graph that stores the int value at that node
// along with a map to the vertices it is connected to via edges.
type Vertex struct {
	name string
	edges map[string]Edge
}

// Edge holds the data bout the connection between two vertices
// Edge represents an edge in the graph and the destination vertex.
type Edge struct {
	weight int
	vertex Vertex
}

// -------------------------------------------------------------- //
// Add vertex
// -------------------------------------------------------------- //
func (g *Graph) addVertex(name string) {
	// create new vertex with initial name
	newVertex := Vertex{
		name: name,
		edges: make(map[string]Edge),
	}

	g.vertexes[name] = newVertex
}

// -------------------------------------------------------------- //
// Add edge
// -------------------------------------------------------------- //
func (g *Graph) addEdge(srcVertexName, destVertexName string, weight int) {
	if g.vertexes == nil {
		g.vertexes = make(map[string]Vertex)
	}

	// check if source and destination vertex exist
	// if not, create it
	if _, ok := g.vertexes[srcVertexName]; !ok {
		g.addVertex(srcVertexName)
	}
	if _, ok := g.vertexes[destVertexName]; !ok {
		g.addVertex(destVertexName)
	}

	// create edge from srcVertex to destVertex
	destVertex := g.vertexes[destVertexName]
	edge := Edge{
		weight: weight,
		vertex: destVertex,
	}
	g.vertexes[srcVertexName].edges[destVertexName] = edge

}


func (g *Graph) printVertices() {
	fmt.Print("Graph:")
	for _, vertex := range g.vertexes {
		if len(vertex.edges) != 0 {
			fmt.Print("\n  " + vertex.name + ": ")

			for _, v := range vertex.edges {
				fmt.Print("(" + v.vertex.name + ", " + strconv.Itoa(v.weight) +  ") ")
			}
		}
	}
	fmt.Println("\n---------------")
}