# 10.2 Combination Sum `M`

[leetcode](https://leetcode.com/problems/combination-sum/) |
[youtube](https://www.youtube.com/watch?v=GBKI9VSKdGg)

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order. The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.
The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:
> - Input: candidates = [2,3,6,7], target = 7
> - Output: [[2,2,3],[7]]

Explanation:
> - 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
> - 7 is a candidate, and 7 = 7.
> These are the only two combinations.

Example 2:
> - Input: candidates = [2,3,5], target = 8
> - Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
> - Input: candidates = [2], target = 1
> - Output: []

Constraints:
> - 1 <= candidates.length <= 30
> - 2 <= candidates[i] <= 40
> - All elements of candidates are distinct.
> - 1 <= target <= 40

<details>
  <summary><b>solution</b></summary>

```go
func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    var temp []int

    var backtrack func(start, remaining int)
    backtrack = func(start, remaining int) {
        if remaining == 0 {
            // If the remaining target is 0, we found a valid combination
            combination := append([]int{}, temp...)
            res = append(res, combination)
            return
        }
        if remaining < 0 {
            // If remaining is negative, no valid combination
            return
        }

        for i := start; i < len(candidates); i++ {
            // Add the current candidate to the temporary combination
            temp = append(temp, candidates[i])
            // Recursively try with the same candidate (unlimited usage allowed)
            backtrack(i, remaining-candidates[i])
            // Backtrack by removing the last added element
            temp = temp[:len(temp)-1]
        }
    }

    backtrack(0, target)
    return res
}
```
</details>
