# 1.1 Contains Duplicate `E`

[leetcode](https://leetcode.com/problems/contains-duplicate/) |
[youtube](https://www.youtube.com/watch?v=3OamzN90kPg)

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
> - Input: nums = [1,2,3,1]
> - Output: true

Example 2:
> - Input: nums = [1,2,3,4]
> - Output: false

Example 3:
> - Input: nums = [1,1,1,3,3,4,3,2,4,2]
> - Output: true

Constraints:
> - 1 <= nums.length <= 10^5
> - -10^9 <= nums[i] <= 10^9

<details>
  <summary><b>O(n^2) solution - bruteforce</b></summary>
  
- loop in nums with i
  - loop in nums with j
    - return true if nums[i] == nums[j]
- return false

</details>

<details>
  <summary><b>O(nlogn) solution: sorting</b></summary>
  
  - sort nums
  - loop in nums
    - return true if consecutive items are equal
  - return false
</details>

<details>
  <summary><b>O(n) solution: hashSet</b></summary>
  
- init hashset
- loop in nums with i
  - return true if nums[i] is in hashset
  - set hashset[nums[i]] to true
- return false

```go
func ContainsDuplicate(nums []int) bool {
    numsSet := map[int] bool {}
    for i := range nums {
        if numsSet[nums[i]] {
            return true
        }
        numsSet[nums[i]] = true
    }
    return false
}
```

```js
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    const seen = new Set()
    for (const num of nums) {
        if (seen.has(num)) return true
        seen.add(num)
    }

    return false
}
```
</details>
