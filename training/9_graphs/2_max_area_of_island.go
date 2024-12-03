package training

// 9.2 `M` Max Area of Island

/*** @LEETCODE https://leetcode.com/problems/max-area-of-island/
You are given an m x n binary matrix grid.
An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
You may assume all four edges of the grid are surrounded by water.
The area of an island is the number of cells with a value 1 in the island.
Return the maximum area of an island in grid. If there is no island, return 0.
---
Example 1:
Input: grid = [
	[0,0,1,0,0,0,0,1,0,0,0,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,1,1,0,1,0,0,0,0,0,0,0,0],
	[0,1,0,0,1,1,0,0,1,0,1,0,0],
	[0,1,0,0,1,1,0,0,1,1,1,0,0],
	[0,0,0,0,0,0,0,0,0,0,1,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,0,0,0,0,0,0,1,1,0,0,0,0]
]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
---
Example 2:
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
---
Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=iJGr1OtmH0c
***/

func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    if rows == 0 || cols == 0 { return 0 }

    var maxArea int
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                maxArea = max(maxArea, getMaxArea(grid, r, c))
            }
        }
    }

    return maxArea
}

func getMaxArea(grid [][]int, r, c int) int {
    rows, cols := len(grid), len(grid[0])
    if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 { return 0 }

    grid[r][c] = 0
    maxDown := getMaxArea(grid, r+1, c)
    maxUp := getMaxArea(grid, r-1, c)
    maxLeft := getMaxArea(grid, r, c-1)
    maxRight := getMaxArea(grid, r, c+1)

    return 1 + maxDown + maxUp + maxLeft + maxRight
}

