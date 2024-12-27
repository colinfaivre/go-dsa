# 10.3. Permutations `M`

[leetcode](https://leetcode.com/problems/permutations/description/) |
[youtube](https://www.youtube.com/watch?v=FZe0UqISmUw)

Given an array nums of distinct integers, return all the possible
permutations. You can return the answer in any order.

Example 1:
> - Input: nums = [1,2,3]
> - Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
> - Input: nums = [0,1]
> - Output: [[0,1],[1,0]]

Example 3:
> - Input: nums = [1]
> - Output: [[1]]

Constraints:
> - 1 <= nums.length <= 6
> - -10 <= nums[i] <= 10
> - All the integers of nums are unique.

<details>
  <summary><b>solution</b></summary>

```go
func permute(nums []int) [][]int {
    var result [][]int

    var backtrack func(path []int, used []bool)
    backtrack = func(path []int, used []bool) {
        if len(path) == len(nums) {
            // Add a copy of the current path to the result
            temp := append([]int{}, path...)
            result = append(result, temp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            // Choose the current number
            used[i] = true
            backtrack(append(path, nums[i]), used)
            // Undo the choice (backtrack)
            used[i] = false
        }
    }

    backtrack([]int{}, make([]bool, len(nums)))
    return result
}
```
</details>
