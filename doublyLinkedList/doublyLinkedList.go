package main

import "fmt"

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (dll *DoublyLinkedList[T]) prepend(item T) {
	//Adding a new node to the star of the list
	node := &Node[T]{value: item}
	dll.length++
	if dll.head == nil {
		//if there is no head, just set the node as a head
		dll.head = node
		dll.tail = node
	}

	//if there is head,point the new node the the now second node, point the second node the the node you are adding, put the head of the dll to the new node
	node.next = dll.head // first point the new node to the next node(previos first node)
	dll.head.prev = node // second point the second node to the new node
	dll.head = node      // third set the list's head to the new node
}

func (dll *DoublyLinkedList[T]) insertAt(item T, index int) {

	if index > dll.length {
		fmt.Println("ERROR", "Cannot insert out fo the list")
	}

	if index == dll.length {
		dll.append(item)
	} else if index == 0 {
		dll.prepend(item)
	}

	dll.length++
	curr := dll.head

	for i := 0; i < index; i++ {
		curr = curr.next
	}

	node := &Node[T]{value: item}
	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if curr.prev != nil {
		curr.prev.next = curr
	}

}

func (dll *DoublyLinkedList[T]) append(item T) {
	node := &Node[T]{value: item}
	dll.length++

	if dll.tail == nil {
		dll.head = node
		dll.tail = node
	}

	node.prev = dll.tail // make the new node to point the the last node(now second to last node)
	dll.tail.next = node // update the last node to point ot he new node(the new last node)
	dll.tail = node      //update the tail to point to the new node

}

// FIXME: remove method doesnt work
func (dll *DoublyLinkedList[T]) remove(item T) T {
	curr := dll.head
	// Traverse the list to find the node with the value to be removed
	for curr != nil {
		if compare(curr.value, item) {
			break
		}
		curr = curr.next
	}

	// If the item wasn't found, return the zero value of T
	if curr == nil {
		var zeroValue T
		return zeroValue
	}

	// Decrease the length of the list
	dll.length--

	// If the list becomes empty after removal
	if dll.length == 0 {
		out := dll.head.value
		dll.head = nil
		dll.tail = nil
		return out
	}

	// If the current node is not the head, update the previous node's next pointer
	if curr.prev != nil {
		curr.prev.next = curr.next
	} else {
		// If curr is the head, move the head to the next node
		dll.head = curr.next
	}

	// If the current node is not the tail, update the next node's prev pointer
	if curr.next != nil {
		curr.next.prev = curr.prev
	} else {
		// If curr is the tail, move the tail to the previous node
		dll.tail = curr.prev
	}

	// Clear the removed node's pointers
	curr.prev = nil
	curr.next = nil

	// Return the value of the removed node
	return curr.value
}

func (dll *DoublyLinkedList[T]) get(index int) T {
	node := dll.getAt(index)
	return node.value
}

func (dll *DoublyLinkedList[T]) removeAt(index int) {}

// TODO: Use getAt method in the above method (the logic is the same)
func (dll *DoublyLinkedList[T]) getAt(index int) *Node[T] {

}
