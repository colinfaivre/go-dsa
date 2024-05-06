package ds

type Stack struct {
	items []int
}

func (stack *Stack) Push(item int) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() int {
	last_idx := len(stack.items) - 1
	to_remove := stack.items[last_idx]
	stack.items = stack.items[:last_idx]

	return to_remove
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}
