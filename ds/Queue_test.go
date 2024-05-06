package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	queue.Dequeue()

	// fmt.Print("queue ", queue)
	if queue.items[0] != 3 {
		t.Errorf("queue first item expected %v but received %v", 3, queue.items[0])
	}
}
