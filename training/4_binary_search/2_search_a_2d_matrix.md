# 4.2 Search a 2D Matrix `M`

[leetcode](https://leetcode.com/problems/search-a-2d-matrix/) |
[youtube](https://www.youtube.com/watch?v=Ber2pi2C0j0)

You are given an m x n integer matrix matrix with the following two properties:
- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.
You must write a solution in O(log(m * n)) time complexity.

Example 1:
> - Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
> - Output: true

Example 2:
> - Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
> - Output: false

Constraints:
> - m == matrix.length
> - n == matrix[i].length
> - 1 <= m, n <= 100
> - -10^4 <= matrix[i][j], target <= 10^4

<details>
  <summary><b>O(mn) solution: brute force</b></summary>

- loop in rows
  - loop in cols
    - return true if matrix[row][col] == target
- return false
</details>

<details>
  <summary><b>O(mlogn) solution: binary search every row</b></summary>

- loop in rows
  - binary search the target value
</details>

<details>
  <summary><b>O(logmn) solution: binary search row then col</b></summary>

- binary search the row containing target
- binary search the col containing target

```go
func SearchMatrix(matrix [][]int, target int) bool {
    top, bot := 0, len(matrix) - 1
    for top <= bot {
        row := (top + bot) / 2
        if target > matrix[row][len(matrix[0]) - 1] {
            top = row + 1
        } else if target < matrix[row][0] {
            bot = row - 1
        } else {
            break
        }
    }

    row := (top + bot) / 2
    l, r := 0, len(matrix[0]) - 1
    for l <= r {
        m := (r + l) / 2
        if target > matrix[row][m] {
            l = m + 1
        } else if target < matrix[row][m] {
            r = m - 1
        } else {
            return true
        }
    }

    return false
}
```

```js
/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    let loRow = 0, hiRow = matrix.length-1
    let targetRow = -1
    while (loRow <= hiRow) {
        let midRow = loRow + Math.floor((hiRow-loRow)/2)
        if (matrix[midRow][0] > target) hiRow = midRow-1
        else if (matrix[midRow][matrix[midRow].length-1] < target) loRow = midRow+1
        else {targetRow = midRow; break}
    }

    if (targetRow === -1) return false

    let lo = 0, hi = matrix[targetRow].length-1
    while (lo <= hi) {
        let mid = lo + Math.floor((hi-lo)/2)
        if (matrix[targetRow][mid] > target) hi = mid-1
        else if (matrix[targetRow][mid] < target) lo = mid+1
        else return true
    }

    return false
};
```
</details>
