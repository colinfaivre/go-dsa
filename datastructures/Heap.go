// @GOSTDLIB https://pkg.go.dev/container/heap
// @WIKI https://en.wikipedia.org/wiki/Heap_(data_structure)
// @MEDIUM https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// @YOUTUBE https://www.youtube.com/watch?v=3DYIgTC4T1o

package datastructures

import (
	"fmt"
)

type Item struct {
	id    int
	Value int
}

type Heap struct {
	isMin bool
	array []Item
}

// O(1): Builds a minheap or max heap depending on isMin
func NewHeap(isMin bool) *Heap {
	return &Heap{
		isMin: isMin,
		array: []Item{},
	}
}

// O(1): Returns the sise of the heap
func (h *Heap) Size() int {
	return len(h.array)
}

// O(1): Returns the item at the top of the heap
func (h *Heap) Peek() Item {
	if h.isMin {
		return Item{
			id:    h.array[0].id,
			Value: -h.array[0].Value,
		}
	}

	return Item{
		id:    h.array[0].id,
		Value: h.array[0].Value,
	}
}

// O(logn): Inserts a new item into the heap
func (h *Heap) Insert(i Item) {
	if h.isMin {
		i.Value = -i.Value
	}
	h.array = append(h.array, i)

	h.siftUp(len(h.array) - 1)
}

// O(logn): Extracts the item on the top of the heap
func (h *Heap) ExtractFrom(index int) Item {
	if len(h.array) == 0 {
		// @TODO return an error instead of empty item?
		fmt.Println("impossible to extract item from empty heap")
		return Item{}
	}

	extracted := h.array[index]

	lastIndex := len(h.array) - 1
	h.array[index] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.siftDown(index)

	if h.isMin {
		return Item{
			id:    extracted.id,
			Value: -extracted.Value,
		}
	}

	return extracted
}

// O(n): Gets the index of the item with a given id
func (h *Heap) GetIndexFromId(id int) int {
	// @TODO ids need to be unique

	if len(h.array) == 0 {
		// @TODO return an error instead of -1?
		fmt.Println("impossible to extract item from empty heap")
		return -1
	}

	for i, item := range h.array {
		if item.id == id {
			return i
		}
	}

	// @TODO return an error instead of -1?
	fmt.Println("item not found")
	return -1
}

// Bubbles up an item to maintain the heap property
func (h *Heap) siftUp(index int) {
	for h.array[parent(index)].Value < h.array[index].Value {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Bubbles down an item to maintain the heap property
func (h *Heap) siftDown(index int) {
	lastIndex := len(h.array) - 1
	li, ri := left(index), right(index)
	childToCompareIdx := 0

	for li <= lastIndex {
		if li == lastIndex {
			childToCompareIdx = li
		} else if h.array[li].Value > h.array[ri].Value {
			childToCompareIdx = li
		} else {
			childToCompareIdx = ri
		}

		if h.array[index].Value < h.array[childToCompareIdx].Value {
			h.swap(index, childToCompareIdx)
			index = childToCompareIdx
			li, ri = left(index), right(index)
		} else {
			return
		}
	}

}

// O(nlogn): Inserts a batch of items into the heap
func (h *Heap) Heapify(arr []int) {
	for _, v := range arr {
		h.Insert(Item{Value: v})
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
