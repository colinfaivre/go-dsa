package training

// 1.8 `M` Valid Sudoku

/*** @LEETCODE https://leetcode.com/problems/valid-sudoku/
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
---
Example 1:
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
---
Example 2:
Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
---
Constraints:
board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=TjFXEUCMqI8
O(9^2) solution - hashmaps
- create hashmaps of hashmaps for rows, cols and squares
- loop in board matrix:
  - ignore dots
  - return false if the char is already in one of the hashmaps
  - add the char to each hashmap
***/

func IsValidSudoku(board [][]byte) bool {
    rowMap := make(map[int]map[byte]bool)
    colMap := make(map[int]map[byte]bool)
    squareMap := make(map[[2]int]map[byte]bool)
    for i := range 9 {
        rowMap[i] = make(map[byte]bool)
        colMap[i] = make(map[byte]bool)
        for j := range 9 {
            squareMap[[2]int{i/3, j/3}] = make(map[byte]bool)
        }
    }

    for i, row := range board {
        for j, val := range row {
            if val == '.' {
                continue
            }

            if rowMap[i][val] || colMap[j][val] || squareMap[[2]int {i/3, j/3}][val]  {
                return false
            }

            rowMap[i][val] = true
            colMap[j][val] = true
            squareMap[[2]int {i/3, j/3}][val] = true
        }
    }

    return true
}
