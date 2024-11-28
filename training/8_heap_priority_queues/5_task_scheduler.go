package training

// 8.5 `M` Task Scheduler

/*** @LEETCODE https://leetcode.com/problems/task-scheduler/
You are given an array of CPU tasks, each represented by letters A to Z, and a cooling time, n.
Each cycle or interval allows the completion of one task.
Tasks can be completed in any order, but there's a constraint:
identical tasks must be separated by at least n intervals due to cooling time.
​Return the minimum number of intervals required to complete all tasks.
---
Example 1:
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
After completing task A, you must wait two cycles before doing A again.
The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle.
By the 4th cycle, you can do A again as 2 intervals have passed.
---
Example 2:
Input: tasks = ["A","C","A","B","D","B"], n = 1
Output: 6
Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
With a cooling interval of 1, you can repeat a task after just one other task.
---
Example 3:
Input: tasks = ["A","A","A", "B","B","B"], n = 3
Output: 10
Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
There are only two types of tasks, A and B, which need to be separated by 3 intervals.
This leads to idling twice between repetitions of these tasks.
---
Constraints:
1 <= tasks.length <= 10^4
tasks[i] is an uppercase English letter.
0 <= n <= 100
***/

/*** @SOLUTION https://www.youtube.com/watch?v=s8p8ukTyA2I
***/

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	// Step 1: Count task frequencies
	tasksFreqs := map[byte]int{}
	for _, val := range tasks {
		tasksFreqs[val]++
	}

	// Step 2: Initialize a max-heap for task frequencies
	maxFreqHeap := &MaxHeap{}
	heap.Init(maxFreqHeap)
	for _, freq := range tasksFreqs {
		heap.Push(maxFreqHeap, freq)
	}

	time := 0
	queue := []struct {
		freq      int
		readyTime int
	}{}

	// Step 3: Simulate the task execution
	for maxFreqHeap.Len() > 0 || len(queue) > 0 {
		time++

		// Execute the most frequent task if available
		if maxFreqHeap.Len() > 0 {
			freq := heap.Pop(maxFreqHeap).(int)
			freq--
			if freq > 0 {
				// Push the task into the cooldown queue
				queue = append(queue, struct {
					freq      int
					readyTime int
				}{freq: freq, readyTime: time + n})
			}
		}

		// Check if any task in the cooldown queue is ready
		if len(queue) > 0 && queue[0].readyTime == time {
			heap.Push(maxFreqHeap, queue[0].freq)
			queue = queue[1:] // Remove the task from the queue
		}
	}

	return time
}
