# 2.5 Trapping Rain Water `M`

[leetcode](https://leetcode.com/problems/trapping-rain-water/) |
[youtube](https://www.youtube.com/watch?v=ZI2z5pq0TqA)

Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.

Example 1:
> - Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
> - Output: 6
> - Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
> - Input: height = [4,2,0,3,2,5]
> - Output: 9

Constraints:
> - n == height.length
> - 1 <= n <= 2 * 10^4
> - 0 <= height[i] <= 10^5

<details>
  <summary><b>O(n) solution - 2 pointers</b></summary>

```go
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}
```
</details>
