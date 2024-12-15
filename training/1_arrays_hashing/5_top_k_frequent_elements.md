# 1.5 Top K Frequent Elements `M`

[LeetCode](https://leetcode.com/problems/top-k-frequent-elements/) [Youtube](https://www.youtube.com/watch?v=YPTqKIgVk-k)

Given an integer array nums and an integer k,
return the k most frequent elements.
You may return the answer in any order.

> **Example 1:**
> - Input: nums = [1,1,1,2,2,3], k = 2
> - Output: [1,2]
>
> **Example 2:**
> - Input: nums = [1], k = 1
> - Output: [1]
>
> **Constraints:**
> - 1 <= nums.length <= 10^5
> - -10^4 <= nums[i] <= 10^4
> - k is in the range [1, the number of unique elements in the array].
> - It is guaranteed that the answer is unique.

### O(nlogn) solution - hashmap/sort
- loop in nums inserting each num into a freqMap - O(n)
- sort the map items in descending order by freqs - O(nlogn)
- pop the first k biggest freq nums - O(k)

### O(klogn) solution - hashmap/maxheap
- loop in nums inserting each num into a freqMap - O(n)
- insert all freqMap items into a maxheap - O(n)
- pop the first k biggest frequency nums from the maxheap - O(klogn)

### O(n) solution - hashmap/bucketsort
- init freqMap w/ key as num and val as freq
- init freqArr as a nums lengthed array of int arrays
- init res as an int array
- loop in nums filling freqMap - O(n)
- loop in freqMap appending each nums in a subarray at freq index in a freqArr - O(n)
- loop in freqArr in descending order inserting num into res - O(n)
   - return res when its length reaches k

```go
// O(n) solution - hashmap/bucketsort
func TopKFrequent(nums []int, k int) []int {
    freqMap := map[int] int {}
    freqArr := [200_001][] int {}
    res := []int {}

    for _, val := range nums {
        freqMap[val]++
    }
    for key, val := range freqMap {
        freqArr[val] = append(freqArr[val], key)
    }
    for i := len(freqArr) - 1; i >= 0; i-- {
        for _, val := range freqArr[i] {
            res = append(res, val)
            if len(res) == k {
                return res
            }
        }
    }

    return res
}
```
