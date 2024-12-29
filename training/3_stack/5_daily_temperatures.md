# 3.5 Daily Temperatures `M`

[leetcode](https://leetcode.com/problems/daily-temperatures/)
[youtube](https://www.youtube.com/watch?v=cTBiBSnjO3c)

Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
> - Input: temperatures = [73,74,75,71,69,72,76,73]
> - Output: [1,1,4,2,1,1,0,0]

Example 2:
> - Input: temperatures = [30,40,50,60]
> - Output: [1,1,1,0]

Example 3:
> - Input: temperatures = [30,60,90]
> - Output: [1,1,0]

Constraints:
> - 1 <= temperatures.length <= 10^5
> - 30 <= temperatures[i] <= 100

<details>
	<summary><b>O(n^2) solution - brute force</b></summary>

- init res array of integers
- loop in temperatures with i
	- loop in temperatures with j
		- if temp at j > temp at i
			- append j-i to res
			- go to outer loop
	- append 0 to res
- return res
</details>

<details>
	<summary><b>O(n) solution - stack</b></summary>

- init res array of integers filled with zeros
- init stack of indexes
- loop in temperatures (i, t)
	- while stack is not empty and t > top temp index from stack
		- pop from the stack
		- set res at popped temp index to i - popped temp index
	- push i on top of the stack
 - return res

```go
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    res := make([]int, n)
    stack := []int{}

    for i, t := range temperatures {
        for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
            topIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[topIndex] = i - topIndex
        }
        stack = append(stack, i)
    }

    return res
}
```
</details>
