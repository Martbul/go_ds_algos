package main

import "fmt"

// Edge structure to hold the destination and weight
type Edge struct {
	To     int // Destination node
	Weight int // Edge weight
}

// WeightedAdjacencyList represents the adjacency list
type WeightedAdjacencyList [][]Edge

func main() {
	// Initialize the weighted adjacency list with empty slices for each node
	list2 := WeightedAdjacencyList{
		{}, // Node 0
		{}, // Node 1
		{}, // Node 2
		{}, // Node 3
		{}, // Node 4
		{}, // Node 5
		{}, // Node 6
	}

	// Populate the adjacency list based on the provided graph structure
	list2[0] = []Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	list2[1] = []Edge{
		{To: 4, Weight: 1},
	}
	list2[2] = []Edge{
		{To: 3, Weight: 7},
	}
	list2[3] = []Edge{} // No outgoing edges
	list2[4] = []Edge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	list2[5] = []Edge{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	}
	list2[6] = []Edge{
		{To: 3, Weight: 1},
	}

	source := 0 // The point where you are starting from
	needle := 6 // The vertex you are looking for

	path := dfs(list2, source, needle)
	fmt.Println(path)
}

func dfs(graph WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}
	walk(graph, source, needle, seen, &path) // Pass path by reference
	return path
}

func walk(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true
	// Pre-order: add current node to the path
	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	// Recur for each adjacent edge
	list := graph[curr]
	for _, edge := range list {
		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	// Post-order: remove the current node from the path if the needle is not found
	*path = pop(*path)

	return false
}

// Pop removes and returns the last element of a slice
func pop(slice []int) []int {
	if len(slice) == 0 {
		return slice // Handle empty slice case
	}
	return slice[:len(slice)-1] // Return the slice without the last element
}
