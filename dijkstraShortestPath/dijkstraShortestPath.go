package main

import (
	"fmt"
	"math"
)

// Edge represents an edge in the graph with a target node and a weight
type Edge struct {
	To     int // The target node
	Weight int // The weight of the edge
}

// WeightedAdjacencyList represents the graph as a list of edges for each node
type WeightedAdjacencyList [][]Edge

func main() {

	// Initialize the weighted adjacency list with empty slices for each node
	list1 := make(WeightedAdjacencyList, 7)

	// Populate the adjacency list
	list1[0] = []Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	list1[1] = []Edge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}
	list1[2] = []Edge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	}
	list1[3] = []Edge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	}
	list1[4] = []Edge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	list1[5] = []Edge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	}
	list1[6] = []Edge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	}

	source := 0
	needle := 6
	res := dijkstrShortestPath(&list1, source, needle)
	fmt.Println(res)
}

func dijkstrShortestPath(list1 *WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(*list1))

	prev := make([]int, len(*list1))

	// Initialize `prev` with -1 to indicate unvisited nodes
	for i := range prev {
		prev[i] = -1
	}

	// Create a slice to hold distances with the same length as arr
	dists := make([]float64, len(*list1))

	// Fill the dists slice with positive infinity
	for i := range dists {
		dists[i] = math.Inf(1) // math.Inf(1) represents positive infinity
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)

		seen[curr] = true

		adjs := (*list1)[curr] //list of our endges
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + float64(edge.Weight)
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := []int{}
	curr := needle

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	return reverseInts(out)

}

func hasUnvisited(seen []bool, dists []float64) bool {
	for i, s := range seen {
		// If the current element of seen is false and dists[i] is less than positive infinity
		if !s && dists[i] < math.Inf(1) {
			return true
		}
	}
	// If no element matches the condition, return false
	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	index := -1
	lowestDistanse := math.Inf(1)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistanse > float64(dists[i]) {
			lowestDistanse = float64(dists[i])
			index = i
		}
	}

	return index
}

// Reverse a slice of integers and return a new slice
func reverseInts(arr []int) []int {
	result := make([]int, len(arr)) // Create a new slice to hold the result
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		result[i], result[j] = arr[j], arr[i] // Assign swapped elements
	}
	return result
}
