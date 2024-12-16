# 1.9 Longest Consecutive Sequence `M`

[leetcode](https://leetcode.com/problems/longest-consecutive-sequence/)
[youtube](https://www.youtube.com/watch?v=P6RZZMu_maU)

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

> **Example 1:**
> - Input: nums = [100,4,200,1,3,2]
> - Output: 4
> - Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
>
> **Example 2:**
> - Input: nums = [0,3,7,2,5,8,4,6,0,1]
> - Output: 9
>
> **Constraints:**
> - 0 <= nums.length <= 10^5
> - -10^9 <= nums[i] <= 10^9

### TC:O(n)/SC:O(n) solution - hashSet
- initialize longest to 0
- loop in nums to create a set containing each num
- loop in set with value
  - check if n-1 is not in set // start of sequence
  - initialize length to 0
  - loop in set while value+length exists incrementing length
  - set longest to max between length and itself

```go
func longestConsecutive(nums []int) int {
    numSet := map[int]bool{}
    longest := 0

    for _, v := range nums {
        numSet[v] = true
    }
    for v := range numSet {
        if !numSet[v-1] {
            length := 0
            for numSet[v+length] {
                length++
            }
            if length > longest {
                longest = length
            }
        }
    }

    return longest
}
```