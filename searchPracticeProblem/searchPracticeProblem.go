//INFO: For comparing 2 trees io value and structure must be used Depth First Search

package main

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func main() {

	tree1 := &BinaryNode[int]{
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

	tree2 := &BinaryNode[int]{
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

	result := compareBinaryTree(tree1, tree2)
	fmt.Println(result)
}

func compareBinaryTree(t1Node *BinaryNode[int], t2Node *BinaryNode[int]) bool {
	//checking the structure of the trees(if both hit a node with no children then they are similat)
	if t1Node == nil && t2Node == nil {
		return true
	}

	if t1Node == nil || t2Node == nil {
		return false
	}

	//checking the value in the node in the trees
	if t1Node.Value != t2Node.Value {
		return false
	}

	return compareBinaryTree(t1Node.Left, t2Node.Left) && compareBinaryTree(t1Node.Right, t2Node.Right)

}
