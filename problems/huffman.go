// @BLOG https://reintech.io/blog/comprehensive-guide-huffman-coding-algorithm-go
// @GOSTD https://pkg.go.dev/container/heap

package problems

import "container/heap"

type HuffmanNode struct {
	charIdx int
	freq    int
	left    *HuffmanNode
	right   *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*HuffmanNode)) }

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func HuffmanCoding(charFreqs []int) *HuffmanNode {
	freqMap := make(map[int]int)

	for idx, number := range charFreqs {
		freqMap[idx] = number
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for idx, freq := range freqMap {
		heap.Push(pq, &HuffmanNode{charIdx: idx, freq: freq})
	}

	for pq.Len() > 1 {
		first := heap.Pop(pq).(*HuffmanNode)
		second := heap.Pop(pq).(*HuffmanNode)

		parent := &HuffmanNode{
			freq:  first.freq + second.freq,
			left:  first,
			right: second,
		}

		heap.Push(pq, parent)
	}

	return heap.Pop(pq).(*HuffmanNode)
}
