package training

// 9.1 `M` Number of Islands

/*** @LEETCODE https://leetcode.com/problems/number-of-islands/
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
---
Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
---
Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
---
Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=pV2kpPD66nE
***/

func dfs(grid [][]byte, r, c int) {
    rows := len(grid)
    cols := len(grid[0])
    if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
        return // Base case: If out of bounds or on water ('0'), return
    }

    grid[r][c] = '0'  // Mark the cell as visited by setting it to '0'

    dfs(grid, r+1, c) // Down
    dfs(grid, r-1, c) // Up
    dfs(grid, r, c+1) // Right
    dfs(grid, r, c-1) // Left
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])
    count := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                count++
                dfs(grid, r, c) // Start DFS to mark the entire island
            }
        }
    }

    return count
}
