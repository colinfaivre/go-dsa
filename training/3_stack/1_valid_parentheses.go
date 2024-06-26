package training

// 3.1 `E` Valid Parentheses

/*** @LEETCODE https://leetcode.com/problems/valid-parentheses/
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
---
Example 1:
Input: s = "()"
Output: true
---
Example 2:
Input: s = "()[]{}"
Output: true
---
Example 3:
Input: s = "(]"
Output: false
---
Constraints:
1 <= s.length <= 10^4
s consists of parentheses only '()[]{}'.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=WTzjTskDFMg
O(n) solution
- init charMap with opening parentheses as keys and closing as values
- init stack
- loop in s
  - if stack is empty or its last parenthesis does not match the current iteration char
    - push the current char on the top of the stack
  - else pop the last item from the stack
- return true if the stack is empty
***/

type Stack struct {
	items []rune
}

func (s *Stack) Peek() rune {
	res := s.items[len(s.items)-1]

	return res
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	removedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return removedItem
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func IsValid(s string) bool {
	charMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := Stack{}

	for _, v := range s {
		if stack.IsEmpty() || charMap[stack.Peek()] != v {
			stack.Push(v)
		} else {
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}
