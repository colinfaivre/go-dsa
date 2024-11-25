package training

// 8.2 `E` Last Stone Weight

/*** @LEETCODE https://leetcode.com/problems/last-stone-weight/
You are given an array of integers stones where stones[i] is the weight of the ith stone.
We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together.
Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.
Return the weight of the last remaining stone. If there are no stones left, return 0.
---
Example 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
---
Example 2:
Input: stones = [1]
Output: 1
---
Constraints:
1 <= stones.length <= 30
1 <= stones[i] <= 1000
***/

/*** @SOLUTION https://www.youtube.com/watch?v=B-QCq79-Vfw
***/

func lastStoneWeight(stones []int) int {
    stoneHeap := MaxHeap{}

    for _, val := range stones {
        stoneHeap.Insert(val)
    }

    for stoneHeap.Size() != 0 {
        y := stoneHeap.ExtractMax()
        if stoneHeap.Size() == 0 { return y }
        x := stoneHeap.ExtractMax()
        if x != y { stoneHeap.Insert(y - x) }
    }

    return 0
}

// MaxHeap represents a max-heap data structure
type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Size() int {
    return len(h.array)
}

// Insert adds a new element to the heap
func (h *MaxHeap) Insert(val int) {
	h.array = append(h.array, val)
	h.heapifyUp(len(h.array) - 1)
}

// ExtractMax removes and returns the largest element in the heap
func (h *MaxHeap) ExtractMax() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	// Get the max element (root)
	max := h.array[0]

	// Replace the root with the last element
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	// Heapify down to maintain heap property
	h.heapifyDown(0)

	return max
}

// heapifyUp restores the heap property by moving a node up
func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown restores the heap property by moving a node down
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	for left <= lastIndex {
		if right <= lastIndex && h.array[right] > h.array[left] {
			childToCompare = right
		} else {
			childToCompare = left
		}

		if h.array[index] > h.array[childToCompare] {
			return
		}

		h.swap(index, childToCompare)
		index = childToCompare
		left, right = leftChild(index), rightChild(index)
	}
}

// parent returns the index of the parent node
func parent(index int) int {
	return (index - 1) / 2
}

// leftChild returns the index of the left child node
func leftChild(index int) int {
	return 2*index + 1
}

// rightChild returns the index of the right child node
func rightChild(index int) int {
	return 2*index + 2
}

// swap exchanges two elements in the heap array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// PrintHeap prints the heap array (for debugging)
func (h *MaxHeap) PrintHeap() {
	fmt.Println(h.array)
}
