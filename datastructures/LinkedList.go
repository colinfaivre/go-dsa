package datastructures

import (
	"fmt"
	"strconv"
)

// @WIKI https://en.wikipedia.org/wiki/Linked_list
// @YOUTUBE https://www.youtube.com/watch?v=1S0_-VxPLJo

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Data int
	next *Node
}

func (l *LinkedList) AddAll(nl []*Node) {}

func (l *LinkedList) AddFirst(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.size = 1
	} else {
		temp := l.head
		l.head = n
		l.head.next = temp
		l.size++
	}
}

func (l *LinkedList) RemoveFirst() int {
	if l.head == nil {
		fmt.Println("Nothing to remove!")
	}

	removed := l.head.Data

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

func (l *LinkedList) AddLast() {}

func (l *LinkedList) RemoveLast() int {
	return 0
}

func (l LinkedList) ToString() string {
	var result string
	curr_node := l.head

	for {
		result = result + strconv.Itoa(curr_node.Data)
		if curr_node.next != nil {
			curr_node = curr_node.next
		} else {
			break
		}
	}

	return result
}
