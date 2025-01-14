# 5.6 Sliding Window Maximum `H`

[leetcode](https://leetcode.com/problems/sliding-window-maximum/) |
[youtube](https://www.youtube.com/watch?v=DfljaUwZsOk)

You are given an array of integers nums, there is a sliding window of size k
which is moving from the very left of the array to the very right.
You can only see the k numbers in the window.
Each time the sliding window moves right by one position.
Return the max sliding window.

Example 1:
> - Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
> - Output: [3,3,5,5,6,7]
> - Explanation:
>
> | Window position | Max |
> | --------------- | --- |
> | [1  3  -1] -3  5  3  6  7 | 3 |
> | 1 [3  -1  -3] 5  3  6  7  | 3 |
> | 1  3 [-1  -3  5] 3  6  7  | 5 |
> | 1  3  -1 [-3  5  3] 6  7  | 5 |
> | 1  3  -1  -3 [5  3  6] 7  | 6 |
> | 1  3  -1  -3  5 [3  6  7] | 7 |

Example 2:
> - Input: nums = [1], k = 1
> - Output: [1]

Constraints:
> - 1 <= nums.length <= 10^5
> - -10^4 <= nums[i] <= 10^4
> - 1 <= k <= nums.length

<details>
 <summary><b>solution - sliding win / heap</b></summary>

```go
type Pair struct {
	val int
	idx int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		heap.Push(h, Pair{val: nums[i], idx: i})
		for (*h)[0].idx <= i-k {
			heap.Pop(h)
		}
		// Add the maximum for the window to the result
		if i >= k-1 {
			res = append(res, (*h)[0].val)
		}
	}

	return res
}
```
</details>
