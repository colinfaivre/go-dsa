package problems

// @LEETCODE https://leetcode.com/problems/two-sum/description/
// Given an unsorted array arr and a target t determine whether or not there exists x and y in arr
// such that x + y = t
// -----------
// Brute force solution: O(n2) exhaustive search
// Better solution: sort first, then binary search "t - x"
// Best solution: use a hashtable

func TwoSum(arr []int, t int) bool {
	sum_map := map[int]bool{}
	for _, v := range arr {
		sum_map[v] = true
	}

	for _, x := range arr {
		y := t - x
		if sum_map[y] && x != y {
			return true
		}
	}

	return false
}

// @LEETCODE https://leetcode.com/problems/count-the-number-of-fair-pairs/
// Returns the right number of distinct 'x+y' in the inclusive interval [-10_000, 10_000]
// Very slow, try using sliding window technique instead of map
func GetSolution(arr []int) int {
	counter := 0

	for k := -10000; k <= 10000; k++ {
		if TwoSum(arr, k) {
			counter++
		}
	}

	return counter
}
