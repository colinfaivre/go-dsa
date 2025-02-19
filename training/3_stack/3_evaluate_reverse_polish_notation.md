# 3.3 Evaluate Reverse Polish Notation `M`

[leetcode](https://leetcode.com/problems/evaluate-reverse-polish-notation/)
[youtube](https://www.youtube.com/watch?v=iu0082c4HDE)

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
Evaluate the expression. Return an integer that represents the value of the expression.
Note that:
The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:
> - Input: tokens = ["2","1","+","3","*"]
> - Output: 9
> - Explanation: ((2 + 1) * 3) = 9

Example 2:
> - Input: tokens = ["4","13","5","/","+"]
> - Output: 6
> - Explanation: (4 + (13 / 5)) = 6

Example 3:
> - Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
> - Output: 22
> - Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
> 	- = ((10 * (6 / (12 * -11))) + 17) + 5
> 	- = ((10 * (6 / -132)) + 17) + 5
> 	- = ((10 * 0) + 17) + 5
> 	- = (0 + 17) + 5
> 	- = 17 + 5
> 	- = 22

Constraints:
> - 1 <= tokens.length <= 10^4
> - tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

<details>
	<summary><b>O(n) solution:</b></summary>

- init a stack
- loop in tokens
	- if the current value is a number push it into the stack
	- else pop two values from the stack and push back the result of the operation
- return last value from stack

```go
type NumStack struct {
	items []int
}

func (s *NumStack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *NumStack) Pop() int {
	removed := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return removed
}

func (s *NumStack) Peek() int {
	return s.items[0]
}

func EvalRPN(tokens []string) int {
	stack := NumStack{items: []int{}}

	for _, v := range tokens {
		if v == "+" {
			stack.Push(stack.Pop() + stack.Pop())
		} else if v == "*" {
			stack.Push(stack.Pop() * stack.Pop())
		} else if v == "/" {
			a, b := stack.Pop(), stack.Pop()
			stack.Push(b / a)
		} else if v == "-" {
			a, b := stack.Pop(), stack.Pop()
			stack.Push(b - a)
		} else {
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		}
	}

	return stack.Peek()
}
```
</details>
