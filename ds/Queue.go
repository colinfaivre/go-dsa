package ds

type Queue struct {
	items []int
}

func (queue *Queue) Enqueue(item int) {
	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() int {
	to_remove := queue.items[0]
	queue.items = queue.items[1:]

	return to_remove
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}
