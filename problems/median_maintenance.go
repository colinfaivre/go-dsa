package problems

import "github.com/colinfaivre/go-dsa/algorithms"

// Median:
// Middle value separating the greater and lesser halves of a data set (sorted)
// ex: 1, 2, 2, *3*, 4, 7, 9

// Naive solution O(n * nlogn) solution <=> O(n2)
// @TODO could be done in O(logn) using MaxHeap and MinHeap
func MedianMaintenance(number_list []int) []int {
	median_list := []int{}

	for i := 0; i < len(number_list); i++ {
		curr_slice := number_list[0 : i+1]
		curr_slice = algorithms.QuickSort(curr_slice)

		if i%2 == 0 {
			median_list = append(median_list, curr_slice[(i+1)/2])

		} else {
			median_list = append(median_list, curr_slice[i/2])
		}
	}

	return median_list
}

func GetResult(median_list []int) int {
	total := 0

	for _, v := range median_list {
		total += v
	}
	return total % 10_000
}
