# 1.8 Valid Sudoku `M`

[leetcode](https://leetcode.com/problems/valid-sudoku/) |
[youtube](https://www.youtube.com/watch?v=TjFXEUCMqI8)

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Example 1:
> - Input: board =
> [["5","3",".",".","7",".",".",".","."]
> ,["6",".",".","1","9","5",".",".","."]
> ,[".","9","8",".",".",".",".","6","."]
> ,["8",".",".",".","6",".",".",".","3"]
> ,["4",".",".","8",".","3",".",".","1"]
> ,["7",".",".",".","2",".",".",".","6"]
> ,[".","6",".",".",".",".","2","8","."]
> ,[".",".",".","4","1","9",".",".","5"]
> ,[".",".",".",".","8",".",".","7","9"]]
> - Output: true

Example 2:
> - Input: board =
> [["8","3",".",".","7",".",".",".","."]
> ,["6",".",".","1","9","5",".",".","."]
> ,[".","9","8",".",".",".",".","6","."]
> ,["8",".",".",".","6",".",".",".","3"]
> ,["4",".",".","8",".","3",".",".","1"]
> ,["7",".",".",".","2",".",".",".","6"]
> ,[".","6",".",".",".",".","2","8","."]
> ,[".",".",".","4","1","9",".",".","5"]
> ,[".",".",".",".","8",".",".","7","9"]]
> - Output: false
> - Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:
> - board.length == 9
> - board[i].length == 9
> - board[i][j] is a digit 1-9 or '.'.

<details>
    <summary><b>O(9^2) solution - arrays</b></summary>

- create arrays of arrays for rows, cols and squares
- loop in board matrix:
  - ignore dots
  - return false if the char is already in one of the arrays
  - add the char to each arrays

```go
func isValidSudoku(board [][]byte) bool {
    var rows [9][9]bool
    var cols [9][9]bool
    var squares [9][9]bool

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' { continue }
            num := board[r][c] - '1' // Convert char '1'-'9' to index 0-8
            squareIndex := (r/3)*3 + c/3 // Unique square index from 0 to 8

            if rows[r][num] || cols[c][num] || squares[squareIndex][num] { return false }

            rows[r][num] = true
            cols[c][num] = true
            squares[squareIndex][num] = true
        }
    }

    return true
}
```

```js
function isValidSudoku(board) {
    const rows = Array.from({ length: 9 }, () => Array(9).fill(false))
    const cols = Array.from({ length: 9 }, () => Array(9).fill(false))
    const squares = Array.from({ length: 9 }, () => Array(9).fill(false))

    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            if (board[r][c] === '.') continue
            const num = board[r][c] - '1' // Convert char '1'-'9' to index 0-8
            const squareIndex = Math.floor(r / 3) * 3 + Math.floor(c / 3) // Unique square index from 0 to 8
            if (rows[r][num] || cols[c][num] || squares[squareIndex][num]) return false
            rows[r][num] = true
            cols[c][num] = true
            squares[squareIndex][num] = true
        }
    }

    return true
}
```
</details>
