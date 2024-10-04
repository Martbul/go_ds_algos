package main

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func main() {
	path := []int{}
	// Construct the tree
	tree := &BinaryNode[int]{
		Value: 20,
		Right: &BinaryNode[int]{
			Value: 50,
			Right: &BinaryNode[int]{
				Value: 100,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode[int]{
				Value: 30,
				Right: &BinaryNode[int]{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
				Left: &BinaryNode[int]{
					Value: 29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &BinaryNode[int]{
			Value: 10,
			Right: &BinaryNode[int]{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode[int]{
				Value: 5,
				Right: &BinaryNode[int]{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	result := post0rderSearch(tree, path)
	fmt.Println(result)
}

func post0rderSearch(head *BinaryNode[int], path []int) []int {
	return walk(head, path)
}

func walk(curr *BinaryNode[int], path []int) []int {
	// the base case is when a node is a leaf(doesnt have chiled so you cannot continue to recurse)
	if curr == nil {
		return path
	}

	// pre

	//recurse
	path = walk(curr.Left, path)

	path = walk(curr.Right, path)

	//post
	path = append(path, curr.Value)

	//no need for post recursion here

	return path
}
