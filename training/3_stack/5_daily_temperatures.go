package training

// 3.5 `M` Daily Temperatures

/*** @LEETCODE https://leetcode.com/problems/daily-temperatures/
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
---
Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
---
Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
---
Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]
---
Constraints:
1 <= temperatures.length <= 10^5
30 <= temperatures[i] <= 100
***/

/*** @SOLUTION https://www.youtube.com/watch?v=cTBiBSnjO3c
O(n^2) solution - brute force:
- init res array of integers
- loop in temperatures with i
	- loop in temperatures with j
		- if temp at j > temp at i
			- append j-i to res
			- go to outer loop
	- append 0 to res
- return res

O(n) solution - stack:
- init res array of integers filled with zeros
- init stack of integer couples
- loop in temperatures (i, t)
	- while stack is not empty and t > top temp from stack
		- pop from the stack
		- set res at popped temp index to i - popped temp index
	- push the couple (t, i) on top of the stack
***/

func DailyTemperatures(temperatures []int) []int {
	res := []int{}
	for i := 0; i < len(temperatures); i++ {
		res = append(res, 0)
	}
	stack := [][2]int{}

	for i, t := range temperatures {
		for len(stack) != 0 && t > stack[len(stack)-1][0] {
			stackInd := stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			res[stackInd] = i - stackInd
		}
		stack = append(stack, [2]int{t, i})
	}

	return res
}
