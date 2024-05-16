package datastructures

import "fmt"

// WIKI https://en.wikipedia.org/wiki/Heap_(data_structure)
// https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// https://www.youtube.com/watch?v=3DYIgTC4T1o

type heap struct {
	isMin bool
	array []int
}

func NewHeap(isMin bool) *heap {
	return &heap{
		isMin: isMin,
		array: []int{},
	}
}

func (h *heap) Size() int {
	return len(h.array)
}

func (h *heap) Peek() int {
	if h.isMin {
		return -h.array[0]
	}

	return h.array[0]
}

func (h *heap) Insert(key int) {
	if h.isMin {
		key = -key
	}
	h.array = append(h.array, key)

	h.siftUp(len(h.array) - 1)
}

func (h *heap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("impossible to extract item from empty heap")
		return -1
	}

	extracted := h.array[0]

	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.siftDown(0)

	if h.isMin {
		return -extracted
	}

	return extracted
}

func (h *heap) siftUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
	}
}

func (h *heap) siftDown(index int) {
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
func (h *heap) Heapify(arr []int) {
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

func (h *heap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
