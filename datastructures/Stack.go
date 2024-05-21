// @WIKI https://en.wikipedia.org/wiki/Stack_(abstract_data_type)

package datastructures

type Stack struct {
	items LinkedList
}

// Returns a string representation of the stack
func (s *Stack) ToString() string {
	return s.items.ToString()
}

// O(1): Adds an item on the top of the stack
func (s *Stack) Push(item int) {
	s.items.AddFirst(item)
}

// O(1): Removes and returns the item on the top of the stack
func (s *Stack) Pop() int {
	return s.items.RemoveFirst()
}

// O(1): Returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.items.isEmpty()
}
