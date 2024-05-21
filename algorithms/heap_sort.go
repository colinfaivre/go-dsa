package algorithms

import "github.com/colinfaivre/go-dsa/datastructures"

// @GOSTDLIB https://pkg.go.dev/sort

func HeapSort(arr []int) []int {
	heap := datastructures.NewHeap(true)
	heap.Heapify(arr)
	new_arr := []int{}

	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, heap.ExtractFrom(0).Value)
	}

	return new_arr
}
