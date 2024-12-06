package training

// 9.6 `M` Pacific Atlantic Water Flow

/*** @LEETCODE https://leetcode.com/problems/pacific-atlantic-water-flow/
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean.
The Pacific Ocean touches the island's left and top edges,
and the Atlantic Ocean touches the island's right and bottom edges.
The island is partitioned into a grid of square cells.
You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
The island receives a lot of rain, and the rain water can flow to neighboring cells directly
north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height.
Water can flow from any cell adjacent to an ocean into the ocean.
Return a 2D list of grid coordinates result where
result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
---
Example 1:
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
Explanation: The following cells can flow to the Pacific and Atlantic oceans, as shown below:
[0,4]: [0,4] -> Pacific Ocean
       [0,4] -> Atlantic Ocean
[1,3]: [1,3] -> [0,3] -> Pacific Ocean
       [1,3] -> [1,4] -> Atlantic Ocean
[1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
       [1,4] -> Atlantic Ocean
[2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
       [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
[3,0]: [3,0] -> Pacific Ocean
       [3,0] -> [4,0] -> Atlantic Ocean
[3,1]: [3,1] -> [3,0] -> Pacific Ocean
       [3,1] -> [4,1] -> Atlantic Ocean
[4,0]: [4,0] -> Pacific Ocean
       [4,0] -> Atlantic Ocean
Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.
---
Example 2:
Input: heights = [[1]]
Output: [[0,0]]
Explanation: The water can flow from the only cell to the Pacific and Atlantic oceans.
---
Constraints:
m == heights.length
n == heights[r].length
1 <= m, n <= 200
0 <= heights[r][c] <= 10^5
***/

/*** @SOLUTION https://www.youtube.com/watch?v=s-VkcjHqkGI
***/

func pacificAtlantic(heights [][]int) [][]int {
    ROWS, COLS := len(heights), len(heights[0])
    if ROWS == 0 {
        return [][]int{}
    }

    pacific := make([][]bool, ROWS)
    atlantic := make([][]bool, ROWS)
    for i := 0; i < ROWS; i++ {
        pacific[i] = make([]bool, COLS)
        atlantic[i] = make([]bool, COLS)
    }

    // Start DFS from Pacific and Atlantic borders
    for i := 0; i < ROWS; i++ {
        dfs(i, 0, pacific, heights)       // Left edge (Pacific)
        dfs(i, COLS-1, atlantic, heights)    // Right edge (Atlantic)
    }
    for j := 0; j < COLS; j++ {
        dfs(0, j, pacific, heights)       // Top edge (Pacific)
        dfs(ROWS-1, j, atlantic, heights)    // Bottom edge (Atlantic)
    }

    result := [][]int{}
    for i := 0; i < ROWS; i++ {
        for j := 0; j < COLS; j++ {
            if pacific[i][j] && atlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func dfs(row, col int, ocean [][]bool, heights [][]int) {
    ROWS, COLS := len(heights), len(heights[0])
    ocean[row][col] = true
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    for _, dir := range directions {
        newRow, newCol := row+dir[0], col+dir[1]
        if newRow >= 0 && newRow < ROWS && newCol >= 0 && newCol < COLS &&
            !ocean[newRow][newCol] && heights[newRow][newCol] >= heights[row][col] {
            dfs(newRow, newCol, ocean, heights)
        }
    }
}
