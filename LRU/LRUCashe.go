package main

import (
	"errors"
)

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type LRU[V any, K comparable] struct {
	length        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func (l *LRU[V, K]) update(key K, value V) {
	// Check for existence
	node, ok := l.lookup[key]
	if !ok {
		// Create a new node
		node = l.createNode(value)
		l.length++
		l.prepend(node)
		l.trimCache()

		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		// Move the node to the front of the list
		l.detach(node)
		l.prepend(node)
		node.value = value
	}
}

func (l *LRU[V, K]) get(key K) (V, error) {
	// Check for existence
	node, ok := l.lookup[key]
	if !ok {
		var zeroValue V
		return zeroValue, errors.New("key not found")
	}

	l.detach(node)  // Removing the node
	l.prepend(node) // Adding the node to the front

	return node.value, nil
}

func (l *LRU[V, K]) createNode(value V) *Node[V] {
	return &Node[V]{value: value}
}

func (l *LRU[V, K]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// If there's no previous node, we're at the head
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// If there's no next node, we're at the tail
		l.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (l *LRU[V, K]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
}

func (l *LRU[V, K]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}

func NewLRU[V any, K comparable](capacity int) *LRU[V, K] {
	return &LRU[V, K]{
		lookup:        make(map[K]*Node[V]),
		reverseLookup: make(map[*Node[V]]K),
		capacity:      capacity,
	}
}
