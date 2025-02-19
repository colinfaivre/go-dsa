# 4.1 Binary Search `E`

[leetcode](https://leetcode.com/problems/binary-search/) |
[youtube](https://www.youtube.com/watch?v=s4DPM8ct1pI)

Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
> - Input: nums = [-1,0,3,5,9,12], target = 9
> - Output: 4
> - Explanation: 9 exists in nums and its index is 4

Example 2:
> - Input: nums = [-1,0,3,5,9,12], target = 2
> - Output: -1
> - Explanation: 2 does not exist in nums so return -1

Constraints:
> - 1 <= nums.length <= 10^4
> - -10^4 < nums[i], target < 10^4
> - All the integers in nums are unique.
> - nums is sorted in ascending order.

<details>
	<summary><b>O(logn) solution: lo|hi pointers</b></summary>

- init lo and hi pointers
- loop in nums while lo is less or equal to hi
	- set mid as lo + (hi - lo) / 2
	- if nums at mid is target return mid
	- else if nums at mid is less than target
		- set lo to mid+1
	- else set hi to mid - 1
- return -1

```go
func Search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
```

```js
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let lo = 0, hi = nums.length-1
    while (lo <= hi) {
        const mid = lo + Math.floor((hi-lo)/2)
        if (nums[mid] > target) hi = mid-1
        if (nums[mid] < target) lo = mid+1
        if (nums[mid] === target) return mid
    }
    return -1
};
```
</details>
