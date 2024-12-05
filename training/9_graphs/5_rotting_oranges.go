package training

// 9.5 `M` Rotting Oranges

/*** @LEETCODE https://leetcode.com/problems/rotting-oranges/
You are given an m x n grid where each cell can have one of three values:
- 0 representing an empty cell,
- 1 representing a fresh orange, or
- 2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.
---
Example 1:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
---
Example 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
---
Example 3:
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
---
Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] is 0, 1, or 2.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=y704fEOx0s0
***/

func orangesRotting(grid [][]int) int {
    var queue [][2]int
    var time, freshTotal int
    ROWS, COLS := len(grid), len(grid[0])

    // Count fresh oranges and initialize the queue with rotten oranges
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == 1 {
                freshTotal++
            } else if grid[r][c] == 2 {
                queue = append(queue, [2]int{r, c})
            }
        }
    }

    if freshTotal == 0 {
        return 0 // No fresh oranges to rot
    }

    var start int // Pointer to the front of the queue
    for start < len(queue) && freshTotal > 0 {
        size := len(queue)
        for i := start; i < size; i++ {
            r, c := queue[i][0], queue[i][1]
            dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
            for _, dir := range dirs {
                row, col := r+dir[0], c+dir[1]
                if row < 0 || row == ROWS || col < 0 || col == COLS || grid[row][col] != 1 {
                    continue
                }
                grid[row][col] = 2
                queue = append(queue, [2]int{row, col})
                freshTotal--
            }
        }
        time++
        start = size // Advance the start pointer
    }

    if freshTotal == 0 {
        return time
    }
    return -1 // Not all fresh oranges could be rotted
}
