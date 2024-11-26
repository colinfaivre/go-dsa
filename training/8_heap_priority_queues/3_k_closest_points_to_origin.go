package training

// 8.3 `M` K Closest Points to Origin

/*** @LEETCODE https://leetcode.com/problems/k-closest-points-to-origin/
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k,
return the k closest points to the origin (0, 0).
The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
---
Example 1:
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
---
Example 2:
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.
---
Constraints:
1 <= k <= points.length <= 10^4
-10^4 <= xi, yi <= 10^4
***/

/*** @SOLUTION https://www.youtube.com/watch?v=rI2EBUEMfTk
***/

import (
	"container/heap"
)

type Point struct {
	coords   []int
	distance int
}

// Creates a new Point with calculated distance from origin
func New(coords []int) Point {
	xSquared := coords[0] * coords[0]
	ySquared := coords[1] * coords[1]

	return Point{
		coords:   coords,
		distance: xSquared + ySquared,
	}
}

// MinHeap is a custom heap type for Points
type MinHeap []Point
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	// Create a MinHeap and initialize it
	pointMinHeap := &MinHeap{}
	heap.Init(pointMinHeap)

	// Populate the heap
	for _, val := range points {
		heap.Push(pointMinHeap, New(val))
	}

	// Extract k closest points
	res := [][]int{}
	for i := 0; i < k; i++ {
		minPoint := heap.Pop(pointMinHeap).(Point).coords
		res = append(res, minPoint)
	}

	return res
}

