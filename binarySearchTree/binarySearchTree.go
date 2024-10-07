//INFO: Rule for Binary Search Tree - everything on the left is smaller than the parant node and everything on the rigth is bigger than the parant node

package main

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func main() {

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

	needle := 10
	result := depthFirstSearch(tree, needle)
	fmt.Println(result)
}

func depthFirstSearch(head *BinaryNode[int], needle int) bool {
	return search(head, needle)
}

func search(curr *BinaryNode[int], needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		// go to the right side of the tree
		return search(curr.Right, needle)
	} else {
		// go to the left side of the tree
		return search(curr.Left, needle)
	}
}
