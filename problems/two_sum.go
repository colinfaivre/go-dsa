package problems

// Given an unsorted array arr and a target t determine whether or not there exists x and y in arr
// such that x + y = t
// -----------
// Brute force solution: O(n2) exhaustive search
// Better solution: sort first, then binary search "t - x"
// Best solution: use a hashtable

func TwoSum(arr []int, t int) bool {
	sum_map := map[int]int{}

	for _, v := range arr {
		sum_map[v] = v
	}

	for _, v := range arr {
		y := t - v
		_, ok := sum_map[y]
		if ok {
			return true
		} else {
			sum_map[y] = y
		}
	}

	return false
}

// Returns the right number of distinct 'x+y' in the inclusive interval [-10_000, 10_000]
func GetSolution(arr []int) int {
	counter := 0

	for k := -10_000; k <= 10_000; k++ {
		if TwoSum(arr, k) {
			counter++
		}
	}

	return counter
}
