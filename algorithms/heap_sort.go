package algorithms

import "github.com/colinfaivre/go-dsa/datastructures"

func HeapSort(arr []int) []int {
	heap := datastructures.NewHeap(true)
	heap.Heapify(arr)
	new_arr := []int{}

	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, heap.ExtractTop())
	}

	return new_arr
}
