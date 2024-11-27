package training

// 8.4 `M` Kth Largest Element In An Array

/*** @LEETCODE https://leetcode.com/problems/kth-largest-element-in-an-array/
Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Can you solve it without sorting?
---
Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
---
Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
---
Constraints:
1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=XEmy13g1Qxc
***/

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
    numsMaxHeap := &MaxHeap{}
    heap.Init(numsMaxHeap)
    var res int
    for _, val := range nums {
        heap.Push(numsMaxHeap, val)
    }
    for i := 1; i <= k; i++ {
        res = heap.Pop(numsMaxHeap).(int)
    }
    return res
}
