package training

// 8.1 `E` Kth Largest Element In a Stream

/*** @LEETCODE https://leetcode.com/problems/kth-largest-element-in-a-stream/
Design a class to find the kth largest element in a stream.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Implement KthLargest class:
- KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
- int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.
---
Example 1:
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]
Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
---
Constraints:
1 <= k <= 10^4
0 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
-10^4 <= val <= 10^4
At most 10^4 calls will be made to add.
It is guaranteed that there will be at least k elements in the array when you search for the kth element.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=hOjcdrqMoQ8
***/

import "container/heap"

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { 
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	minHeap MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		minHeap: MinHeap{},
		k:       k,
	}
	heap.Init(&kl.minHeap)
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.minHeap, val)
	if kl.minHeap.Len() > kl.k {
		heap.Pop(&kl.minHeap) // Remove the smallest element to maintain size k
	}
	return kl.minHeap[0] // Return the kth largest element
}
