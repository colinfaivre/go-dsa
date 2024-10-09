package training

// 4.5 `M` Search In Rotated Sorted Array

/*** @LEETCODE https://leetcode.com/problems/search-in-rotated-sorted-array/
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the possible rotation and an integer target,
return the index of target if it is in nums, or -1 if it is not in nums.
You must write an algorithm with O(log n) runtime complexity.
---
Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
---
Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
---
Example 3:
Input: nums = [1], target = 0
Output: -1
---
Constraints:
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-10^4 <= target <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=U8XENwh8Oy8
O(n) solution:
- loop in nums w/ i and val:
  - return i when val equals target
- return - 1
O(logn) solution:
- init l and r pointers to 0 and nums length - 1
- loop while l is less or equal to r:
  - init m to (l+r)/2
  - return m if nums at m equals target
  - if nums at l is less or equal to nums at m --> left sorted portion
    - if target is greater than nums at m or target is less than nums at l
      - set l to m+1
    - else set r to m-1
  - else --> right sorted portion
    - if target is less than nums at m or target is greater than nums at r
      - set r to m-1
    - else set l to m+1
- return -1
***/

func Search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        m := (l+r)/2
        if nums[m] == target {
            return m
        }

        if nums[l] <= nums[m] { // left sorted portion
            if target > nums[m] || target < nums[l] {
                l = m+1
            } else {
                r = m-1
            }
        } else { // right sorted portion
            if target < nums[m] || target > nums[r] {
                r = m-1
            } else {
                l = m+1
            }
        }
    }
    return -1
}
