# 2.4 `M` Container With Most Water

[leetcode](https://leetcode.com/problems/container-with-most-water/)
[youtube](https://www.youtube.com/watch?v=UuiTKBwPgAo)

You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
> - Input: height = [1,8,6,2,5,4,8,3,7]
> - Output: 49
> - Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
> In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
> - Input: height = [1,1]
> - Output: 1

Constraints:
> - n == height.length
> - 2 <= n <= 10^5
> - 0 <= height[i] <= 10^4

<details>
	<summary><b>O(n^2) solution - brute force</b></summary>

- init res to 0
- loop in height with l
	- loop in height with r
		- set area to min of height at l and height at r times r minus l
		- set res to max of area and res 
</details>

<details>
	<summary><b>O(n) solution - 2 pointers</b></summary>

- init res to 0
- init l and r pointers on both sides of height
- loop in height while l < r
	- set area to min of height at l and height at r times r minus l
	- set res to max of area and res
	- if height at l less than or equals height at r
		- increment l
	- else decrement r

```go
func MaxAreaOn(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		res = max(area, res)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
```

```js
/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let lo = 0, hi = height.length - 1, res = 0
    while (lo < hi) {
        const area = Math.min(height[lo], height[hi])*(hi-lo)
        res = Math.max(res, area)
        if (height[lo] <= height[hi]) lo++
        else hi--
    }
    return res
};
```

</details>
