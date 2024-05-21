package problems

import (
	"github.com/colinfaivre/go-dsa/algorithms"
	"github.com/colinfaivre/go-dsa/datastructures"
)

// @LEETCODE https://leetcode.com/problems/find-median-from-data-stream/description/
// @MEDIUM https://yuminlee2.medium.com/leetcode-295-find-median-from-data-stream-9d3b4ff5270f

// Median:
// Middle value separating the greater and lesser halves of a data set (sorted)
// ex: 1, 2, 2, *3*, 4, 7, 9

// Naive solution O(n * nlogn) solution <=> O(n2)
func NaiveMedianMaintenance(number_list []int) []int {
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

// https://kb.novaordis.com/index.php/The_Median_Maintenance_Problem
func HeapMedianMaintenance(number_list []int) []int {
	median_list := []int{}
	h_low := datastructures.NewHeap(false)
	h_high := datastructures.NewHeap(true)

	for _, v := range number_list {
		if h_low.Size() == 0 || v <= h_low.Peek().Value {
			h_low.Insert(datastructures.Item{Value: v})
		} else {
			h_high.Insert(datastructures.Item{Value: v})
		}
		rebalance(h_low, h_high)
		median_list = append(median_list, h_low.Peek().Value)
	}

	return median_list
}

func rebalance(h_low, h_high *datastructures.Heap) {
	diff := h_low.Size() - h_high.Size()

	if diff == 2 {
		h_high.Insert(h_low.ExtractFrom(0))
	} else if diff == -1 {
		h_low.Insert(h_high.ExtractFrom(0))
	}
}

func GetResult(median_list []int) int {
	total := 0

	for _, v := range median_list {
		total += v
	}
	return total % 10_000
}
