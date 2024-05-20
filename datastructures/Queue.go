package datastructures

// WIKI https://en.wikipedia.org/wiki/Queue_(abstract_data_type)

type Queue struct {
	items LinkedList
}

// Returns a string representation of the queue
func (q *Queue) ToString() string {
	return q.items.ToString()
}

// O(1): Adds an item at the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items.AddLast(item)
}

// O(1): Removes and returns the item in front of the queue
func (q *Queue) Dequeue() int {
	return q.items.RemoveFirst()
}

// O(1): Returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.items.isEmpty()
}
