package datastructures

import "fmt"

// WIKI https://en.wikipedia.org/wiki/Heap_(data_structure)
// https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// https://www.youtube.com/watch?v=3DYIgTC4T1o

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) GetArray() []int {
	return h.array
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)

	h.HeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract(key int) int {
	if len(h.array) == 0 {
		fmt.Println("impossible to extract item from empty Heap")
		return -1
	}

	extracted := h.array[0]

	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.HeapifyDown(0)

	return extracted
}

func (h *MaxHeap) HeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
	}
}

func (h *MaxHeap) HeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

// n batched inserts: O(n)
func (h *MaxHeap) Heapify(arr []int) {
	for _, v := range arr {
		h.Insert(v)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
