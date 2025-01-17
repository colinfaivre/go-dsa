# 1.3 `M` 3Sum

[leetcode](https://leetcode.com/problems/3sum/) |
[youtube](https://www.youtube.com/watch?v=jzZsG8n2R9A.io%2F)

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that:
i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0
Notice that the solution set must not contain duplicate triplets.

Example 1:
> - Input: nums = [-1,0,1,2,-1,-4]
> - Output: [[-1,-1,2],[-1,0,1]]
> - Explanation:
> nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
> nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
> nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
> The distinct triplets are [-1,0,1] and [-1,-1,2].
> Notice that the order of the output and the order of the triplets does not matter.

Example 2:
> - Input: nums = [0,1,1]
> - Output: []
> - Explanation: The only possible triplet does not sum up to 0.

Example 3:
> - Input: nums = [0,0,0]
> - Output: [[0,0,0]]
> - Explanation: The only possible triplet sums up to 0.

Constraints:
> - 3 <= nums.length <= 3000
> - -10^5 <= nums[i] <= 10^5

<details>
	<summary><b>O(n^3) solution - brute force</b></summary>

- loop in nums
  - loop in nums
    - loop in nums
     - add the three numbers to res if they add up to zero
</details>

<details>
	<summary><b>O(n^2) solution - 2 pointers</b></summary>

- sort nums
- loop in nums
  - if the value is the same as in the previous iteration skip this one
  - init left and right pointers
  - loop while left < right
    - init threesum
    - decrement right if threesum > 0
    - increment left if threesum < 0
    - add num triplet to res if threesum == 0
    - increment left
    - inrement left while value is the same as previous

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    res := make([][]int, 0)
    for i := 0; i < len(nums)-2; i++ { // Ensure at least 3 elements
        if i > 0 && nums[i] == nums[i-1] { continue } // Skip duplicates

        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum < 0 { l++ }
            if sum > 0 { r-- }
            if sum == 0 {
                res = append(res, []int{ nums[i], nums[l], nums[r] })
                for l < r && nums[l] == nums[l+1] { l++ } // Skip duplicates
                for l < r && nums[r] == nums[r-1] { r-- } // Skip duplicates
                l++
                r--
            }
        }
    }

    return res
}
```

```js
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    const res = []
    nums.sort((a, b) => a-b)

    for (let i = 0; i < nums.length-2; i++) {
        if (i > 0 && nums[i] === nums[i-1]) continue
        let lo = i+1, hi = nums.length-1
        while (lo < hi) {
            const sum = nums[i]+nums[lo]+nums[hi]
            if (sum > 0) hi--
            if (sum < 0) lo++
            if (sum === 0) {
                res.push([nums[i], nums[lo], nums[hi]])
                while (lo < hi && nums[lo] === nums[lo+1]) lo++
                while (lo < hi && nums[hi] === nums[hi-1]) hi--
                hi--;lo++
            }
        }
    }

    return res
}
```
</details>
