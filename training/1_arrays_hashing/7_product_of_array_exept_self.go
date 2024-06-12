package training

// 1.7 `M` Product of Array Except Self

/*** @LEETCODE https://leetcode.com/problems/product-of-array-except-self/
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.
---
Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
---
Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
---
Constraints:
2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=bNvIQI2wAjk
TC:O(n)/SC:O(n) solution - prefix-suffix
- loop in nums computing arr of prefixProduct
- loop desc in nums computing arr of suffixProduct
- loop in nums getting the product of the previous prefixProduct and the next suffixProduct
TC:O(1)/SC:O(n) solution - prefix-suffix
- loop in nums filling res with prefix products
- loop desc in nums multiplying res items with suffix products
***/

// TC:O(n)/SC:O(n) solution - prefix-suffix
func ProductExceptSelfSCOn(nums []int) []int {
	prefixProducts := make([]int, len(nums))
	suffixProducts := make([]int, len(nums))
	res := make([]int, len(nums))

	prefixProducts[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixProducts[i] = prefixProducts[i-1] * nums[i]
	}

	suffixProducts[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffixProducts[i] = suffixProducts[i+1] * nums[i]
	}

	res[0] = suffixProducts[1]
	res[len(nums)-1] = prefixProducts[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = prefixProducts[i-1] * suffixProducts[i+1]
	}

	return res
}

// TC:O(1)/SC:O(n) solution - prefix-suffix
func ProductExceptSelfSCO1(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
