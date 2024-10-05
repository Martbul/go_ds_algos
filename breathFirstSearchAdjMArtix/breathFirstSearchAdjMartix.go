package main

import "fmt"

func main() {
	matrix2 := [][]int{
		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	source := 0 // the point where you are starting from
	needle := 6 // the vertex you are looking for

	path := bfs(matrix2, source, needle)
	fmt.Println(path)
}

func bfs(graph [][]int, source int, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))

	// Initialize `prev` with -1 to indicate unvisited nodes
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}

	for len(q) > 0 {
		var curr int
		curr, q = shift(q) // Get the first element and update queue

		if curr == needle {
			break
		}

		// Get adjacent nodes
		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue // Skip if no edge
			}

			if seen[i] {
				continue // Skip if already seen
			}

			// Mark the node as seen and record its predecessor
			seen[i] = true
			prev[i] = curr
			q = append(q, i) // Add to queue
		}
	}

	// Build the path by walking backwards using `prev`
	curr := needle
	out := []int{curr} // Start from the `needle`

	for prev[curr] != -1 {
		curr = prev[curr]
		out = append([]int{curr}, out...) // Prepend the predecessor to the path
	}

	// Check if the path is valid (i.e., reaches the source)
	if out[0] == source {
		return out
	}

	return []int{} // Return an empty array if no path found
}

// Helper functions

// shift removes the first element of a slice and returns it along with the updated slice
func shift(slice []int) (int, []int) {
	if len(slice) == 0 {
		return 0, slice
	}
	return slice[0], slice[1:]
}

// reverse reverses a slice in-place
func reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
