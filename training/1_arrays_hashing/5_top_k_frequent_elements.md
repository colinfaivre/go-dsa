# 1.5 Top K Frequent Elements `M`

[leetCode](https://leetcode.com/problems/top-k-frequent-elements/) |
[youtube](https://www.youtube.com/watch?v=YPTqKIgVk-k)

Given an integer array nums and an integer k,
return the k most frequent elements.
You may return the answer in any order.

Example 1:
> - Input: nums = [1,1,1,2,2,3], k = 2
> - Output: [1,2]

Example 2:
> - Input: nums = [1], k = 1
> - Output: [1]

Constraints:
> - 1 <= nums.length <= 10^5
> - -10^4 <= nums[i] <= 10^4
> - k is in the range [1, the number of unique elements in the array].
> - It is guaranteed that the answer is unique.

<details>
    <summary><b>O(nlogn) solution - hashmap/sort</b></summary>

- loop in nums inserting each num into a freqMap - O(n)
- sort the map items in descending order by freqs - O(nlogn)
- pop the first k biggest freq nums - O(k)

```js
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function(nums, k) {
    const freqMap = new Map()
    for (const num of nums) freqMap.set(num, (freqMap.get(num) || 0) + 1)
    const sorted = [...freqMap.entries()].sort((a, b) => b[1] - a[1])

    return sorted.slice(0, k).map(entry => entry[0])
}
```
</details>

<details>
    <summary><b>O(klogn) solution - hashmap/maxheap</b></summary>

- loop in nums inserting each num into a freqMap - O(n)
- insert all freqMap items into a maxheap - O(n)
- pop the first k biggest frequency nums from the maxheap - O(klogn)

```go
type IntHeap []heapItem
type heapItem struct {
    num int
    freq int
}
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) { *h = append(*h, x.(heapItem)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int] int)
    for _, val := range nums {
        freqMap[val]++
    }

    h := &IntHeap{}
    heap.Init(h)
    for key, val := range freqMap {
        heap.Push(h, heapItem{num: key, freq: val})
    }

    var res []int
    for i := 1; i <= k; i++ {
        res = append(res, heap.Pop(h).(heapItem).num)
    }

    return res
}
```
</details>

<details>
    <summary><b>O(n) solution - hashmap/bucketsort</b></summary>


- init freqMap w/ key as num and val as freq
- init freqArr as a nums lengthed array of int arrays
- init res as an int array
- loop in nums filling freqMap - O(n)
- loop in freqMap appending each nums in a subarray at freq index in a freqArr - O(n)
- loop in freqArr in descending order inserting num into res - O(n)
   - return res when its length reaches k

```go
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
</details>
