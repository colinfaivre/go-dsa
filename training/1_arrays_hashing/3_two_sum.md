## 1.3 `E` Two Sum

[leetcode](https://leetcode.com/problems/two-sum/)
[youtube](https://www.youtube.com/watch?v=KLlXCFG5TnA)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

> **Example 1:**
> - Input: nums = [2,7,11,15], target = 9
> - Output: [0,1]
> - Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
>
> **Example 2:**
> - Input: nums = [3,2,4], target = 6
> - Output: [1,2]
>
> **Example 3:**
> - Input: nums = [3,3], target = 6
> - Output: [0,1]
>
> **Constraints:**
> - 2 <= nums.length <= 10^4
> - -10^9 <= nums[i] <= 10^9
> - -10^9 <= target <= 10^9
> - Only one valid answer exists.

### O(n^2) solution - bruteforce
- loop in nums w/ i
  - loop in nums w/ j > i
    - if nums[i] + nums[j] == target return [i, j]

### O(n) solution - hashmap
- init hashmap
- loop in nums w/ i
  - if nums[i] key in hashmap return [i, hashmap[arr[i]]]
  - set hashmap[t-arr[i]] to i

```go
func TwoSum(nums []int, target int) []int {
    hashMap := map[int] int {}

    for i, v := range nums {
        if _, ok := hashMap[v]; ok {
            return []int {i, hashMap[v]}
        }
        hashMap[target - v] = i
    }

    return []int {0, 0} // default return
}
```