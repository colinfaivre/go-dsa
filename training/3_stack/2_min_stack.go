package training

// 3.2 `M` Min Stack

/*** @LEETCODE https://leetcode.com/problems/min-stack/
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
Implement the MinStack class:
MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.
---
Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
Output
[null,null,null,null,-3,null,0,-2]
Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
---
Constraints:
-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104^ calls will be made to push, pop, top, and getMin.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=qkLl7nAwDPo
O(1) solution - two stacks
a stack for items and another for minimums
update the two stacks when pushing and popping
***/

type MinStack struct {
	items    []int
	minimums []int
}

func Constructor() MinStack {
	return MinStack{items: []int{}, minimums: []int{}}
}

func (s *MinStack) Push(val int) {
	s.items = append(s.items, val)
	minValue := val
	if len(s.minimums) != 0 {
		minValue = s.minimums[len(s.minimums)-1]
	}
	s.minimums = append(s.minimums, min(minValue, val))
}

func (s *MinStack) Pop() {
	s.minimums = s.minimums[:len(s.minimums)-1]
	s.items = s.items[:len(s.items)-1]
}

func (s *MinStack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *MinStack) GetMin() int {
	return s.minimums[len(s.minimums)-1]
}
