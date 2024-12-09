package training

// 9.7 `M` Surrounded Regions

/*** @LEETCODE https://leetcode.com/problems/surrounded-regions/
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:
- Connect: A cell is connected to adjacent cells horizontally or vertically.
- Region: To form a region connect every 'O' cell.
- Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.
A surrounded region is captured by replacing all 'O's with 'X's in the input matrix board.
---
Example 1:
Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation:
In the above diagram, the bottom region is not captured because it is on the edge of the board and cannot be surrounded.
---
Example 2:
Input: board = [["X"]]
Output: [["X"]]
---
Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=9z2BunfoZ5Y
***/

func solve(board [][]byte) {
    if len(board) == 0 {
        return
    }

    ROWS, COLS := len(board), len(board[0])

    // Helper function to mark border-connected regions using DFS
    var markBorderConnected func(r, c int)
    markBorderConnected = func(r, c int) {
        if r < 0 || r >= ROWS || c < 0 || c >= COLS || board[r][c] != 'O' {
            return
        }

        board[r][c] = 'B' // Mark as border-connected
        markBorderConnected(r-1, c) // Up
        markBorderConnected(r+1, c) // Down
        markBorderConnected(r, c-1) // Left
        markBorderConnected(r, c+1) // Right
    }

    // Step 1: Mark all border-connected 'O' regions
    for r := 0; r < ROWS; r++ {
        if board[r][0] == 'O' {
            markBorderConnected(r, 0)
        }
        if board[r][COLS-1] == 'O' {
            markBorderConnected(r, COLS-1)
        }
    }
    for c := 0; c < COLS; c++ {
        if board[0][c] == 'O' {
            markBorderConnected(0, c)
        }
        if board[ROWS-1][c] == 'O' {
            markBorderConnected(ROWS-1, c)
        }
    }

    // Step 2: Replace all remaining 'O' with 'X' (captured regions)
    // and restore all 'B' to 'O' (border-connected regions)
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if board[r][c] == 'O' {
                board[r][c] = 'X'
            } else if board[r][c] == 'B' {
                board[r][c] = 'O'
            }
        }
    }
}
