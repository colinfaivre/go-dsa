package sort

import "slices"

// O(nlogn)
// The quick sort is probably the most used sorting algorithm
// It performs better than other nlogn algorithms
// It uses divide and conquer approach
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left []int
	var right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return slices.Concat(QuickSort(left), []int{pivot}, QuickSort(right))
}
