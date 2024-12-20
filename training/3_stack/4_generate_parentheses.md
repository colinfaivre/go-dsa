# 3.4 Generate Parentheses `M`

[leetcode](https://leetcode.com/problems/generate-parentheses/)
[youtube](https://www.youtube.com/watch?v=s9fokUqJ76A)

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
> - Input: n = 3
> - Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
> - Input: n = 1
> - Output: ["()"]

Constraints:
> - 1 <= n <= 8

<details>
	<summary><b>backtrack solution</b></summary>

```go
func GenerateParenthesis(n int) []string {
	// only add open parentheses if open < n
	// only add closing parentheses if closed < open
	// valid IIF if open == closed == n

	stack := []string{}
	res := []string{}

	backtrack(0, 0, &stack, &res, n)
	return res
}

func backtrack(openN, closedN int, stack *[]string, res *[]string, n int) {
	if openN == closedN && openN == n {
		*res = append(*res, strings.Join(*stack, ""))
		return
	}

	if openN < n {
		*stack = append(*stack, "(")
		backtrack(openN+1, closedN, stack, res, n)
		*stack = (*stack)[:len(*stack)-1]
	}

	if closedN < openN {
		*stack = append(*stack, ")")
		backtrack(openN, closedN+1, stack, res, n)
		*stack = (*stack)[:len(*stack)-1]
	}
}
```
</details>
