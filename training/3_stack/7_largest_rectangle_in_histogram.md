# 3.7 Largest Rectangle In Histogram `H`

[leetcode](https://leetcode.com/problems/largest-rectangle-in-histogram/) |
[youtube](https://www.youtube.com/watch?v=zx5Sw9130L0&source_ve_path=OTY3MTQ)

Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
return the area of the largest rectangle in the histogram.

Example 1:
> - Input: heights = [2,1,5,6,2,3]
> - Output: 10
> - Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

Example 2:
> - Input: heights = [2,4]
> - Output: 4

Constraints:
> - 1 <= heights.length <= 10^5
> - 0 <= heights[i] <= 10^4

<details>
  <summary><b>solution</b></summary>

```go
func largestRectangleArea(heights []int) int {
    stack := []int{}
    maxArea := 0
    heights = append(heights, 0)

    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            maxArea = max(maxArea, h*width)
        }

        stack = append(stack, i)
    }

    return maxArea
}
```
</details>
