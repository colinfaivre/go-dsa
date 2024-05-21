// @WIKI https://en.wikipedia.org/wiki/Heap_(data_structure)
// @MEDIUM https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// @YOUTUBE https://www.youtube.com/watch?v=3DYIgTC4T1o

package datastructures

import "fmt"

type Heap struct {
	isMin bool
	array []int
}

// O(1): Builds a minheap or max heap depending on isMin
func NewHeap(isMin bool) *Heap {
	return &Heap{
		isMin: isMin,
		array: []int{},
	}
}

// O(1): Returns the sise of the heap
func (h *Heap) Size() int {
	return len(h.array)
}

// O(1): Returns the value at the top of the heap
func (h *Heap) Peek() int {
	if h.isMin {
		return -h.array[0]
	}

	return h.array[0]
}

// O(logn): Inserts a new value into the heap
func (h *Heap) Insert(key int) {
	if h.isMin {
		key = -key
	}
	h.array = append(h.array, key)

	h.siftUp(len(h.array) - 1)
}

// O(logn): Extracts the value on the top of the heap
func (h *Heap) ExtractTop() int {
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

// O(logn): Extracts a value from the heap
func (h *Heap) Extract() int {

	return 0
}

// Bubbles up a value to maintain the heap property
func (h *Heap) siftUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Bubbles down a value to maintain the heap property
func (h *Heap) siftDown(index int) {
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

// O(nlogn): Inserts a batch of values into the heap
func (h *Heap) Heapify(arr []int) {
	for _, v := range arr {
		h.Insert(v)
	}
}

// O(1): Swaps to values in the heap
func (h *Heap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
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
