package main

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func main() {

	//INFO: Dont think of the three as a JS object. When it is being passed to the btsearch func it basicaly passes the value 20 and 2 pointer(leaf and rigth)
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

	needle := 7

	res := breadthFirstSearch(tree, needle)
	fmt.Println(res)

}

func breadthFirstSearch(head *BinaryNode[int], needle int) bool {

	q := []*BinaryNode[int]{head}

	for len(q) > 0 {
		var curr *BinaryNode[int]
		//extracting the value of the curr and removing it from the queue
		curr, q = shift(q)

		//search
		if curr.Value == needle {
			return true
		}

		if curr.Left != nil {

			q = append(q, curr.Left)
		}

		if curr.Right != nil {

			q = append(q, curr.Right)
		}

	}

	return false
}

// Helper function that simulates JavaScript's shift() method
func shift[T any](queue []*BinaryNode[T]) (*BinaryNode[T], []*BinaryNode[T]) {
	// Return the first element and the rest of the slice
	return queue[0], queue[1:]
}
