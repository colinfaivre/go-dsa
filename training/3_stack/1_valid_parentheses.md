# 3.1 Valid Parentheses `E`

[leetcode](https://leetcode.com/problems/valid-parentheses/)
[youtube](https://www.youtube.com/watch?v=WTzjTskDFMg)

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
> - Input: s = "()"
> - Output: true

Example 2:
> - Input: s = "()[]{}"
> - Output: true

Example 3:
> - Input: s = "(]"
> - Output: false

Constraints:
> - 1 <= s.length <= 10^4
> - s consists of parentheses only '()[]{}'.

<details>
	<summary><b>O(n) solution</b></summary>
	
- init charMap with opening parentheses as keys and closing as values
- init stack
- loop in s
  - if stack is empty or its last parenthesis does not match the current iteration char
    - push the current char on the top of the stack
  - else pop the last item from the stack
- return true if the stack is empty

```go
func IsValid(s string) bool {
	charMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	var stack []rune
	for _, v := range s {
		if len(stack) == 0 || charMap[stack[len(stack)-1]] != v {
			stack = append(stack, v) // Push to the stack
		} else {
			stack = stack[:len(stack)-1] // Pop from the stack
		}
	}

	return len(stack) == 0
}
```

```js
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const parenMap = {
        "(": ")",
        "[": "]",
        "{": "}",
    }
    stack = []

    for (const char of s) {
        if (parenMap[stack[stack.length-1]] != char || stack.length === 0) stack.push(char)
        else stack.pop()
    }

    return stack.length === 0
};
```
</details>
