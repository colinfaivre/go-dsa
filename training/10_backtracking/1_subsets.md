# 10.1 Subsets `M`

[leetcode](https://leetcode.com/problems/subsets/description/) |
[youtube](https://www.youtube.com/watch?v=REOH22Xwdkk)

Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
> - Input: nums = [1,2,3]
> - Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
> - Input: nums = [0]
> - Output: [[],[0]]

Constraints:
> - 1 <= nums.length <= 10
> - -10 <= nums[i] <= 10
> - All the numbers of nums are unique.

<details>
  <summary><b>O(n2^n) solution - backtracking</b></summary>

```go
func subsets(nums []int) [][]int {
    var result [][]int
    var subset []int

    var backtrack func(int)
    backtrack = func(start int) {
        // Append a copy of the current subset to the result
        temp := make([]int, len(subset))
        copy(temp, subset)
        result = append(result, temp)

        // Generate all subsets starting from the current position
        for i := start; i < len(nums); i++ {
            subset = append(subset, nums[i]) // Include nums[i] in the current subset
            backtrack(i + 1) // Recur with the next element
            subset = subset[:len(subset)-1] // Backtrack by removing nums[i]
        }
    }

    backtrack(0) // Start backtracking from index 0
    return result
}
```
</details>
