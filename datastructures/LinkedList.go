package datastructures

import (
	"fmt"
	"strconv"
)

// @GOSTDLIB https://pkg.go.dev/container/list
// @WIKI https://en.wikipedia.org/wiki/Linked_list
// @YOUTUBE https://www.youtube.com/watch?v=1S0_-VxPLJo

type LinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	data int
	next *node
}

// Insert the given list of elements at the front of the linked list
func (l *LinkedList) AddAllFront(nl []int) {
	for i := len(nl) - 1; i >= 0; i-- {
		l.AddFirst(nl[i])
	}
}

// Insert the given element at the front of the linked list
func (l *LinkedList) AddFirst(n int) {
	node := &node{data: n, next: nil}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.size = 1
	} else {
		temp := l.head
		l.head = node
		l.head.next = temp
		l.size++
	}
}

// Retrieves, but does not remove the first element of the linked list
func (l *LinkedList) PeekFirst() int {
	if l.isEmpty() {
		fmt.Println("The linked list is empty!")
	}

	return l.head.data
}

// Removes and returns the first item of the linked list
func (l *LinkedList) RemoveFirst() int {
	if l.head == nil {
		fmt.Println("Nothing to remove!")
	}

	removed := l.head.data

	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.size = 0
	} else {
		l.head = l.head.next
		l.size--
	}

	return removed
}

// Insert the given element at the back of the linked list
func (l *LinkedList) AddLast(n int) {
	node := &node{data: n, next: nil}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.size = 1
	} else {
		l.tail.next = node
		l.tail = node
		l.size++
	}
}

// Retrieves, but does not remove the last element of the linked list
func (l *LinkedList) PeekLast() int {
	if l.isEmpty() {
		fmt.Println("The linked list is empty!")
	}

	return l.tail.data
}

// Returns the size of the linked list
func (l *LinkedList) Size() int {
	return l.size
}

// Returns true if the linked list is empty
func (l *LinkedList) isEmpty() bool {
	return l.size == 0
}

// Represents the linked list as a string
func (l LinkedList) ToString() string {
	var result string
	curr_node := l.head

	for {
		result = result + strconv.Itoa(curr_node.data)
		if curr_node.next != nil {
			curr_node = curr_node.next
		} else {
			break
		}
	}

	return result
}
