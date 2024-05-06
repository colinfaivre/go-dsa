package ds

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	stack.Pop()
	stack.Pop()

	// fmt.Print("stack ", stack)
	if stack.items[0] != 3 {
		t.Errorf("stack first item expected %v but received %v", 3, stack.items[0])
	}
}
