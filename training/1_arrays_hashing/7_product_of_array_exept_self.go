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
O(n) solution - prefix-suffix
- loop in nums filling res with prefix products
- loop desc in nums multiplying res items with suffix products
***/

func ProductExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    prefix, suffix := 1, 1
    for i := 0; i < len(nums); i++ {
        res[i] = prefix
        prefix *= nums[i]
    }
    for j := len(nums) - 1; j >= 0; j-- {
        res[j] *= suffix
        suffix *= nums[j]
    }
    return res
}
